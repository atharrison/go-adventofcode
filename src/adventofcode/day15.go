package adventofcode

import "fmt"

/*
Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1
*/

type Ingredient struct {
	Name       string
	Capacity   int64
	Durability int64
	Flavor     int64
	Texture    int64
	Calories   int64
}

func ExecuteDay15() {

	ingredients := []*Ingredient{
		&Ingredient{
			//Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
			Name:       "Frosting",
			Capacity:   4,
			Durability: -2,
			Flavor:     0,
			Texture:    0,
			Calories:   5,
		},
		&Ingredient{
			//Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
			Name:       "Candy",
			Capacity:   0,
			Durability: 5,
			Flavor:     -1,
			Texture:    0,
			Calories:   8,
		},
		&Ingredient{
			//Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
			Name:       "Butterscotch",
			Capacity:   -1,
			Durability: 0,
			Flavor:     5,
			Texture:    0,
			Calories:   6,
		},
		&Ingredient{
			//Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1
			Name:       "Sugar",
			Capacity:   0,
			Durability: 0,
			Flavor:     -2,
			Texture:    2,
			Calories:   1,
		},
	}

	// Part 1
	best := ProcessBestCookieScore(ingredients)
	fmt.Printf("Best: %d\n", best)
}

func ProcessBestCookieScore(ingredients []*Ingredient) int64 {

	var bestScore int64
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			for c := 0; c <= 100; c++ {
				for d := 0; d <= 100; d++ {
					if a+b+c+d != 100 {
						continue // So much wasted looping
					}
					// Part 2 adds calories restriction
					score, calories := CalculateCookieScore(ingredients, a, b, c, d)
					if score > bestScore && calories == 500 {
						fmt.Printf("[%d] New best for %d %d %d %d\n", score, a, b, c, d)
						bestScore = score
					}
				}
			}
		}
	}

	return bestScore
}

func CalculateCookieScore(ingredients []*Ingredient, a int, b int, c int, d int) (int64, int64) {

	capacity := ingredients[0].Capacity*int64(a) + ingredients[1].Capacity*int64(b) +
		ingredients[2].Capacity*int64(c) + ingredients[3].Capacity*int64(d)
	durability := ingredients[0].Durability*int64(a) + ingredients[1].Durability*int64(b) +
		ingredients[2].Durability*int64(c) + ingredients[3].Durability*int64(d)
	flavor := ingredients[0].Flavor*int64(a) + ingredients[1].Flavor*int64(b) +
		ingredients[2].Flavor*int64(c) + ingredients[3].Flavor*int64(d)
	texture := ingredients[0].Texture*int64(a) + ingredients[1].Texture*int64(b) +
		ingredients[2].Texture*int64(c) + ingredients[3].Texture*int64(d)
	calories := ingredients[0].Calories*int64(a) + ingredients[1].Calories*int64(b) +
		ingredients[2].Calories*int64(c) + ingredients[3].Calories*int64(d)

	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0, 0
	}

	return capacity * durability * flavor * texture, calories
}
