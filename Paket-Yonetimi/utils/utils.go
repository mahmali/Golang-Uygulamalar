package utils

func FullName(ad, soyad string) (string, int) { //ilk parantezler giriş veri alamaktadır. 2. parantez ise geri döndürülecek olan parametrelerin tipleri veirir
	full := ad + " " + soyad
	length := len(full)
	return full, length
}
func FullNameLen(isim, soyisim string) (full string, length int) {
	full = isim + " " + soyisim
	length = len(full)
	return
}
