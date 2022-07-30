package main

import (
	"fmt"
	"longnd/tutorial/builder2"
)

func main() {
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		fmt.Printf("%p\n", singleton.GetInstance())
	// 	}()
	// }

	/**
		Create normalBuilder
	**/
	// normalBuilder := builder.GetBuilder("normal")
	// iglooBuilder := builder.GetBuilder("igloo")

	// director := builder.NewDirector(normalBuilder)
	// normalHouse := director.BuildHouse()
	// fmt.Printf("Normal House: %s | %s | %d\n", normalHouse.GetWindowType(), normalHouse.GetDoorType(), normalHouse.GetNumberFloor())

	// director.SetBuilder(iglooBuilder)
	// iglooHouse := director.BuildHouse()
	// fmt.Printf("Igloo House: %s | %s | %d\n", iglooHouse.GetWindowType(), iglooHouse.GetDoorType(), iglooHouse.GetNumberFloor())

	/**
		Create carBuilder
	**/
	carBuilder := builder2.GetCarBuilder()
	fmt.Println(carBuilder)

	director := builder2.NewDirector(carBuilder)
	car := director.BuildCar()
	fmt.Printf("Car: %d | %s | %t | %t\n", car.GetSeats(), car.GetEngine(), car.GetTripComputer(), car.GetGPS())
}
