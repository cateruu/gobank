package main

func main() {
	server := NewApiServer(":8888")
	server.Run()
}