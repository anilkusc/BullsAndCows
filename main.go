package main

var (
	store = make(map[int]map[int]int)
)

func main() {

	a := App{}
	a.Init("test.db")
	a.Run(":8080")

}
