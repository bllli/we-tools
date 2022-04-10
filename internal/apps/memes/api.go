package memes

import (
	"github.com/gin-gonic/gin"
	"strings"
	"we-tools/internal/common"
)

type Api struct {
	usecase Usecase
}

// NewApi creates a new memes api
func NewApi(usecase Usecase) *Api {
	return &Api{usecase}
}

func (api *Api) UploadMeme(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.Fail(c, "file parameter is missing")
		return
	}
	allowedSuffixes := map[string]bool{
		"jpg": true,
		"png": true,
	}
	filename := file.Filename
	fileSuffix := strings.ToLower(strings.Split(filename, ".")[1])
	if !allowedSuffixes[fileSuffix] {
		common.Fail(c, "file suffix is not allowed")
		return
	}
	f, err := file.Open()
	if err != nil {
		common.Fail(c, "file open error"+err.Error())
		return
	}
	fileContent := make([]byte, file.Size)
	_, err = f.Read(fileContent)

	title, ok := c.GetPostForm("title")
	if !ok {
		common.Fail(c, "title parameter is missing")
		return
	}

	inputDto := &CreateMemeInputDto{
		Title:          title,
		FileContent:    fileContent,
		Type:           MemeTypeImage,
		FilenameSuffix: fileSuffix,
	}
	outputDto, err := api.usecase.CreateMeme(inputDto)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.OK(c, outputDto)
}
