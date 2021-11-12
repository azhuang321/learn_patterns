package builder

import (
	"fmt"
	"go-patterns/build_mode/builder/item"
	"go-patterns/build_mode/builder/item/drink"
	"go-patterns/build_mode/builder/item/food"
)

type Meal struct {
	Item []item.Itemer
}

func (m *Meal) AddItem(item item.Itemer) {
	m.Item = append(m.Item, item)
}

func (m *Meal) GetCost() float64 {
	var cost float64
	for _, v := range m.Item {
		cost += v.Price()
	}
	return cost
}

func (m *Meal) ShowItems() {
	for _, v := range m.Item {
		fmt.Printf("Name:%s,Packing:%s,Price:%.2f\n", v.Name(), v.Packing().Pack(), v.Price())
	}
}

type MealBuilder struct{}

func (m *MealBuilder) PrepareVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(food.VegBurger))
	meal.AddItem(new(drink.Coke))
	return meal
}

func (m *MealBuilder) PrepareNonDrinkVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(food.VegBurger))
	return meal
}
