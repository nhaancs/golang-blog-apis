package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"nhaancs/common"
	"nhaancs/component/uploadprovider"
	"nhaancs/modules/upload/model"
	"path/filepath"
	"strings"
	"time"
)

type ImageStore interface {
	CreateImage(context context.Context, data *uploadmodel.UploadedImage) error
}

type uploadImageBiz struct {
	provider uploadprovider.UploadProvider
	imgStore ImageStore
}

func NewUploadImageBiz(provider uploadprovider.UploadProvider, imgStore ImageStore) *uploadImageBiz {
	return &uploadImageBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadImageBiz) UploadImage(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	dst := fmt.Sprintf("%s/%s", folder, fileName)
	img, err := biz.provider.SaveFile(ctx, data, dst)
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt
	
	uploadedImage := &uploadmodel.UploadedImage{}
	uploadedImage.ImageInfo = img
	if err := biz.imgStore.CreateImage(ctx, uploadedImage); err != nil {
		biz.provider.DeleteFile(ctx, dst)
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
