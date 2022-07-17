package rpc_client

import (
	"github.com/seata/seata-go/pkg/base/protocal"
)

type RpcRMMessage struct {
	RpcMessage    protocal.RpcMessage
	ServerAddress string
}
