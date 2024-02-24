package ymo

import (
	"net/http"
)

type YMOClient struct {
	token      string       //
	counter    string       //
	clientType string       //
	httpClient *http.Client // http client session
	debug      bool         // send events for validation, used for debug
}

type Event struct {
	ClientId string `json:"ClientId"`
	Target   string `json:"Target"`
	DateTime string `json:"DateTime"`
	Price    string `json:"Price"`
	Currency string `json:"Currency"`
}
