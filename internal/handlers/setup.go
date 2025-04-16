package handlers

var (
	Index IIndex
)

// Init 初始化 handler 层的逻辑
func Init() {
	Index = new(index)
}
