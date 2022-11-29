package main

import "fmt"

func main() {
	a := make([]int, 5) //type, length, capacity (icindeki elemanlar sifir olan dizi yapar)
	printSlice("a", a)

	b := make([]int, 0, 5) // icinde eleman olmayan ama temel dizinin eleman sayisi 5 olan bir dizi yarattik
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:] //slice lere baslangic degeri verince neden kapasitesi azaliyor? ust deger verdigimizde neden cap azalmiyor?
	printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%v: %v, len: %v cap: %v,\n", s, x, len(x), cap(x))
}
