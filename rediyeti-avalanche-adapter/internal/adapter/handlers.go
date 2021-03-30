package adapter

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ybbus/jsonrpc/v2"
	"gitlab.red.dev/redi-yeti/rediyeti-avalanche-adapter/pkg/utils"
)

type Response struct {
	ID         string              `json:"jobRunID"`
	StatusCode int                 `json:"StatusCode"`
	Payload    jsonrpc.RPCResponse `json:"data"`
}

func DecodeRequest(c *gin.Context) (*Request, error) {
	var env Request
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &env); err != nil {
		return nil, err
	}
	return &env, nil
}

func (a *Adapter) DefaultHandler(c *gin.Context) {
	request, err := DecodeRequest(c)
	var rpcClient jsonrpc.RPCClient
	if err != nil {
		message := fmt.Sprintf("Decoding request failed with error: %s", err)
		utils.SendError(c, 500, message)
		return
	}
	switch request.Data.Chain {
	case "P":
		rpcClient = a.rpcP
	case "X":
		rpcClient = a.rpcX
	case "C":
		rpcClient = a.rpcC
	default:
		message := "Unknown chain"
		utils.SendError(c, 500, message)
		return
	}
	response, err := rpcClient.Call(request.Data.Method, request.Data.Params)
	if err != nil {
		message := fmt.Sprintf("RPC call failed with error: %s", err)
		utils.SendError(c, 500, message)
		return
	}
	ChainlinkResponse := Response{
		ID:         request.ID,
		StatusCode: 200,
		Payload:    *response,
	}
	c.JSON(200, ChainlinkResponse)
}
