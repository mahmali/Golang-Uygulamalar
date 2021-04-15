package main

func main() {
	/*
		for i := 0; ; i++ {
			println("deger", i)
		}
	*/
	sum := 1
	for sum < 10 {
		sum += sum
		println(sum)
	}

}
