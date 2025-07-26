package odev3

import (
	"fmt"
	"go_pointers_tasks/odev1"
)

// Odev3_1
func Odev3_1() {
	odev1.Space_between_tasks("Odev3_1")

	kisi := Kisi{"Alper", 23, "alper@email.com"}
	p_kisi := &kisi
	fmt.Println("Önce: ", kisi)
	p_kisi.bilgileriGuncelle("Ahmet", 30)
	fmt.Println("Sonra: ", kisi)

}

type Kisi struct {
	ad   string
	yas  int
	mail string
}

func (k *Kisi) bilgileriGuncelle(newName string, newAge int) {
	k.ad = newName
	k.yas = newAge

}

// Odev3_2

type BankaHesabi struct {
	hesapNo string
	bakiye  float64
	sahibi  string
}

func Odev3_2() {
	odev1.Space_between_tasks("Odev3_2")

	hesap := BankaHesabi{"12345", 1000.0, "Mehmet"}
	p_hesap := &hesap
	fmt.Printf("En başta hesabın adresi: %p\n ", p_hesap)

	p_hesap.paraYatir(500) //500 yatırcak
	fmt.Printf("Para yatırıldıktan sonra: %2.f, adresi: %p\n", hesap.bakiye, p_hesap)

	p_hesap.paraCek(200) //200 cekilecek
	fmt.Printf("Para çekildikten sonra: %2.f, adresi: %p\n", hesap.bakiye, p_hesap)

}

func (h *BankaHesabi) paraYatir(miktar float64) {
	h.bakiye += miktar
}

func (h *BankaHesabi) paraCek(miktar float64) {
	if h.bakiye >= miktar {
		h.bakiye -= miktar
	} else {
		fmt.Println("Yetersiz bakiye")
	}

}

// Odev3_3
func Odev3_3() {
	odev1.Space_between_tasks("Odev3_3")

	ogrenci := Ogrenci{ad: "Alper", notlar: []int{}}
	p_ogrenci := &ogrenci
	p_ogrenci.notEkle(85)
	p_ogrenci.notEkle(92)
	p_ogrenci.notEkle(78)
	p_ogrenci.notEkle(88)

	p_ogrenci.ortalamaHesapla()

	fmt.Printf("Öğrenci: %s\n", ogrenci.ad)
	fmt.Printf("Notlar: %v\n", ogrenci.notlar)
	fmt.Printf("Ortalama: %.2f\n", ogrenci.ortalama)

}

type Ogrenci struct {
	ad       string
	notlar   []int
	ortalama float64
}

func (o *Ogrenci) notEkle(not int) {
	o.notlar = append(o.notlar, not)

}
func (o *Ogrenci) ortalamaHesapla() {
	var notlarinToplami float64
	for i := 0; i < len(o.notlar); i++ {
		notlarinToplami += float64(o.notlar[i])
	}
	o.ortalama = notlarinToplami / float64(len(o.notlar))

}
