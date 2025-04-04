package interface


type LogQueue interface {
	SendLog(data entity.LogErro) error
}