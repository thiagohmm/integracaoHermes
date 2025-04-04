package interfaceFile

import (
	"github.com/thiagohmm/integracaoHermes/internal/domain/entity"
)

type ProdutoRepository interface {
	GetIntegrRmsProductsIn() ([]entity.IntegrRmsProductIn, error)
	RemoveProductService(rms entity.IntegrRmsProductIn) error
	ProcessProduct(iprID int) entity.Result
}
