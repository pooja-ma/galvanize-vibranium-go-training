package main

import "fmt"

func main() {

	// Maps are used to store data values in key:value pairs.
	// Each element in a map is a key:value pair.

	var capitalMap map[string]string

	// The make function returns a map of the given type, initialized and ready for use.
	capitalMap1 := make(map[string]string)

	fmt.Println(capitalMap)
	fmt.Println(capitalMap1)

	var capiTalMap = map[string]string{"height": "tall"}

	capiTalMap["weight"] = "80kg"
	fmt.Println(capiTalMap)

	capitalMap = make(map[string]string)
	/* insert key-value pairs in the map*/
	capitalMap["Karnataka"] = "Bangalore"
	capitalMap["Andra Pradesh"] = "Hydarabad"

	capitalMap1["Karnataka"] = "Bangalore"
	capitalMap1["Andra Pradesh"] = "Hydarabad"

	/* print map using keys*/
	for capital := range capitalMap {
		fmt.Println("Capital of", capital, "is", capitalMap[capital])
	}
}
