package server

import "github.com/dashotv/codex/internal/plex"

func setupPlex(s *Server) error {
	p := plex.New(&plex.ClientOptions{
		URL:   s.Config.PlexURL,
		Token: s.Config.PlexToken,
		Debug: false,
	})

	s.Plex = p
	return nil
}
