package main

import (
	"fmt"
)

// bu şekilde genel degisken tanimlayabiliyoruz
var Version = "1.2.3"

const Pi = 3.14

type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE    Brand = "Google"
	APPLE     Brand = "Apple"
)

func PrintBrand(brand Brand) {
	fmt.Println(brand)
}

func main() {
	//degiskenler
	/*
		fmt.Println("selam")

		var name string = "golang"
		fmt.Println(name)

		var message = "merhaba"
		fmt.Println(message)

		var a, b, c int = 3, 4, 5
		fmt.Println(a, b, c)

		kisaTanimlama := 4
		fmt.Println(kisaTanimlama)
	*/

	//deneme
	/*
	   fmt.Println("selam")
	   	fmt.Println("The time is", time.Now())
	*/

	/*
		fmt.Println("selam")

		var name string = "golang"
		fmt.Println(name)
	*/

	//public(buyuk harfle baslarsa) private(kucuk harfle başlarsa)

	//fmt.Println(Version)

	//char bastırılırken int'e cast edilir
	/*
		M, E, L, I, H := 'M', 'E', 'L', 'I', 'H'

		fmt.Println(M, E, L, I, H)
	*/

	/*
		var myFloat float32 = 4
		fmt.Println(myFloat)

		var myComplex complex64 = 5.0

		fmt.Println(myComplex)

	*/

	/*
		//boş tanımlayıcı (alt cizgi) //tanımlamak zorunda olduğumuz ama ileride kullanmayacağımız bir değişken olduğunda kullanılır
		//bir daha çağırılamaz sadece boş geçer
		_, variable, _ := 100, 123, 1000
		fmt.Println(variable)
	*/

	//enum kullaniminin go karsiligi (gayet akillica)
	//PrintBrand(GOOGLE)

	//convert
	/*
		myString := "1"

		myNumber, _ := strconv.Atoi(myString)

		fmt.Println(myNumber)

		result := myNumber + 1

		fmt.Println(result)

	*/

	//Casting
	/*
		//float'in boyutu daha büyük olmasına ragmen int'ten float'a cast ederken otomatik olarak yerine atama yapamiyoruz
		//go dilinde her halukarda cast işlemini yazarak belirterek yapiyoruz
		var myInteger int = 5
		var castedFloat float64 = float64(myInteger)

		var myFloat float64 = 4.0
		var castedInt int = int(myFloat)

		fmt.Println(castedFloat, castedInt)
	*/

	//Space Holders
	/* string1, string2, string3 := "Melih", "Emre", "Guler"
	aNumber := 42

	Length, _ := fmt.Println(string1, string2, string3) //Printf fonksiyonu iki deger donuyor birincisi icindekinin lengt' i digeri hata olursa hata mesaji
	fmt.Println("Length: ", Length)

	fmt.Printf("My first name is %v \n", string1)                     // %v value olarak kullanılır virgülden sonraki degeri %v nin oldugu yere yapistirir. (printf ile kullanilir	)
	fmt.Printf("Value of a number as float: %f \n", float64(aNumber)) //float tipinde tutmak icin var
	fmt.Printf("Type of aNumber: %T \n", aNumber) // (type)
 */

	

}
