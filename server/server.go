package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// New MCPサーバーを初期化
func New() *server.MCPServer {
	s := server.NewMCPServer("uuid", "1.0.0")
	tool := mcp.NewTool(
		"generate_uuid_v4",
		mcp.WithDescription("UUID v4を生成する"),
	)
	s.AddTool(tool, generateUUIDV4Handler)
	return s
}

// generateUUIDV4Handler UUIDv4を生成する
func generateUUIDV4Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	u := uuid.New()
	return mcp.NewToolResultText(u.String()), nil
}
