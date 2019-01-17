package main

import "fmt"

var categories = map[int]string{
	0: "substantive",
	1: "verb",
	2: "adjective",
	3: "adverb",
	4: "conjunction",
}

var plurality = map[int]string{
	0: "singular",
	1: "dual",
	2: "plural",
}

var gender = map[string]bool{
	"m": true,
	"f": false,
}

func getKey(s string) (int, bool) {
	i := -1

	for k, v := range categories {
		if s == v {
			i = k
		}
	}

	return i, i == -1
}

type word struct {
	gender              bool
	category, plurality int
	written, definition string
}

type term struct {
	word
	synonyms []term
}

func main() {
	defer bye()
	fmt.Println(add(20, 30))

	substantive, _ := getKey("substantive")
	singular, _ := getKey("singular")

	conceito := term{
		word: word{
			written:    "Conceito",
			definition: "",
			category:   substantive,
			plurality:  singular,
			gender:     gender["m"],
		},
	}

	conceito.addSynonyms("Ideia", "Modelo").removeSynonym("Modelo").write()

	fmt.Println(conceito.synonyms)

	if adverb, err := getKey("adverb"); !err {
		fmt.Println("Adverb category is", adverb)
	}
}

// func (r receiver) identifier(parameters) (returns) { body }

func add(a float64, b float64) float64 {
	return a + b
}

func bye() {
	fmt.Println("See you")
}

func (t term) write() {
	fmt.Print(t.written)

	for _, w := range t.synonyms {
		fmt.Println(" has synonyms :", w.written)
	}
}

func (t *term) addSynonyms(words ...string) *term {
	for _, w := range words {
		newTerm := term{
			word: word{
				written: w,
			},
		}
		t.synonyms = append(t.synonyms, newTerm)
	}

	return t
}

func (t *term) removeSynonym(word string) *term {
	var newSynonyms []term

	for _, w := range t.synonyms {
		if w.written != word {
			newSynonyms = append(newSynonyms, w)
		}
	}

	t.synonyms = newSynonyms

	return t
}
