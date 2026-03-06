package main

import (
	routers "GO-ECOMMERCE-BACKEND-API/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}
