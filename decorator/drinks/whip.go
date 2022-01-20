package drinks

// Whip : Condiment
type Whip struct {
	beverage    Beverage
	description string
	cost        float32
}

// NewWhip : Constructor
func NewWhip(b Beverage) Beverage {
	m := Whip{b, "Whip", 0.10}
	return &m
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + " " + w.description
}

func (w *Whip) Cost() float32 {
	return w.beverage.Cost() + w.cost
}
