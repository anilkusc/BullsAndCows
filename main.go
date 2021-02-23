package main

func main() {

	a := App{}
	a.Init("test.db")
	a.Run(":8080")

}
