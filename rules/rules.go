package rules

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name                string `validate:"required,min=4,max=15"`
	Email               string `validate:"required,email"`
	Age                 int    `validate:"required,numeric,min=18"`
	DriverLicenseNumber string `validate:"omitempty,len=12,numeric"`
}

func main() {
	//validate := validator.New()
	p := Person{
		Name:                "Ume",
		Email:               "usonawane@flexera.com",
		Age:                 16,
		DriverLicenseNumber: "",
	}
	err := p.ValidatePerson()
	fmt.Println(err)
}

func (p *Person) ValidatePerson() error {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return err
	}
	return err
}
