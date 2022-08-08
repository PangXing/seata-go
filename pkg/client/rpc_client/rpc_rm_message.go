package rpc_client

import (
	"github.com/PangXing/seata-go/pkg/base/protocal"
)

type RpcRMMessage struct {
	RpcMessage    protocal.RpcMessage
	ServerAddress string
}
