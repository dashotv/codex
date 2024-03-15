package main

import (
	"log"

	_ "github.com/dotenv-org/godotenvvault/autoload"

	"github.com/dashotv/codex/internal/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

// TODO: consider using Run function pattern (passing in args, env, etc to make it easier to test)
