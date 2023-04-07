package usecase

import (
	"errors"
	"fmt"
	"regexp"
)

type IinGetterUsecase interface {
	GetIINs(s uint64) ([]string, error)
}

type IinGetter struct{}

func NewIinGetter() *IinGetter {
	return &IinGetter{}
}

func (i *IinGetter) GetIINs(s uint64) ([]string, error) {
	var results []string
	// iin := regexp.MustCompile(`(?i)\b(\d{12}|\d{3}-\d{2}-\d{4})\b`)
	iin := regexp.MustCompile(`^((0[48]|[2468][048]|[13579][26])0229[1-6]|000229[34]|\d\d((0[13578]|1[02])(0[1-9]|[12]\d|3[01])|(0[469]|11)(0[1-9]|[12]\d|30)|02(0[1-9]|1\d|2[0-8]))[1-6])\d{5}$`)
	n := fmt.Sprint(s)
	if len(n) <= 11 {
		return nil, errors.New("invalid length, IIN's length must be 12 digits")
	}

	for i := 0; i < len(n)-11; i++ {
		iins := iin.FindAllString(n[i:i+12], -1)
		results = append(results, iins...)
	}
	return results, nil
}
