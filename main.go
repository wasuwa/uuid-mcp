package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	uuidserver "github.com/wasuwa/uuid-mcp/server"
)

func main() {
	s := uuidserver.New()
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("サーバーエラー: %v\n", err)
	}
}
