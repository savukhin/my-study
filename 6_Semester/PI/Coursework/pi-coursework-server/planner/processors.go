package planner

type IProcessor interface {
	GetName() string
	DoProcess() error
}

type Aggregator struct {
	IProcessor
}

type Deletor struct {
	IProcessor
}

type Returner struct {
	IProcessor
}
