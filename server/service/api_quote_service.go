package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Ryan-18-system/desafio-client-server-goexpert/server/model"
	"github.com/Ryan-18-system/desafio-client-server-goexpert/server/util"
)

const (
	quoteUrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func GetCurrentQuote() (*model.QuoteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {

		return nil, util.HandlerErro(err)
	}
	defer resp.Body.Close()
	responseJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, util.HandlerErro(err)
	}
	var quote *model.QuoteResponse
	json.Unmarshal(responseJson, &quote)
	return quote, nil
}
