package uploadmodel

import (
	"nhaancs/common"
)

type UploadedImage struct {
	common.SQLCreateModel 
	ImageInfo *common.Image `gorm:"column:image_info;"`
}

func (UploadedImage) TableName() string {
	return Image{}.TableName()
}
