package restapis

import (
	"common/wrapper"
	"github.com/sirupsen/logrus"
)

type APIService struct {
	baseURI string
	*wrapper.EchoWrapper
	logger *logrus.Entry
}

func NewAPIService(port string) *APIService {
	api := &APIService{
		baseURI:     "/api/v1/area",
		EchoWrapper: wrapper.NewEchoWrapper(port),
	}

	return api
}
