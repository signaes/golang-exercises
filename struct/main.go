package main

import "fmt"

func main() {
	type state struct {
		name   string
		region string
		abbr   string
	}

	df := state{
		name:   "Distrito Federal",
		region: "Centro Oeste",
		abbr:   "DF",
	}

	sc := state{
		name:   "Santa Catarina",
		region: "Sul",
		abbr:   "SC",
	}

	type animal struct {
		name           string
		from           string
		maxHeight      float64
		naturalHabitat string
	}

	type mammal struct {
		animal
		gestationDays [2]int
	}

	panda := mammal{
		animal: animal{
			name:           "Panda",
			from:           "China",
			maxHeight:      1.6,
			naturalHabitat: "Forest",
		},
		gestationDays: [2]int{95, 160},
	}

	fmt.Println(df.name)
	fmt.Println(sc.name)

	fmt.Println(panda)
	fmt.Println(panda.from)
	fmt.Println(panda.animal.name)
}
