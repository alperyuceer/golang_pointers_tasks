package odev1

import "fmt"

func Odev1_1() {
	Space_between_tasks("Odev1_1")
	sayi := 42
	p_sayi := &sayi
	fmt.Println("Sayı pointerının adresi: ", p_sayi)
	fmt.Println("Sayı pointerının değeri: ", *p_sayi)

}

func Odev1_2() {
	Space_between_tasks("Odev1_2")

	x := 10
	p_x := &x
	fmt.Println("x sayısının adresi: ", p_x)
	*p_x = 25
	fmt.Println("x sayısının yeni değeri:", x)
	fmt.Println("x sayısının adresi: ", p_x)

}

func Odev1_3() {
	Space_between_tasks("Odev1_3")
	sayi := 7
	p_sayi := &sayi
	fmt.Println("İki kat yapmadan önceki adres: ", p_sayi)
	fmt.Println("İki kat yapmadan önceki değer: ", *p_sayi)
	ikiKatYap(p_sayi)
	fmt.Println(sayi) //14 olmalı
	fmt.Println("İki kat yaptıktan sonraki adres: ", p_sayi)
	fmt.Println("İki kat yaptıktan sonraki değer: ", *p_sayi)

}
func ikiKatYap(sayi *int) {
	*sayi = *sayi * 2
}

func Space_between_tasks(odevNumarasi string) {
	fmt.Printf("\n\n ****************** %v ******************\n", odevNumarasi)

}
