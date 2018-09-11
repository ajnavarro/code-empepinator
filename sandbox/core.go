package sandbox

type Executor interface {
	Execute(...float64) (float64, error)
}
