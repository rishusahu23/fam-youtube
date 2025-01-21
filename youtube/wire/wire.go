//go:build wireinject
// +build wireinject

package wire

import (
	"crypto/tls"
	"github.com/google/wire"
	"github.com/rishusahu23/fam-youtube/youtube"
	"net/http"
)

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func InitialiseYoutubeService() *youtube.Service {
	wire.Build(
		youtube.NewService,
	)
	return &youtube.Service{}
}
