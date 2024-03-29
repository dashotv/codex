package server

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/dashotv/minion"
)

var walking uint32

type FileWalk struct {
	minion.WorkerDefaults[*FileWalk]
}

func (j *FileWalk) Kind() string { return "file_walk" }
func (j *FileWalk) Timeout(job *minion.Job[*FileWalk]) time.Duration {
	return 60 * time.Minute
}
func (j *FileWalk) Work(ctx context.Context, job *minion.Job[*FileWalk]) error {
	s := getServer(ctx)
	l := s.Logger.Named("file_walk")
	if !atomic.CompareAndSwapUint32(&walking, 0, 1) {
		l.Warnf("walkFiles: already running")
		return fmt.Errorf("already running")
	}
	defer atomic.StoreUint32(&walking, 0)

	libs, err := s.Plex.GetLibraries()
	if err != nil {
		l.Errorw("libs", "error", err)
		return fmt.Errorf("getting libraries: %w", err)
	}

	w := newWalker(s.db, l.Named("walker"), libs)
	if err := w.Walk(); err != nil {
		l.Errorw("walk", "error", err)
		return fmt.Errorf("walking: %w", err)
	}

	s.bg.Enqueue(&FileMatch{})
	return nil
}

type FileMatch struct {
	minion.WorkerDefaults[*FileMatch]
}

func (j *FileMatch) Kind() string { return "file_match" }
func (j *FileMatch) Timeout(job *minion.Job[*FileMatch]) time.Duration {
	return 60 * time.Minute
}
func (j *FileMatch) Work(ctx context.Context, job *minion.Job[*FileMatch]) error {
	// 	s := getServer(ctx)
	// 	l := s.Logger.Named("file_match")
	// 	q := s.db.File.Query().In("medium_id", bson.A{nil, "", primitive.NilObjectID})
	//
	// 	total, err := q.Count()
	// 	if err != nil {
	// 		l.Errorw("total", "error", err)
	// 		return fmt.Errorf("counting: %w", err)
	// 	}
	// 	l.Debugf("total: %d", total)
	//
	// 	skip := 0
	// 	limit := 25
	// 	for skip < int(total) {
	// 		list, err := q.Limit(limit).Skip(skip).Run()
	// 		if err != nil {
	// 			l.Errorw("query", "error", err)
	// 			return fmt.Errorf("querying: %w", err)
	// 		}
	//
	// 		for _, f := range list {
	// 			// l.Debugf("match: %s", f.Path)
	// 			m, err := s.db.MediumByFile(f)
	// 			if err != nil {
	// 				l.Warnw("medium", "error", err)
	// 				continue
	// 			}
	// 			if m == nil {
	// 				l.Warnw("medium", "error", "not found", "file", f.Path)
	// 				continue
	// 			}
	//
	// 			// l.Debugf("found: %s", m.Title)
	// 			f.MediumId = m.ID
	// 			if err := app.DB.File.Save(f); err != nil {
	// 				l.Errorf("save", "error", err)
	// 				return fmt.Errorf("saving: %w", err)
	// 			}
	// 		}
	// 		skip += limit
	// 	}

	return nil
}
