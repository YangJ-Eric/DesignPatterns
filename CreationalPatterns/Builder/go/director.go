package main

type director struct {
	build iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{build: b}
}

func (d *director) setBuild(b iBuilder)  {
	d.build = b
}

func (d *director) buildHouse() house {
	d.build.setDoorType()
	d.build.setWindowType()
	d.build.setNumFloor()
	return d.build.getHouse()
}