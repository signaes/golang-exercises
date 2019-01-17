package main

import "fmt"

type term struct {
	word     string
	synonyms []string
}

func main() {
	fmt.Println(add(20, 30))

	conceito := term{
		word: "Conceito",
	}

	conceito.addSynonyms("Ideia", "Modelo").removeSynonym("Modelo")

	fmt.Println(conceito.synonyms)
}

// func (r receiver) identifier(parameters) (returns) { body }

func add(a float64, b float64) float64 {
	return a + b
}

func (t *term) addSynonyms(words ...string) *term {
	for _, w := range words {
		t.synonyms = append(t.synonyms, w)
	}

	return t
}

func (t *term) removeSynonym(word string) *term {
	var newSynonyms []string

	for _, w := range t.synonyms {
		if w != word {
			newSynonyms = append(newSynonyms, w)
		}
	}

	t.synonyms = newSynonyms

	return t
}
