package car

import "fmt"

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func New(make string, model string, year int) *Car {
	return &Car{
		Make:  make,
		Model: model,
		Year:  year,
	}
}

func (c *Car) Honk() {
	fmt.Printf("This %v says: Beep! Beep!", c.Make+" "+c.Model)
}
