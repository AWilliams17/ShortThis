package main

import (
	"github.com/awilliams17/ShortThis/api"
)

func main() {
	e := api.CreateRouter()

	e.Start(":8000")
}
