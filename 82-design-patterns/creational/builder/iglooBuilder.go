package builder

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Iron Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Iron Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 3
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
