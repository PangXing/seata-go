package session

import (
	"testing"

	"github.com/PangXing/seata-go/pkg/base/meta"
	"github.com/PangXing/seata-go/pkg/util/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBranchSession_Encode_Decode(t *testing.T) {
	bs := branchSessionProvider()
	result, _ := bs.Encode()
	newBs := &BranchSession{}
	newBs.Decode(result)

	assert.Equal(t, bs.TransactionID, newBs.TransactionID)
	assert.Equal(t, bs.BranchID, newBs.BranchID)
	assert.Equal(t, bs.ResourceID, newBs.ResourceID)
	assert.Equal(t, bs.LockKey, newBs.LockKey)
	assert.Equal(t, bs.ClientID, newBs.ClientID)
	assert.Equal(t, bs.ApplicationData, newBs.ApplicationData)
}

func branchSessionProvider() *BranchSession {
	bs := NewBranchSession(
		WithBsTransactionID(uuid.NextID()),
		WithBsBranchID(1),
		WithBsResourceGroupID("my_test_tx_group"),
		WithBsResourceID("tb_1"),
		WithBsLockKey("t_1"),
		WithBsBranchType(meta.BranchTypeAT),
		WithBsStatus(meta.BranchStatusUnknown),
		WithBsClientID("c1"),
		WithBsApplicationData([]byte("{\"data\":\"test\"}")),
	)

	return bs
}
