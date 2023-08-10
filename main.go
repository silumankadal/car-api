package main

import "gin/routers"

func main() {
	var PORT = ":8000"

	routers.StartServer().Run(PORT)
}