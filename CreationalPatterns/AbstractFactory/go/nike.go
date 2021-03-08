package main

type nike struct {

}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{shoe{
		logo: "nike",
		size: 14,
	}}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{shirt{
		logo: "nike",
		size: 14,
	}}
}
