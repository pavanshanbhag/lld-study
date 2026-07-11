package decorator

// SugarDecorator adds sugar to a beverage.
type SugarDecorator struct {
	*BeverageDecorator
}

func NewSugarDecorator(beverage Beverage) *SugarDecorator {
	return &SugarDecorator{
		BeverageDecorator: NewBeverageDecorator(beverage),
	}
}

func (d *SugarDecorator) GetDescription() string {
	return d.beverage.GetDescription() + " with Sugar"
}

func (d *SugarDecorator) Cost() float64 {
	return d.beverage.Cost() + 0.2
}
