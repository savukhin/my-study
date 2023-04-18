package processors

type IProcessor interface {
	DoProcess() error
}

type Deletor struct {
	IProcessor
}

type Returner struct {
	IProcessor
}
