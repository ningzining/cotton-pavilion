# Makefile

APP=cotton-pavilion

.PHONY: build
build:
	@go build -o $(APP)

.PHONY: swag
## swag: 生成swagger文档并导入到目标apifox的项目当中
swag: tools.verify.L-ctl
	@swag fmt
	@swag init
	@L-ctl swag -f docs/swagger.json -p 3495682

tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: tools.install.L-ctl
tools.install.L-ctl:
	@go install github.com/ningzining/L-ctl@latest

#help:
#	@echo "Usage:"
#	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
