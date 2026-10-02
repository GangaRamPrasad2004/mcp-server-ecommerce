#  MCP E-Commerce Server

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![MCP](https://img.shields.io/badge/MCP-2025--06--18-blueviolet?style=for-the-badge)](https://modelcontextprotocol.io/)
[![JSON--RPC](https://img.shields.io/badge/JSON--RPC-2.0-orange?style=for-the-badge)](https://www.jsonrpc.org/specification)


A **production-grade Model Context Protocol (MCP) server** built from scratch in Go that enables AI assistants (Claude, Cursor, GPT-based hosts) to interact with an E-Commerce platform — browsing products, managing shopping carts, and placing orders — all through natural language.

> **No SDK. No framework. Pure Go implementation of MCP + JSON-RPC 2.0 over stdio.**

---



##  Features

- **Full MCP Protocol Implementation** — Handshake (`initialize`/`initialized`), tool discovery (`tools/list`), and tool execution (`tools/call`)
- **Custom JSON-RPC 2.0 Engine** — Built from scratch with spec-compliant error codes, notification support, and stdio transport
-  **6 E-Commerce Tools** — Product browsing, search with filters, cart management, and order placement
-  **JWT Authentication** — Secure Bearer token injection for user-scoped actions (cart & orders)
-  **Structured Logging** — Logrus-powered JSON logs to `stderr`, keeping `stdout` clean for protocol communication
-  **Clean Architecture** — Layered design with clear separation of concerns (transport → protocol → registry → tools → client)
-  **Zero External MCP Dependencies** — No third-party MCP SDK; the entire protocol stack is hand-written

---

##  Architecture (from modeelcontextprotocol.io)

```
┌─────────────────────────────────────────────────────────┐
│                      MCP Client                         │
│          (Claude Desktop / Cursor / AI Host)            │
└────────────────────────┬────────────────────────────────┘
                         │ stdio (JSON-RPC 2.0)
                         ▼
┌─────────────────────────────────────────────────────────┐
│               JSON-RPC Transport Layer                  │
│             (internal/jsonRPC/server.go)                │
│    Reads os.Stdin ─── Writes os.Stdout ─── Logs Stderr  │
└────────────────────────┬────────────────────────────────┘
                         ▼
┌─────────────────────────────────────────────────────────┐
│                 MCP Protocol Layer                      │
│              (internal/mcp/server.go)                   │
│   initialize │ initialized │ tools/list │ tools/call    │
└────────────────────────┬────────────────────────────────┘
                         ▼
┌─────────────────────────────────────────────────────────┐
│                  MCP Tool Registry                      │
│              (internal/mcp/registry.go)                 │
├─────────────┬──────────────────┬────────────────────────┤
│  Products   │      Cart        │       Orders           │
│  Toolset    │    Toolset       │      Toolset           │
└──────┬──────┴────────┬─────────┴──────────┬─────────────┘
       └───────────────┼────────────────────┘
                       ▼
┌─────────────────────────────────────────────────────────┐
│                  REST Client Layer                      │
│             (internal/client/http.go)                   │
│         Resty HTTP │ Bearer Auth │ Logging Hooks        │
└────────────────────────┬────────────────────────────────┘
                         │ HTTP/REST
                         ▼
┌─────────────────────────────────────────────────────────┐
│            External E-Commerce REST API                 │
│            (default: http://localhost:8080)              │
└─────────────────────────────────────────────────────────┘
```

---

##  Tech Stack

| Layer | Technology | Purpose |
|:------|:-----------|:--------|
| **Language** | ![Go](https://img.shields.io/badge/Go_1.26-00ADD8?logo=go&logoColor=white) | Systems-level performance, strong concurrency, static typing |
| **Protocol** | ![MCP](https://img.shields.io/badge/MCP-Model_Context_Protocol-blueviolet) | Standardized AI ↔ Tool communication (Anthropic's open protocol) |
| **Transport** | ![JSON-RPC](https://img.shields.io/badge/JSON--RPC_2.0-orange) | Lightweight RPC over stdio, hand-implemented to spec |
| **HTTP Client** | ![Resty](https://img.shields.io/badge/Resty_v2-REST_Client-blue) | Fluent HTTP client with middleware hooks for external API calls |
| **Config** | ![godotenv](https://img.shields.io/badge/godotenv-Environment_Config-yellow) | `.env` file loading for local development |
| **Logging** | ![Logrus](https://img.shields.io/badge/Logrus-Structured_Logging-red) | JSON-formatted structured logging to stderr |
| **Auth** | ![JWT](https://img.shields.io/badge/JWT-Bearer_Token-black) | Token-based authentication for user-scoped operations |

---

##  Project Structure

```
mcp-server-Ecommerce/
├── cmd/
│   └── server/
│       └── main.go              # Entry point — bootstraps all layers
├── config/
│   └── config.go                # Environment & .env configuration loader
├── internal/
│   ├── client/
│   │   └── http.go              # Resty-based REST client with auth support
│   ├── jsonRPC/
│   │   ├── errors.go            # JSON-RPC 2.0 standard error codes & helpers
│   │   ├── server.go            # Stdio JSON-RPC transport engine
│   │   └── types.go             # Request/Response/Error type definitions
│   ├── mcp/
│   │   ├── helpers.go           # MCP tool result utility functions
│   │   ├── registry.go          # Tool catalog & handler registry
│   │   ├── server.go            # MCP protocol handler (initialize, tools/*)
│   │   └── types.go             # MCP specification type definitions
│   └── tools/
│       ├── cart/
│       │   ├── cart.go           # add_to_cart, view_cart tool handlers
│       │   └── data.go          # Cart API response DTOs
│       ├── orders/
│       │   ├── orders.go        # place_order tool handler
│       │   └── data.go          # Order API response DTOs
│       └── products/
│           ├── products.go      # list, search, get_details tool handlers
│           └── data.go          # Product API response DTOs
├── .env                         # Environment variables (gitignored)
├── .gitignore
├── go.mod
└── go.sum
```

---

##  Getting Started

### Prerequisites

- **Go 1.26+** — [Install Go](https://go.dev/doc/install)
- **An E-Commerce REST API** running (or any compatible backend on `localhost:8080`)

### Installation

```bash
# Clone the repository
git clone https://github.com/GangaRamPrasad2004/mcp-server-Ecommerce.git
cd mcp-server-Ecommerce

# Download dependencies
go mod download
```

### Configuration

Create a `.env` file in the project root:

```env
# Base URL of the E-Commerce REST API
API_URL=http://localhost:8080 ->https://gothub.com/GangaRamPrasad2004/ECommerce-Platform

# JWT Bearer token for authenticated operations (cart, orders)
AUTH_TOKEN=your_jwt_token_here

# Log level: debug | info | warn | error
LOG_LEVEL=debug

# Transport mode (currently supports: stdio)
TRANSPORT=stdio
```

### how to run the Server

```bash
# Build and run
go build -o mcp-server ./cmd/server
./mcp-server

# Or run directly
go run ./cmd/server
```

The server will start listening for JSON-RPC messages on `stdin` and respond on `stdout`.

---


##  Integration with AI Clients

### Claude Desktop

Add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "ecommerce": {
      "command": "/path/to/mcp-server",
      "env": {
        "API_URL": "http://localhost:8080",
        "AUTH_TOKEN": "your_jwt_token"
      }
    }
  }
}
```




