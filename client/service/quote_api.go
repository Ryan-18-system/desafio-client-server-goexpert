package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Ryan-18-system/desafio-client-server-goexpert/client/util"
)

const (
	url string = "http://localhost:8080/cotacao"
)

func GetCurrentQuote() (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, util.HandlerError(err)
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, util.HandlerError(err)
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("erro ao consumir api de cotação %s", string(responseBody)))
	}
	quote := string(responseBody)
	return &quote, nil
}
