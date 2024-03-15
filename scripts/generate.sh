#!/usr/bin/env bash
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
BASE_DIR=$(cd "$SCRIPT_DIR/.." && pwd)

SERVER="$BASE_DIR/internal/server"
TEMPLATES="$BASE_DIR/internal/templates"
DEFINITIONS="$BASE_DIR/internal/definitions"

# Generate the code
mkdir -p "$SERVER"
oto -template "$TEMPLATES/echo.go.plush" \
  -out "$SERVER/server.gen.go" \
  -ignore Ignorer \
  -pkg server \
  "$DEFINITIONS"
oto -template "$TEMPLATES/models.go.plush" \
  -out "$SERVER/models.gen.go" \
  -ignore Ignorer \
  -pkg server \
  "$DEFINITIONS"

goimports -w -local github.com/dashotv "$SERVER"/*.go
