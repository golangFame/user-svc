package validators

import "github.com/go-playground/validator/v10"

type Validator struct {
	*validator.Validate
}

func New() (v *Validator) {
	v = &Validator{
		validator.New(),
	}
	v.RegisterValidation("dummy", validateDummyTag)
	return
}

func (v *Validator) Validating(req interface{}) (err error) {
	return v.Struct(req)
}
