package rpc_client

import (
	"time"

	"github.com/PangXing/seata-go/pkg/base/protocal"
)

type ClientMessageSender interface {

	// Send msg with response object.
	SendMsgWithResponse(msg interface{}) (interface{}, error)

	// Send msg with response object.
	SendMsgWithResponseAndTimeout(msg interface{}, timeout time.Duration) (interface{}, error)

	// Send response.
	SendResponse(request protocal.RpcMessage, serverAddress string, msg interface{})
}
