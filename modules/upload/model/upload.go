package uploadmodel

import (
	"nhaancs/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLCreateModel `json:",inline"`
	common.Image          `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

//
//func (u *Upload) Mask(isAdmin bool) {
//	u.GenUID(common.DBTypeUpload, 1)
//}

var (
	ErrFileTooLarge = common.NewCustomError(nil, "file too large", "ErrFileTooLarge")
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}
