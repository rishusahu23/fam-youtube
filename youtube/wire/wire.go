//go:build wireinject
// +build wireinject

package wire

import (
	"crypto/tls"
	"github.com/google/wire"
	"gorm.io/gorm"
	"github.com/rishusahu23/fam-youtube/config"
	"github.com/rishusahu23/fam-youtube/external/ohttp"
	youtube2 "github.com/rishusahu23/fam-youtube/external/youtube"
	"github.com/rishusahu23/fam-youtube/youtube"
	"github.com/rishusahu23/fam-youtube/youtube/dao"
	"net/http"
)

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func youtubeProvider(impl *youtube2.ClientImpl) youtube2.Client {
	return impl
}

func httpHandlerProvider(impl *ohttp.HttpRequestHandler) ohttp.IHttpRequestHandler {
	return impl
}

func InitialiseYoutubeService(conf *config.Config, db *gorm.DB) *youtube.Service {
	wire.Build(
		youtube.NewService,
		youtube2.NewYoutubeClientImpl,
		youtubeProvider,
		ohttp.NewHttpRequestHandler,
		getHttpClient,
		dao.RecordWireSet,
	)
	return &youtube.Service{}
}
