package adminController

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	util "hmshop/utils"
)

type CommonApi struct {
}

func (cm CommonApi) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	path, err := util.UploadFile(file)
	if err != nil {
		res.FailWithMessage(code.UploadError, c)
		return
	}
	res.Ok(path, code.UploadSuccess, c)
}
