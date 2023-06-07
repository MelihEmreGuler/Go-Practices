package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var map1 map[string]Vertex

//var map2 map[int]int

func main() {
	map1 = make(map[string]Vertex)
	map1["My Home"] = Vertex{
		40.90428998993353, 31.180442550293332,
	}
	fmt.Println(map1["My Home"])
}
