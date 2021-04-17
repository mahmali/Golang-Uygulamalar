package main

import (
	"fmt"
	"strconv"
)

func main() {
	ferr := new(Ferrari)
	ferr.Model = "f11"
	ferr.Marka = "ferrari"
	ferr.Renk = "kırmızı"
	ferr.Hiz = 300
	ferr.Fiyat = 300
	ferr.ozel = true

	fmt.Println(ferr.Information())

}

//interface
type arabaface interface {
	Run() bool
	Stop() bool
	Information() string
}
type otomabil struct {
	Marka string
	Model string
	Renk  string
	Hiz   int
	Fiyat float64
}
type ozelurun struct {
	ozel bool
}

//ferrari
type Ferrari struct {
	otomabil //kalıtım olarak alınmış oldu
	ozelurun
}

func (_ Ferrari) Run() bool {
	return true
}
func (_ Ferrari) Stop() bool {
	return true
}
func (x *Ferrari) Information() string {
	ret := x.Marka + " " + x.Model + " " + x.Renk + "\n" + strconv.Itoa(x.Hiz) + "\n" + strconv.FormatFloat(x.Fiyat, 'g', -1, 64)
	add := "Evet"
	if x.ozel {
		ret += "\n" + "\t" + "ozellik:" + add
	}
	return ret
}

//toros
type Toros struct {
	otomabil //kalıtım olarak alınmış oldu
	ozelurun
}

func (_ Toros) Run() bool {
	return true
}
func (_ Toros) Stop() bool {
	return true
}
func (x *Toros) Information() string {
	ret := x.Marka + " " + x.Model + " " + x.Renk + "\n" + strconv.Itoa(x.Hiz) + "\n" + strconv.FormatFloat(x.Fiyat, 'E', -1, 64)
	add := "Evet"
	if x.ozel {
		ret += "\n" + "\t" + "ozellik:" + add
	}
	return ret
}

//mercedes
type Mercedes struct {
	otomabil //kalıtım olarak alınmış oldu
	ozelurun
}

func (_ Mercedes) Run() bool {
	return true
}
func (_ Mercedes) Stop() bool {
	return true
}
func (x *Mercedes) Information() string {
	ret := x.Marka + " " + x.Model + " " + x.Renk + "\n" + strconv.Itoa(x.Hiz) + "\n" + strconv.FormatFloat(x.Fiyat, 'E', -1, 64)
	return ret
}
