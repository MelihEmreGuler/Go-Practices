package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var map1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var map2 = map[string]Vertex{ //her key value tanımlamamızda struct' in ismini belirtmemize gerek yok 
	"Bell Labs": {40.68433, -74.39967,},
	"Google" : {37.42202, -122.08408,},
}

func main() {
	fmt.Println(map1)
}
