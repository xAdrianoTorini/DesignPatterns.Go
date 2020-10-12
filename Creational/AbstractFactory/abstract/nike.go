package abstract

//Nike type
type Nike struct {
}

//MakeShoe interface implementation
func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

//MakeShort interface implementation
func (n *Nike) MakeShort() IShort {
	return &NikeShort{
		Short: Short{
			logo: "nike",
			size: 14,
		},
	}
}
