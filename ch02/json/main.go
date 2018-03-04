package main

import (
	"encoding/json"
	"os"
)

func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}

// {
//     "example": "encoding/json",
//     "hello": "world"
// }
