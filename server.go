package zinxlearn

type IServer interface {
	Start()

	Serve()

	Stop()
}
