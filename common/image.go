package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width,omitempty" gorm:"column:width;"`
	Height    int    `json:"height,omitempty" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"column:cloud_name;"`
	Extension string `json:"extension,omitempty" gorm:"column:extension;"`
}

func (img *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var res Image
	if err := json.Unmarshal(bytes, &res); err != nil {
		return err
	}

	*img = res
	return nil
}

// Value return json value, implement driver.Valuer interface
func (img *Image) Value() (driver.Value, error) {
	if img == nil {
		return nil, nil
	}
	return json.Marshal(img)
}

type Images []Image

func (imgs *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*imgs = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (imgs *Images) Value() (driver.Value, error) {
	if imgs == nil {
		return nil, nil
	}
	return json.Marshal(imgs)
}
