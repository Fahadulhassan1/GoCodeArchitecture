package rest

import (
	"io/ioutil"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	addr             string
	sampleController SampelController
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}
func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// gbeConfig := conf.GetConfig()
	// gbeConfig.TempDir = conf.TempDir{
	// 	Path: tempDirectoryName,
	// }
	public := r.Group("/")
	{
		public.GET("/api/sample", server.sampleController.SampleMessage)

	}

	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}
}
