package scrape

import (
	"invest/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoldApi(t *testing.T) {

	var s = Scraper{}
	info, _ := config.NewConfig()

	url := info.Api["gold"].Url
	head := info.Api["gold"].Header

	rtn, err := s.CallApi(url, head)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, rtn)

	t.Log(rtn)
}

func TestBitcoinApi(t *testing.T) {

	var s = Scraper{}
	// var info = config.NewConfigInfo()

	url := "https://api.upbit.com/v1/candles/minutes/1?market=KRW-BTC&count=1" //info.Bitcoin.API.Url
	// id := info.Bitcoin.API.ID
	// key := info.Bitcoin.API.ApiKey

	rtn, err := s.CallApi(url, nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf(rtn)
}

func TestAlpaca(t *testing.T) {
	cp, err := AlpacaCrypto("BTC/USD")
	if err != nil {
		t.Error(err)
	}
	t.Log(cp)
}
