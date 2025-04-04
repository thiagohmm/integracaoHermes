package repository

import (
	"github.com/thiagohmm/integracaoHermes/internal/domain/entity"
)

type ProdutoOracleRepository struct {
	// Conn *sql.DB ou outra conexão Oracle
}

func (r *ProdutoOracleRepository) GetIntegrRmsProductsIn() ([]entity.IntegrRmsProductIn, error) {
	// buscar produtos no banco
	return []entity.IntegrRmsProductIn{}, nil
}

func (r *ProdutoOracleRepository) RemoveProductService(rms entity.IntegrRmsProductIn) error {
	// remover produto do Oracle
	return nil
}

func (r *ProdutoOracleRepository) ProcessProduct(iprID int) entity.Result {
	// chamada ao dopkg_produto
	return entity.Result{
		Success: true,
		Message: "Processado com sucesso",
	}
}
