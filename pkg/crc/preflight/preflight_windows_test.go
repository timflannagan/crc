package preflight

import (
	"testing"

	"github.com/code-ready/crc/pkg/crc/config"
	"github.com/code-ready/crc/pkg/crc/network"
	"github.com/stretchr/testify/assert"
)

func TestCountConfigurationOptions(t *testing.T) {
	cfg := config.New(config.NewEmptyInMemoryStorage())
	RegisterSettings(cfg)
	assert.Len(t, cfg.AllConfigs(), 13)
}

func TestCountPreflights(t *testing.T) {
	assert.Len(t, getPreflightChecks(false, false, network.SystemNetworkingMode), 18)
	assert.Len(t, getPreflightChecks(true, true, network.SystemNetworkingMode), 18)

	assert.Len(t, getPreflightChecks(false, false, network.UserNetworkingMode), 19)
	assert.Len(t, getPreflightChecks(true, true, network.UserNetworkingMode), 19)
}
