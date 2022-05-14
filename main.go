package main

import "zshf.private/initialize"

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Redis()
	initialize.Router()
}
