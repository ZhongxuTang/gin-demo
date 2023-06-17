package main

import (
	"github.com/lemonsoul/gin-demo/config"
	"github.com/lemonsoul/gin-demo/routers"
)

func main() {
	config.InitDB()
	routers.InitRouter()
}
