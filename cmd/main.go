package main

import (
	"github.com/bkohler93/bootdev-blog-aggregator/internal/app/bloggo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bloggo.RunApp()
}
