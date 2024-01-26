package promptuiutils

import "errors"

type Validators []struct {
	Validator func(input string) bool
	Message   string
}

func (v Validators) GenValidatorFunc(input string) error {
	for _, item := range v {
		if !item.Validator(input) {
			return errors.New(item.Message)
		}
	}
	return nil
}
