package abstract

import (
	"fmt"
)

//ISportsFactory Abstract Factory
type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

//GetSportsFactory Creates a type
func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
