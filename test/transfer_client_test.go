package test

import (
	"testing"

	"github.com/tx7do/huobi_futures_golang/sdk/linearswap/restful"
	"github.com/tx7do/huobi_futures_golang/sdk/linearswap/restful/response/transfer"
)

var tfClient restful.TransferClient

func init() {
	tfClient = restful.TransferClient{}
	tfClient.Init(config.AccessKey, config.SecretKey, "api.huobi.pro")
}

func TestTransferClient_TransferAsync(t *testing.T) {
	data := make(chan transfer.TransferResponse)

	go tfClient.TransferAsync(data, "spot", "linear-swap", 1, "BTC-USDT", "USDT")
	x, ok := <-data
	if !ok || x.Success != true {
		t.Logf("%d:%s", x.Code, x.Message)
		t.Fail()
	} else {
		t.Log(x)
	}
}
