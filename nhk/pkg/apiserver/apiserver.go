package apiserver

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"aiart2023/pkg/global"
	"aiart2023/pkg/model"
	"aiart2023/pkg/util"
)

var (
	GlobalApiserverController *ApiServerController
)

type ApiServerController struct{}

func NewApiServer() *ApiServerController {
	return &ApiServerController{}
}

func (m *ApiServerController) Run(ctx *gin.Context) {
	meta := &model.RequestMeta{
		StoreLoc: fmt.Sprintf("./%s/", strconv.FormatInt(time.Now().Unix(), 10)),
	}

	if algorithmName := ctx.PostForm("algorithm_name"); algorithmName == "" {
		ctx.JSON(http.StatusBadRequest, "algorithm_name为空")
		return
	} else {
		meta.AlgorithmName = algorithmName
	}

	if fileName := ctx.PostForm("file_name"); fileName == "" {
		ctx.JSON(http.StatusBadRequest, "algorithm_name为空")
		return
	} else {
		meta.FileName = fileName
	}

	folderPath := filepath.Join(meta.StoreLoc, "source")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(folderPath, os.ModePerm); mkdirErr != nil {
			ctx.JSON(http.StatusInternalServerError, "创建文件夹失败")
			return
		}
	}

	folderPath = filepath.Join(meta.StoreLoc, "result")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(folderPath, os.ModePerm); mkdirErr != nil {
			ctx.JSON(http.StatusInternalServerError, "创建文件夹失败")
			return
		}
	}

	defer func() {
		if rmdirErr := os.RemoveAll(meta.StoreLoc); rmdirErr != nil {
			util.InfoLogger.Printf("文件夹%s删除失败\n", meta.StoreLoc)
			return
		}
	}()

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "获取image失败")
		return
	}
	if err := ctx.SaveUploadedFile(file, filepath.Join(meta.StoreLoc, "source", meta.FileName)); err != nil {
		util.ErrorLogger.Printf("存储image失败，err:%s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, "存储image失败")
		return
	} else {
		util.InfoLogger.Printf("存储image成功，临时地址:%s\n", filepath.Join(meta.StoreLoc, "source", meta.FileName))
	}

	ch := make(chan int, 1)
	defer close(ch)

	go func() {
		for {
			if _, err := os.Stat(filepath.Join(meta.StoreLoc, "result", meta.FileName)); !os.IsNotExist(err) {
				ch <- 1
				return
			}
		}
	}()


	cmd := exec.Command("bash",
		global.AlgorithmRegister[meta.AlgorithmName],
		filepath.Join(meta.StoreLoc, "source", meta.FileName),
		filepath.Join(meta.StoreLoc, "result", meta.FileName))

	if cmdErr := cmd.Run(); cmdErr != nil {
		util.ErrorLogger.Printf("调用shell脚本失败，err:%s, cmd:%s\n", cmdErr.Error(),  cmd.String())
		ctx.JSON(http.StatusInternalServerError, "调用shell脚本失败")
		return
	}

	if f := <- ch; f == 0 {
		ctx.JSON(http.StatusInternalServerError, "运行超时")
		util.ErrorLogger.Printf("运行超时\n")
		return
	} else {
		ctx.JSON(http.StatusOK, "服务调用成功")
	}
}
