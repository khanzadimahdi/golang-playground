package builder

import "errors"

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

type house struct {
	windowType string
	doorType   string
	floor      int
}

func getBuilder(builder string) (iBuilder, error) {
	if builder == "normal" {
		return new(normalBuilder), nil
	} else if builder == "igloo" {
		return new(iglooBuilder), nil
	}

	return nil, errors.New("builder is not exists")
}

// how to use:
// normalBuilder := getBuilder("normal")
// iglooBuilder := getBuilder("igloo")

// director := newDirector(normalBuilder)
// normalHouse := director.buildHouse()
// fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)

// director.setBuilder(iglooBuilder)
// iglooHouse := director.buildHouse()
// fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
