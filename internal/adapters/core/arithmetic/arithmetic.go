package arithmetic

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (_ Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (_ Adapter) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (_ Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (_ Adapter) Division(a, b int32) (int32, error) {
	return a / b, nil
}
