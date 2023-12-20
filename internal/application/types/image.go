package types

import "mime/multipart"

type UploadDTO struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`
	UserId uint                  `form:"user_id"`
}

type UploadRet struct {
	Url string `json:"url" binding:"required"`
}
