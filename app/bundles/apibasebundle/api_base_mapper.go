package apibasebundle

// KittiesMapper define the base contract for mapper
type ApiBaseMapper interface {
	Alive() float64
	Env() interface{}
}