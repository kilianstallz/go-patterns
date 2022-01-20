package drinks

type Beverage interface {
	Cost() float32
	GetDescription() string
}

type Espresso struct {
	description string
	cost        float32
}

func NewEspresso() Beverage {
	m := Espresso{"Espresso", 1.99}
	return &m
}

func (e *Espresso) Cost() float32 {
	return e.cost
}

func (e *Espresso) GetDescription() string {
	return e.description
}
