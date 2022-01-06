package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var myStoredVariable interface{}
	myJsonString := `[{"one":"two"},{"three":"four"}]`

	// `&myStoredVariable` is the address of the variable we want to store our
	// parsed data in
	json.Unmarshal([]byte(myJsonString), &myStoredVariable)
	fmt.Printf("%+v\n", myStoredVariable)
	fmt.Printf("Type:%T\n", myStoredVariable)
}
