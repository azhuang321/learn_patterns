package builder

func main() {
	mealBuilder := &MealBuilder{}
	vegMeal := mealBuilder.PrepareVegMeal()
	vegMeal.ShowItems()
	vegMeal.GetCost()

	nonDrinkMeal := mealBuilder.PrepareNonDrinkVegMeal()
	nonDrinkMeal.ShowItems()
	nonDrinkMeal.GetCost()
}
