package pkg

import (
	"fmt"
)

type SyncRequest interface {
	GetMethod() string
	GetURL() string
	GetResponse() Response
	Marshal() ([]byte, error)
}

type Response interface {
	Unmarshal(b []byte) (interface{}, error)
}

type Request interface {
}

type SyncRequestFactory func(req interface{}) SyncRequest

func NewVendorRequest(req Request, reqFact SyncRequestFactory) (SyncRequest, error) {
	vendorReq := reqFact(req)
	if vendorReq == nil {
		return nil, fmt.Errorf("invalid request %v", req)
	}
	return vendorReq, nil
}
