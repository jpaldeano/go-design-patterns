package main

import "fmt"

// TreeDecorator defines the strategy interface
type TreeDecorator interface {
	Decorate()
}

// ClassicTreeDecorator implements the classic style
type ClassicTreeDecorator struct{}

func (c *ClassicTreeDecorator) Decorate() {
	fmt.Println("Decorating the tree with tinsel, baubles, and a star.")
}

// ModernTreeDecorator implements the modern style
type ModernTreeDecorator struct{}

func (m *ModernTreeDecorator) Decorate() {
	fmt.Println("Decorating the tree with monochrome ornaments and LED lights.")
}

// MinimalistTreeDecorator implements the minimalist style
type MinimalistTreeDecorator struct{}

func (min *MinimalistTreeDecorator) Decorate() {
	fmt.Println("Decorating the tree with a few white lights and simple ornaments.")
}

// TreeContext uses a decorator strategy
type TreeContext struct {
	decorator TreeDecorator
}

// SetDecorator assigns the strategy
func (t *TreeContext) SetDecorator(decorator TreeDecorator) {
	t.decorator = decorator
}

// Decorate executes the selected decoration strategy
func (t *TreeContext) Decorate() {
	if t.decorator == nil {
		fmt.Println("No decoration strategy selected.")
		return
	}
	t.decorator.Decorate()
}

func main() {
	// Create the context
	tree := &TreeContext{}

	// Use the classic style
	tree.SetDecorator(&ClassicTreeDecorator{})
	tree.Decorate()

	// Use the modern style
	tree.SetDecorator(&ModernTreeDecorator{})
	tree.Decorate()

	// Use the minimalist style
	tree.SetDecorator(&MinimalistTreeDecorator{})
	tree.Decorate()
}
