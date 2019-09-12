package app

import (
	comp "collector/components"
	"github.com/sirupsen/logrus"
)

func ApiClient(log *logrus.Entry) *comp.ApiClient {
	apiClient := comp.NewApiClient(Config().MyTargetHost, log)
	return apiClient
}

func Proxy(log *logrus.Entry) *comp.Proxy {
	proxy := comp.NewProxy(Config().TokenProxyHost, log)
	return proxy
}
