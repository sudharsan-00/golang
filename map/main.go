package main

import "fmt"

func main() {
	m := make(map[string][]string)

	m["chris_hemsworth"] = []string{"thor", "extraction", "ragnarok"}
	m["tony_stark"] = []string{"endgame", "ironman", "avengers"}
	m["James_bond"] = []string{"goldfinger", "octopussy", "thunderball"}

	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
	fmt.Println("---------record deleted-----------")

	delete(m, "James_bond")

	fmt.Println("---------record deleted-----------")

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
