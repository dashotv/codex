package server

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/term"

	"github.com/dashotv/codex/internal/plex"
	"github.com/dashotv/minion"
)

type Server struct {
	Config *Config

	Router *echo.Echo
	Logger *zap.SugaredLogger
	Plex   *plex.Client

	db *connection
	bg *minion.Minion

	// Services
	file *fileService
}

func New() (*Server, error) {
	logger := setupLogger()

	s := &Server{
		Logger: logger,
	}

	if err := setupConfig(s); err != nil {
		return nil, err
	}
	if err := setupDatabase(s); err != nil {
		return nil, err
	}
	if err := setupPlex(s); err != nil {
		return nil, err
	}
	if err := setupWorkers(s); err != nil {
		return nil, err
	}

	setupRouter(s)

	file := &fileService{log: logger.Named("services.file"), bg: s.bg, db: s.db}
	job := &jobService{log: logger.Named("services.file"), bg: s.bg, db: s.db}

	g := s.Router.Group("/api")
	RegisterFileService(g, file)
	RegisterJobService(g, job)

	return s, nil
}

func (s *Server) Start() error {
	startWorkers(context.Background(), s)

	count, err := s.db.File.Query().Count()
	if err != nil {
		return err
	}
	s.Logger.Debugf("managing %d files", count)

	return s.Router.Start(":" + s.Config.Port)
}

func setupLogger() *zap.SugaredLogger {
	isTTY := term.IsTerminal(int(os.Stderr.Fd()))
	verbosity := 1
	logStdoutWriter := zapcore.Lock(os.Stderr)
	log := zap.New(zapcore.NewCore(logging.NewEncoder(verbosity, isTTY), logStdoutWriter, zapcore.DebugLevel))
	return log.Named("codex").Sugar()
}

var pageSize = 100

func reqLimitSkip(req *Request) (int, int) {
	limit := pageSize
	if req.Limit > 0 {
		limit = req.Limit
	}

	return limit, req.Skip
}
