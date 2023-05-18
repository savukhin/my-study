package processors

type Limiter struct {
	Count int32
	IProcessor
}

func NewLimiter(count int32) *Limiter {
	return &Limiter{
		Count: count,
	}
}
