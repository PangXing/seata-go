package getty

import (
	"github.com/PangXing/seata-go/pkg/base/protocal"
)

// MessageFuture ...
type MessageFuture struct {
	ID       int32
	Err      error
	Response interface{}
	Done     chan bool
}

// NewMessageFuture ...
func NewMessageFuture(message protocal.RpcMessage) *MessageFuture {
	return &MessageFuture{
		ID:   message.ID,
		Done: make(chan bool),
	}
}
