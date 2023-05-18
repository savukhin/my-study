package processors

type Returner struct {
	IProcessor
}

func NewReturner() *Returner {
	return &Returner{}
}
