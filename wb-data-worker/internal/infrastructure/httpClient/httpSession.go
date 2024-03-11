package httpClient

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	"wb-data-service-golang/wb-data-worker/internal/domain"
)

type _HttpSession struct {
	Client       *http.Client
	SuccessCodes map[int]bool
}

func NewHttpClient(client *http.Client, successCodes map[int]bool) domain.HttpClient {
	return &_HttpSession{
		Client:       client,
		SuccessCodes: successCodes,
	}
}

func (session *_HttpSession) SendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := session.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if session.SuccessCodes[resp.StatusCode] {
		return resp, nil
	}

	return nil, errors.New("Status Cods is Error")
}
