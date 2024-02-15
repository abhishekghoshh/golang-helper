package abstractfactory

type Nike struct {
}
type NikeShirt struct {
	Shirt
}
type NikeShoe struct {
	Shoe
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}
func (n *Nike) makeCustomShoe(logo string, size int) IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: logo,
			size: size,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}
