package wire

import (
	"hmshop/internal/api/provider"

	"github.com/google/wire"
)

func InitHandlers() (*provider.AllController, error) {
	wire.Build(provider.ProviderController)
	return nil, nil
}
