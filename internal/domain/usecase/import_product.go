package usecase

import (
	"encoding/json"
	"time"

	"github.com/thiagohmm/integracaoHermes/internal/domain/entity"
	interfaceFile "github.com/thiagohmm/integracaoHermes/internal/domain/interfacesFile"
)

type ImportProductUseCase struct {
	ProdutoRepo interfaceFile.ProdutoRepository
	LogQueue    interfaceFile.LogQueue
}

func NewImportProductUseCase(repo interfaceFile.ProdutoRepository, queue interfaceFile.LogQueue) *ImportProductUseCase {
	return &ImportProductUseCase{
		ProdutoRepo: repo,
		LogQueue:    queue,
	}
}

func (uc *ImportProductUseCase) Execute() (bool, error) {
	success := []bool{}
	produtos, err := uc.ProdutoRepo.GetIntegrRmsProductsIn()
	if err != nil {
		return false, err
	}

	for _, rms := range produtos {
		var jsonProduto map[string]interface{}
		if err := json.Unmarshal([]byte(rms.JSON), &jsonProduto); err != nil {
			uc.LogQueue.SendLog(entity.LogErro{
				Tabela: "LogIntegrRMS",
				Fields: []string{"TRANSACAO", "TABELA", "DATARECEBIMENTO", "DATAPROCESSAMENTO", "STATUSPROCESSAMENTO", "JSON", "DESCRICAOERRO"},
				Values: []interface{}{"IN", "PRODUTOS", rms.DATARECEBIMENTO, time.Now(), 1, rms.JSON, err.Error()},
			})
			success = append(success, false)
			uc.ProdutoRepo.RemoveProductService(rms)
			continue
		}

		result := uc.ProdutoRepo.ProcessProduct(*rms.IPR_ID)

		status := 0
		msg := "Integração de Produtos Realizada com Sucesso"
		if !result.Success {
			status = 1
			msg = result.Message
		}

		uc.LogQueue.SendLog(entity.LogErro{
			Tabela: "LogIntegrRMS",
			Fields: []string{"TRANSACAO", "TABELA", "DATARECEBIMENTO", "DATAPROCESSAMENTO", "STATUSPROCESSAMENTO", "JSON", "DESCRICAOERRO"},
			Values: []interface{}{"IN", "PRODUTOS", rms.DATARECEBIMENTO, time.Now(), status, rms.JSON, msg},
		})

		success = append(success, result.Success)
		uc.ProdutoRepo.RemoveProductService(rms)
	}

	for _, ok := range success {
		if !ok {
			return false, nil
		}
	}
	return true, nil
}
