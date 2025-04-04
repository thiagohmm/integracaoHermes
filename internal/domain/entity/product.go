package entity

import "time"

type IntegrRmsProductIn struct {
	IPR_ID          *int
	JSON            string
	DATARECEBIMENTO time.Time
}

type Result struct {
	Success bool
	Message string
}

type ProductRepository interface {
	GetIntegrRmsProductsIn() ([]IntegrRmsProductIn, error)
	RemoveProductService(rms IntegrRmsProductIn) error
}

type ProductProcessor interface {
	Process(IPR_ID int) Result
}

type LogQueue interface {
	SendLog(data LogErro) error
}

type LogErro struct {
	Tabela string
	Fields []string
	Values []interface{}
}
