package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/ningzining/cotton-pavilion/pkg/logger"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
)

type consoleColorModeValue int

const (
	autoColor consoleColorModeValue = iota
	disableColor
	forceColor
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

var consoleColorMode = autoColor

type LoggerConfig struct {
	Formatter LogFormatter
	Output    io.Writer
	SkipPaths []string
}

type LogFormatter func(params LogFormatterParams) string

type LogFormatterParams struct {
	TimeStamp  time.Time
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration

	ClientIP string
	Method   string
	Path     string

	ErrorMessage string

	isTerm   bool
	Request  []byte
	Response []byte

	Keys map[string]any
}

func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.StatusCode

	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

func (p *LogFormatterParams) MethodColor() string {
	method := p.Method

	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

func (p *LogFormatterParams) ResetColor() string {
	return reset
}

func (p *LogFormatterParams) IsOutputColor() bool {
	return consoleColorMode == forceColor || (consoleColorMode == autoColor && p.isTerm)
}

var defaultLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	return fmt.Sprintf("%s%3d%s - [%s] %v %s%s%s %s\n%s\n%s\n %s",
		statusColor, param.StatusCode, resetColor,
		param.ClientIP,
		param.Latency,
		methodColor, param.Method, resetColor,
		param.Path,
		string(param.Request),
		string(param.Response),
		param.ErrorMessage,
	)
}

func ForceConsoleColor() {
	consoleColorMode = forceColor
}

func Logger() gin.HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: nil,
		Output:    nil,
		SkipPaths: nil,
	})
}

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(conf LoggerConfig) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = gin.DefaultWriter
	}

	notLogged := conf.SkipPaths

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	if isTerm {
		ForceConsoleColor()
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notLogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		requestBytes := requestFromCtx(c)

		// 生成自定义的writer代替gin.DefaultWriter
		writer := newWriter(c)
		c.Writer = writer

		c.Next()

		if _, ok := skip[path]; !ok {
			end := time.Now()
			param := LogFormatterParams{
				TimeStamp:    end,
				StatusCode:   c.Writer.Status(),
				Latency:      end.Sub(start),
				ClientIP:     c.ClientIP(),
				Method:       c.Request.Method,
				Path:         c.Request.RequestURI,
				ErrorMessage: c.Errors.ByType(gin.ErrorTypePrivate).String(),
				isTerm:       false,
				Request:      requestBytes,
				Response:     writer.bodyBuf.Bytes(),
				Keys:         c.Keys,
			}

			// todo: 待研究存入接口日志
			log.Info(formatter(param))
		}
	}
}

type responseWriter struct {
	gin.ResponseWriter               // 继承原有 gin.ResponseWriter
	bodyBuf            *bytes.Buffer // Body 内容临时存储位置，这里指针，原因这个存储对象要复用
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if count, err := w.bodyBuf.Write(b); err != nil { // 写入数据时，也写入一份数据到缓存中
		return count, err
	}
	return w.ResponseWriter.Write(b) // 原始框架数据写入
}

func newWriter(c *gin.Context) *responseWriter {
	return &responseWriter{
		ResponseWriter: c.Writer,
		bodyBuf:        bytes.NewBuffer([]byte{}),
	}
}

func requestFromCtx(c *gin.Context) []byte {
	if c.Request.Method == http.MethodGet {
		return requestParam(c)
	}

	return requestBody(c)
}

func requestBody(c *gin.Context) []byte {
	requestBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	// 重新在缓冲区中写入数据
	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBytes))

	// 删除空格和换行符,windows下换行符是\r\n,linux环境下换行符是\n
	requestBytes = bytes.ReplaceAll(requestBytes, []byte(" "), []byte{})
	requestBytes = bytes.ReplaceAll(requestBytes, []byte("\r\n"), []byte{})
	requestBytes = bytes.ReplaceAll(requestBytes, []byte("\n"), []byte{})

	return requestBytes
}

func requestParam(c *gin.Context) []byte {
	query := c.Request.URL.Query()

	queryParams := make(map[string]any)
	for k, v := range query {
		queryParams[k] = v
	}

	requestParamBytes, err := json.Marshal(queryParams)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return requestParamBytes
}
