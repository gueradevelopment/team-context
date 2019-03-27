package main

import (
	"net/http"

	"github.com/gueradevelopment/team-context/routers"
)

func main() {
	r := routers.GetRouter()
	http.ListenAndServe(":8080", r)
}
