package pipeline

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type Pipeline func(ctx context.Context) <-chan error

type Stream struct {
	pipelines []Pipeline
}

func (s *Stream) Add(pipeline Pipeline) {
	s.pipelines = append(s.pipelines, pipeline)
}

func (s *Stream) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	signalStream := make(chan os.Signal)
	signal.Notify(signalStream, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		for _, line := range s.pipelines {
			go func(p Pipeline) {
				defer cancel()
				for err := range p(ctx) {
					logrus.WithError(err).Error("pipline error")
					cancel()
				}
			}(line)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-signalStream:
			logrus.Error("got interrupt signal canceling context")
			cancel()
		}
	}
}
