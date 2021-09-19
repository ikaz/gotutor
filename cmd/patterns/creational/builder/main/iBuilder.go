package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		// instance
		return newNormalBuilder()
	}
	if builderType == "igloo" {
		// vs type
		return &iglooBuilder{}
	}
	return nil
}
