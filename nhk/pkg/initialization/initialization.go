package initialization

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"aiart2023/pkg/apiserver"
	"aiart2023/pkg/global"
	"aiart2023/pkg/util"
)

func Init() {
	apiserver.GlobalApiserverController = apiserver.NewApiServer()
	global.GlobalEngine = gin.New()

	initLogger()
	initRouter()
	initAlgorithmRegister()
}

func initLogger() {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	util.InfoLogger = log.New(io.MultiWriter(f, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	util.ErrorLogger = log.New(io.MultiWriter(f, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func initRouter() {
	user := global.GlobalEngine.Group("/algorithm")
	{
		user.POST("run", apiserver.GlobalApiserverController.Run)
	}
}

func initAlgorithmRegister() {
	global.AlgorithmRegister = make(map[string]string)

	global.AlgorithmRegister["PGP"] = "./algorithm/pgp.sh"

}