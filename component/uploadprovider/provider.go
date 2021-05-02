package uploadprovider

import (
	"context"
	"nhaancs/common"
)

type UploadProvider interface {
	SaveFile(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
