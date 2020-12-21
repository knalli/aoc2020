package day21

import (
	"fmt"
	"github.com/knalli/aoc"
	"sort"
	"strings"
)

func solve1(lines []string) error {
	foods := parseFoods(lines)
	countIngredientFreeAllergens, _ := solver(foods)
	aoc.PrintSolution(fmt.Sprintf("Allergen free ingredients: %d", countIngredientFreeAllergens))
	return nil
}

func solve2(lines []string) error {
	foods := parseFoods(lines)
	_, ingredientList := solver(foods)
	aoc.PrintSolution(fmt.Sprintf("Ingredient list: %s", ingredientList))
	return nil
}

type Food struct {
	Ingredients []string
	Allergens   []string
}

func parseFoods(lines []string) []Food {
	result := make([]Food, 0)
	for _, line := range lines {
		result = append(result, Food{
			Ingredients: strings.Split(line[:strings.Index(line, " (")], " "),
			Allergens:   strings.Split(line[strings.Index(line, " (")+11:len(line)-1], ", "),
		})
	}
	return result
}

func solver(foods []Food) (int, string) {
	ingredientAllergens := make(map[string]map[string]int)
	allergenIngredients := make(map[string]map[string]int)
	for _, food := range foods {
		for _, ingredient := range food.Ingredients {
			for _, allergen := range food.Allergens {

				if _, exist := ingredientAllergens[ingredient]; !exist {
					ingredientAllergens[ingredient] = make(map[string]int)
				}
				ingredientAllergens[ingredient][allergen] = ingredientAllergens[ingredient][allergen] + 1

				if _, exist := allergenIngredients[allergen]; !exist {
					allergenIngredients[allergen] = make(map[string]int)
				}
				allergenIngredients[allergen][ingredient] = allergenIngredients[allergen][ingredient] + 1

			}
		}
	}

	for f, food := range foods {
		for _, ingredient := range food.Ingredients {
			exclusiveIngredient := true
			for o, otherFood := range foods {
				if f == o {
					continue
				}
				for _, otherIngredient := range otherFood.Ingredients {
					if ingredient == otherIngredient {
						exclusiveIngredient = false
						break
					}
				}
				if !exclusiveIngredient {
					break
				}
			}

			// if exclusive, all well-known allergens are not possible any more
			if exclusiveIngredient {
				for allergen, allergenNum := range ingredientAllergens[ingredient] {
					if allergenNum == 0 {
						delete(ingredientAllergens[ingredient], allergen)
						continue
					}
					used := false
					for o, otherFood := range foods {
						if f == o {
							continue
						}
						for _, otherIngredient := range otherFood.Ingredients {
							for otherAllergen, otherAllergenNum := range ingredientAllergens[otherIngredient] {
								if otherAllergenNum == 0 {
									delete(ingredientAllergens[otherIngredient], otherAllergen)
									continue
								}
								if allergen == otherAllergen {
									used = true
									break
								}
							}
							if used {
								break
							}
						}
						if used {
							break
						}
					}
					if used {
						delete(ingredientAllergens[ingredient], allergen)
					}
				}
			}

			for _, allergen := range food.Allergens {
				for o, otherFood := range foods {
					if f == o {
						continue
					}
					otherHasIngredient := false
					for _, otherIngredient := range otherFood.Ingredients {
						if ingredient == otherIngredient {
							otherHasIngredient = true
							break
						}
					}
					// look into all other foods allergen list: if marked but ingredient is not listed, this ingredient cannot have that allergen.
					if !otherHasIngredient {
						otherHasAllergen := false
						for _, otherAllergen := range otherFood.Allergens {
							if allergen == otherAllergen {
								otherHasAllergen = true
							}
						}
						if otherHasAllergen {
							delete(ingredientAllergens[ingredient], allergen)
							delete(allergenIngredients[allergen], ingredient)
						}
					}
				}
			}
		}
	}

	// reduce list (an exclude allergen can be removed from all other options)
	for i := 0; i < len(allergenIngredients); i++ {
		for allergen := range allergenIngredients {
			if len(allergenIngredients[allergen]) == 1 {
				for ingredient := range allergenIngredients[allergen] {
					for otherAllergen := range allergenIngredients {
						if allergen == otherAllergen {
							continue
						}
						delete(allergenIngredients[otherAllergen], ingredient)
						delete(ingredientAllergens[ingredient], otherAllergen)
					}
				}
			}
		}
	}

	ingredientFreeAllergens := make([]string, 0)
	for ingredient, m := range ingredientAllergens {
		if len(m) == 0 {
			ingredientFreeAllergens = append(ingredientFreeAllergens, ingredient)
		}
	}

	countIngredientFreeAllergens := 0
	for _, food := range foods {
		for _, ingredient := range food.Ingredients {
			for _, s := range ingredientFreeAllergens {
				if ingredient == s {
					countIngredientFreeAllergens++
				}
			}
		}
	}

	allAllergens := make([]string, 0)
	for allergen := range allergenIngredients {
		allAllergens = append(allAllergens, allergen)
	}
	sort.Strings(allAllergens)

	allIngredients := make([]string, 0)
	for _, allergen := range allAllergens {
		for ingredient := range allergenIngredients[allergen] {
			allIngredients = append(allIngredients, ingredient)
		}
	}

	return countIngredientFreeAllergens, strings.Join(allIngredients, ",")
}
