package main

func main() {
	/*
		var a=10
		var b=11

		if a==b{
			fmt.Println("a ve b birlerine eşittir")
		} else if a<b{
			fmt.Println("a küçükktür b")
		} else{
			fmt.Println("b < a dır")
		}
	*/
	if foo := 10; foo == 1 { //if içinde nesne tanımlama ve nesne kullanma işlemi
		println("bar")
	} else {
		println("buz")
	}
	println(foo) //burada foo çalışmaz. çünki yukarıda if içinde tanımlanan ifade sadece if içinde kullanılır

}
