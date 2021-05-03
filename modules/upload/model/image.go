package uploadmodel

import (
	"nhaancs/common"
)

const EntityName = "Image"

type Image struct {
	common.SQLModel `json:",inline"`
	ImageInfo       *common.Image `json:",inline" gorm:"column:image_info;"`
}

func (Image) TableName() string {
	return "images"
}

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
