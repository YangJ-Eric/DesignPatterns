package main

type iglooBuilder struct {
	house
}

func newIglooBuilder() *iglooBuilder {
	return  &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType()  {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType()  {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor()  {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return b.house
}