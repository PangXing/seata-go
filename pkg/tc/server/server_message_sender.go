package server

import (
	"time"

	"github.com/PangXing/seata-go/pkg/base/protocal"
	getty "github.com/apache/dubbo-getty"
)

type ServerMessageSender interface {

	// Send response.
	SendResponse(request protocal.RpcMessage, session getty.Session, msg interface{})

	// Sync call to RM
	SendSyncRequest(resourceID string, clientID string, message interface{}) (interface{}, error)

	// Sync call to RM with timeout.
	SendSyncRequestWithTimeout(resourceID string, clientID string, message interface{}, timeout time.Duration) (interface{}, error)

	// Send request with response object.
	SendSyncRequestByGetty(session getty.Session, message interface{}) (interface{}, error)

	// Send request with response object.
	SendSyncRequestByGettyWithTimeout(session getty.Session, message interface{}, timeout time.Duration) (interface{}, error)

	// ASync call to RM
	SendASyncRequest(session getty.Session, message interface{}) error
}
