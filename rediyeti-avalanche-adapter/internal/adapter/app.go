package adapter

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ybbus/jsonrpc/v2"
)

type Config struct {
	Port int

	ApiHost string
	ApiPort int
}

type Adapter struct {
	cfg  Config
	rpcP jsonrpc.RPCClient
	rpcX jsonrpc.RPCClient
	rpcC jsonrpc.RPCClient
}

type Data struct {
	Method string      `json:"method"`
	Chain  string      `json:"chain"`
	Params interface{} `json:"params"`
}
type Request struct {
	ID   string `json:"id"`
	Data Data   `json:"data"`
}

func NewAdapter(cfg Config) *Adapter {
	baseUrl := fmt.Sprintf("%s:%d", cfg.ApiHost, cfg.ApiPort)
	pRpcClient := jsonrpc.NewClient(fmt.Sprintf("%s/ext/P", baseUrl))
	xRpcClient := jsonrpc.NewClient(fmt.Sprintf("%s/ext/bc/X", baseUrl))
	cRpcClient := jsonrpc.NewClient(fmt.Sprintf("%s/ext/bc/C/rpc", baseUrl))
	a := Adapter{
		cfg:  cfg,
		rpcP: pRpcClient,
		rpcX: xRpcClient,
		rpcC: cRpcClient,
	}
	return &a
}

func (a *Adapter) Serve() error {
	r := gin.Default()
	r.POST("/", a.DefaultHandler)
	err := r.Run(fmt.Sprintf(":%d", a.cfg.Port))
	return err
}
