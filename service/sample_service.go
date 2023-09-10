package service

type SampleService interface {
	SampleFunc() (string, error)
}
type sampleService struct {
}

func NewSampleService() SampleService {
	return &sampleService{}
}

func (s *sampleService) SampleFunc() (string, error) {
	return "this is sample", nil

}
