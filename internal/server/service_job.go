package server

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/dashotv/minion"
)

var workers = map[string]minion.Payload{
	"FileWalk":  &FileWalk{},
	"FileMatch": &FileMatch{},
}

type jobService struct {
	db  *connection
	log *zap.SugaredLogger
	bg  *minion.Minion
}

func (s *jobService) Index(c echo.Context, req *Request) (*JobsResponse, error) {
	limit, skip := reqLimitSkip(req)

	count, err := s.db.Job.Query().Count()
	if err != nil {
		return nil, err
	}

	list, err := s.db.Job.Query().Limit(limit).Skip(skip).Desc("created_at").Run()
	if err != nil {
		return nil, err
	}

	return &JobsResponse{Total: count, Results: list}, nil
}

func (s *jobService) Create(c echo.Context, req *Request) (*JobResponse, error) {
	s.log.Infof("job create: %+v", req)
	id := req.ID
	j, ok := workers[id]
	if !ok {
		return nil, errors.New("unknown job:" + id)
	}

	if err := s.bg.Enqueue(j); err != nil {
		return nil, err
	}

	return &JobResponse{Job: &Job{Kind: id}}, nil
}

func (s *jobService) Update(c echo.Context, req *Job) (*JobResponse, error) {
	return nil, errors.New("not implemented")
}
