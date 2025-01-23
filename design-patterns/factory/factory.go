package patterns

import "fmt"

// When building logic that depends on the type of object
// we are working with, we can use a factory to abstract
// the object creation and avoid dependencies on concrete classes

// Factory centralizes de object creation
// and abstracts away concrete class dependencies

type Oven interface {
	Bake()
}

type GasOven struct {
}

func (g *GasOven) Bake() {
	fmt.Println("Burning gas")
}

type ElectricOven struct {
}

func (e *ElectricOven) Bake() {
	fmt.Println("Burning electricity")
}

type WoodOven struct {
}

func (w *WoodOven) Bake() {
	fmt.Println("Burning wood")
}

func OvenFactory(food string) Oven {
	switch food {
	case "pizza":
		return &GasOven{}
	case "bread":
		return &WoodOven{}
	case "cake":
		return &ElectricOven{}
	default:
		return nil
	}
}

func main() {
	// Example food item
	food := "pizza"

	oven := OvenFactory(food)
	if oven == nil {
		fmt.Println("No suitable oven for this food")
		return
	}

	oven.Bake()
}
