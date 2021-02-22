package main

func main() {

	a := App{}
	a.Init("test.db", "admin", "admin")
	a.Run(":8080")

}
