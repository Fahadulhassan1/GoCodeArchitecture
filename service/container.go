package service

type Container struct {
	SampleService    SampleService
	GbeConfigService GbeConfigService
}

func NewServiceContainer() *Container {
	sampelService := NewSampleService()
	gGbeConfigService := NewGbeConfigService()

	return &Container{
		SampleService:    sampelService,
		GbeConfigService: gGbeConfigService,
	}
}
