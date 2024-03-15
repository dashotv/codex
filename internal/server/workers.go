package server

import (
	"context"

	"github.com/pkg/errors"

	"github.com/dashotv/minion"
)

func startWorkers(ctx context.Context, s *Server) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		// s.Logger.Infof("starting workers (%d)...", s.Config.MinionConcurrency)
		if err := s.bg.Start(); err != nil {
			s.Logger.Errorf("starting workers: %s", err)
			return
		}

		<-ctx.Done()
	}()

	return nil
}

func setupWorkers(s *Server) error {
	ctx := context.Background()

	dbname := s.Config.Name + "_development"
	if s.Config.Production {
		dbname = s.Config.Name + "_production"
	}

	mcfg := &minion.Config{
		Logger:      s.Logger.Named("minion"),
		Debug:       s.Config.MinionDebug,
		Concurrency: s.Config.MinionConcurrency,
		BufferSize:  s.Config.MinionBufferSize,
		DatabaseURI: s.Config.Mongo,
		Database:    dbname,
		Collection:  "job",
	}

	ctx = context.WithValue(ctx, "server", s)

	m, err := minion.New(ctx, mcfg)
	if err != nil {
		return errors.Wrap(err, "creating minion")
	}

	// add something like the below line in app.Start() (before the workers are
	// started) to subscribe to job notifications.
	// minion sends notifications as jobs are processed and change status
	// m.Subscribe(app.MinionNotification)
	// an example of the subscription function and the basic setup instructions
	// are included at the end of this file.

	if err := minion.Register(m, &FileWalk{}); err != nil {
		return errors.Wrap(err, "registering worker: scrape_page (ScrapePage)")
	}

	if s.Config.Production {
		if _, err := m.Schedule("0 */15 * * * *", &FileWalk{}); err != nil {
			return errors.Wrap(err, "scheduling worker: scrape_pages (ScrapePages)")
		}
	}

	s.bg = m
	return nil
}

func getServer(ctx context.Context) *Server {
	return ctx.Value("server").(*Server)
}

// run the following commands to create the events channel and add the necessary models.
//
// > golem add event jobs event id job:*Minion
// > golem add model minion_attempt --struct started_at:time.Time duration:float64 status error 'stacktrace:[]string'
// > golem add model minion queue kind args status 'attempts:[]*MinionAttempt'
//
// then add a Connection configuration that points to the same database connection information
// as the minion database.

// // This allows you to notify other services as jobs change status.
//func (a *Application) MinionNotification(n *minion.Notification) {
//	if n.JobID == "-" {
//		return
//	}
//
//	j := &Minion{}
//	err := app.DB.Minion.Find(n.JobID, j)
//	if err != nil {
//		log.Errorf("finding job: %s", err)
//		return
//	}
//
//	if n.Event == "job:created" {
//		events.Send("runic.jobs", &EventJob{"created", j.ID.Hex(), j})
//		return
//	}
//	events.Send("runic.jobs", &EventJob{"updated", j.ID.Hex(), j})
//}
