package main

type IlceBilgi struct {
	Ulke     string     `json:"ulke"`
	Il       string     `json:"il"`
	Ilce     string     `json:"ilce"`
	UlkeId   int        `json:"ulke_id"`
	IlId     int        `json:"il_id"`
	IlceId   int        `json:"ilce_id"`
	Vakitler []Imsakiye `json:"vakitler"`
}
