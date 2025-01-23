package factory

import "fmt"

type GasOven struct {
}

func (g *GasOven) Bake() {
	fmt.Println("Baking with gas oven")
}

type ElectricOven struct {
}

func (e *ElectricOven) Bake() {
	fmt.Println("Baking with electric oven")
}

type WoodOven struct {
}

func (w *WoodOven) Bake() {
	fmt.Println("Baking with wood-fired oven")
}

func main() {
	// Example food item
	food := "pizza"

	// Directly coupled logic to select ovens based on food
	var oven interface{}
	if food == "pizza" {
		oven = &GasOven{}
	} else if food == "bread" {
		oven = &WoodOven{}
	} else if food == "cake" {
		oven = &ElectricOven{}
	} else {
		fmt.Println("No suitable oven for this food")
		return
	}

	switch o := oven.(type) {
	case *GasOven:
		o.Bake()
	case *ElectricOven:
		o.Bake()
	case *WoodOven:
		o.Bake()
	default:
		fmt.Println("No baking performed")
	}
}
