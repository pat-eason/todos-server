package main

import (
	"fmt"
	"github.com/pateason/todo-server/internal/transport"
)

func main() {
	fmt.Println("Starting router")
	transport.StartRouter()
}
