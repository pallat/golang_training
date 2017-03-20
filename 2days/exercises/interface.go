package main

type drinking interface {
	brand() string
}

func main() {
		
}

func favorite(d drinking) {
	fmt.Println(d.brand())
}
