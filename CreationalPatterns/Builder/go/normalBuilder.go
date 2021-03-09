package main

type normalBuilder struct {
	house
}

func newNormalBuilder() *normalBuilder{
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType()  {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType()  {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor()  {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return b.house
}