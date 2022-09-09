package main

type (
	Contact struct {
		Name  string `hcl:"name"`
		Phone string `hcl:"phone"`
	}

	Address struct {
		Street  string `hcl:"street"`
		City    string `hcl:"city"`
		Country string `hcl:"country"`
	}

	Order struct {
		Contact *Contact `hcl:"contact,block"`
		Address *Address `hcl:"address,block"`
		Pizzas  []*Pizza `hcl:"pizza,block"`
	}

	Pizza struct {
		Size     string   `hcl:"size"`
		Count    int      `hcl:"count,optional"`
		Toppings []string `hcl:"toppings,optional"`
	}
)

func main() {
}
