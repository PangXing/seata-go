package file

import (
	"strings"

	"github.com/PangXing/seata-go/pkg/base/constant"
	"github.com/PangXing/seata-go/pkg/base/extension"
	"github.com/PangXing/seata-go/pkg/base/registry"
	"github.com/PangXing/seata-go/pkg/client/config"
	"github.com/PangXing/seata-go/pkg/util/log"
)

func init() {
	extension.SetRegistry(constant.FileKey, newFileRegistry)
}

type fileRegistry struct {
}

func (r *fileRegistry) Register(addr *registry.Address) error {
	//文件不需要注册
	log.Info("file register")
	return nil
}

func (r *fileRegistry) UnRegister(addr *registry.Address) error {
	return nil
}
func (r *fileRegistry) Lookup() ([]string, error) {
	addressList := strings.Split(config.GetClientConfig().TransactionServiceGroup, ",")
	return addressList, nil
}
func (r *fileRegistry) Subscribe(notifyListener registry.EventListener) error {
	return nil
}

func (r *fileRegistry) UnSubscribe(notifyListener registry.EventListener) error {
	return nil
}

func (r *fileRegistry) Stop() {
	// TODO: Implement Stop interface
	return
}

// newNacosRegistry will create new instance
func newFileRegistry() (registry.Registry, error) {
	tmpRegistry := &fileRegistry{}
	return tmpRegistry, nil
}
