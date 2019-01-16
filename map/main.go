package main

import "fmt"

func main() {
	m := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	fmt.Println(m)

	v, ok := m[13]

	fmt.Println(v, ok)

	dates := map[int]string{}

	dates[27082015] = "Ilyas"
	dates[30121985] = "Hyun Jin"

	dates[0000000] = "Init"

	delete(dates, 0000000)

	for k, v := range dates {
		fmt.Println(k, v)
	}
}
