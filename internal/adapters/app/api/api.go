package api

import "github.com/lschulzes/hex_arch/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DBPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DBPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiA.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apiA.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiA Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiA.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiA.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
