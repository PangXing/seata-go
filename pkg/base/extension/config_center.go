package extension

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/seata/seata-go/pkg/base/config"
	"github.com/seata/seata-go/pkg/base/config_center"
)

var (
	configCentersMu sync.RWMutex
	configCenters   = make(map[string]func(conf *config.ConfigCenterConfig) (config_center.DynamicConfigurationFactory, error))
)

// SetConfigCenter set config center
func SetConfigCenter(name string, v func(conf *config.ConfigCenterConfig) (config_center.DynamicConfigurationFactory, error)) {
	configCentersMu.Lock()
	defer configCentersMu.Unlock()
	if v == nil {
		panic("configCenter: Register  configCenter is nil")
	}
	if _, dup := configCenters[name]; dup {
		panic("configCenter: Register called twice for configCenter " + name)
	}
	configCenters[name] = v
}

// GetConfigCenter get config center
func GetConfigCenter(name string, conf *config.ConfigCenterConfig) (config_center.DynamicConfigurationFactory, error) {
	configCentersMu.RLock()
	configCenter := configCenters[name]
	configCentersMu.RUnlock()
	if configCenter == nil {
		return nil, errors.Errorf("config center for " + name + " is not existing, make sure you have import the package.")
	}
	return configCenter(conf)
}
