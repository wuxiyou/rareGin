package utils

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"os"
	"io"
)

func SingeUpload(c *gin.Context)  {
	name := c.PostForm("name")
	fmt.Println(name)
	file, header, err := c.Request.FormFile("upload")

	if err != nil {
		c.String(http.StatusBadRequest, "错误请求")
		return
	}

	fileName := header.Filename

	out, err := os.Create(fileName)

	if err != nil {
		c.String(400, "文件创建失败")
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {

	}

	c.String(http.StatusCreated, "上传成功")
}
