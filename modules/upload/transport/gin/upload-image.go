package ginupload

import (
	_ "image/jpeg"
	_ "image/png"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/upload/biz"
	"nhaancs/modules/upload/store"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstore.NewSQLStore(db)
		biz := uploadbiz.NewUploadImageBiz(appCtx.UploadProvider(), imgStore)
		img, err := biz.UploadImage(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
