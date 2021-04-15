package main

import "fmt"

type marka string

const (
	audi     marka = "audi"
	mercedes marka = "mercedes"
	fiat     marka = "fiat"
	google   marka = "google"
)

func markaYazdır(yazdır marka) {
	fmt.Println(yazdır)
}

func main() {
	markaYazdır(audi)
	markaYazdır(fiat)

}
