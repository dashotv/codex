package server

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/dashotv/minion"
)

type fileService struct {
	log *zap.SugaredLogger
	bg  *minion.Minion
	db  *connection
}

func (s *fileService) Index(c echo.Context, req *IndexRequest) (*FilesResponse, error) {
	count, err := s.db.File.Query().Count()
	if err != nil {
		return nil, err
	}

	files, err := s.db.File.Query().Desc("modified_at").Run()
	if err != nil {
		return nil, err
	}

	return &FilesResponse{Count: count, Result: files}, nil
}

func (s *fileService) Walk(c echo.Context, req *KeyRequest) (*EmptyResponse, error) {
	s.bg.Enqueue(&FileWalk{})
	return &EmptyResponse{}, nil
}
