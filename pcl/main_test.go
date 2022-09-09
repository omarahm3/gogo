package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func TestParseInput(t *testing.T) {
	var o Order
	if err := hclsimple.DecodeFile("./testdata/order.hcl", nil, &o); err != nil {
		t.Fatalf("failed: %s", err)
	}

	require.EqualValues(t, Order{
		Contact: &Contact{
			Name:  "Sherlock Holmes",
			Phone: "+44 20 7224 3688",
		},
		Address: &Address{
			Street:  "221B Baker St",
			City:    "London",
			Country: "England",
		},
	}, o)
}

func TestPizza(t *testing.T) {
	var o Order
	if err := hclsimple.DecodeFile("./testdata/pizza.hcl", ctx(), &o); err != nil {
		t.Fatalf("failed: %s", err)
	}

	require.EqualValues(t, Order{
		Pizzas: []*Pizza{
			{
				Size:  "XL",
				Count: 1,
				Toppings: []string{
					"olives",
					"feta_cheese",
					"onions",
				},
			},
		},
	}, o)
}

func TestDiners(t *testing.T) {
	var o Order
	if err := hclsimple.DecodeFile("./testdata/diners.hcl", ctx(), &o); err != nil {
		t.Fatalf("failed: %s", err)
	}

	require.EqualValues(t, 2, o.Pizzas[0].Count)
}

func ctx() *hcl.EvalContext {
	vars := make(map[string]cty.Value)

	for _, size := range []string{"S", "M", "L", "XL"} {
		vars[size] = cty.StringVal(size)
	}
	for _, topping := range []string{"olives", "onions", "feta_cheese", "garlic", "tomato"} {
		vars[topping] = cty.StringVal(topping)
	}

	spec := &function.Spec{
		Type: function.StaticReturnType(cty.Number),
		Params: []function.Parameter{
			{
				Name:      "diners",
				Type:      cty.Number,
				AllowNull: false,
			},
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			d := args[0].AsBigFloat()

			if !d.IsInt() {
				return cty.NilVal, fmt.Errorf("expected int got %q", d)
			}

			di, _ := d.Int64()
			neededSlices := di * 3
			return cty.NumberFloatVal(math.Ceil(float64(neededSlices) / 8)), nil
		},
	}

	return &hcl.EvalContext{
		Variables: vars,
		Functions: map[string]function.Function{
			"for_diners": function.New(spec),
		},
	}
}
