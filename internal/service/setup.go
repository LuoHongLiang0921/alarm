package service

var (
	IndexService IIndexService
)

func Init() {
	IndexService = NewIndexService()
}
