package odev2

import (
	"fmt"
	"go_pointers_tasks/odev1"
	"sort"
)

//Odev2_1

func Odev2_1() {
	odev1.Space_between_tasks("Odev2_1")

	a, b := 5, 10
	p_a, p_b := &a, &b
	fmt.Printf("Önce: a=%d, a'nın adresi: %d, b=%d, b'nin adresi: %d \n", a, &a, b, &b)
	swap(p_a, p_b)
	fmt.Printf("Sonra: a=%d, a'nın adresi: %d, b=%d, b'nin adresi: %d \n", a, &a, b, &b)

}
func swap(i, j *int) {
	*i, *j = *j, *i

}

//Odev2_2

func Odev2_2() {
	odev1.Space_between_tasks("Odev2_2")

	var result int
	p_result := &result
	topla(15, 7, p_result)
	fmt.Printf("Toplam: %d, Adres: %d\n ", result, p_result)

	cikar(20, 8, p_result)
	fmt.Printf("Fark: %d, Adres: %d\n ", result, p_result)

}

func topla(a, b int, sonuc *int) {
	*sonuc = a + b

}

func cikar(a, b int, sonuc *int) {
	*sonuc = a - b
}

//Odev2_3

func Odev2_3() {
	odev1.Space_between_tasks("Odev2_3")
	numbers := []int{45, 12, 89, 3, 67, 23}
	var min, max int

	minMax(numbers, &min, &max)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

}
func minMax(sayilar []int, minPtr, maxPtr *int) {
	sort.Ints(sayilar)
	*minPtr = sayilar[0]
	*maxPtr = sayilar[len(sayilar)-1]
}
