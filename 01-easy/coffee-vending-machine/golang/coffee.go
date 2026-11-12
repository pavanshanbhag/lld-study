package coffeevendingmachine

type Coffee struct {
	name   string
	price  float64
	recipe map[*Ingredient]int
}

func NewCoffee(name string, price float64, recipe map[*Ingredient]int) *Coffee {
	return &Coffee{
		name:   name,
		price:  price,
		recipe: recipe,
	}
}

func (c *Coffee) Name() string {
	return c.name
}

func (c *Coffee) Price() float64 {
	return c.price
}

func (c *Coffee) Recipe() map[*Ingredient]int {
	return c.recipe
}
