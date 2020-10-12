package abstract

//Adidas type
type Adidas struct {
}

//MakeShoe interface implementation
func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

//MakeShort interface implementation
func (a *Adidas) MakeShort() IShort {
	return &AdidasShort{
		Short: Short{
			logo: "adidas",
			size: 14,
		},
	}
}
