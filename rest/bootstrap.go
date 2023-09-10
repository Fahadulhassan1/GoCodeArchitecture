package rest

import "github.com/goPractice/service"

func StartServer(container *service.Container) *HttpServer {
	httpServer := NewHttpServer(
		container.GbeConfigService.GetConfig().RestServer.Addr,
	)
	sampleService := NewSampleController(container.SampleService)
	httpServer.sampleController = sampleService
	go httpServer.Start()
	return httpServer
}
