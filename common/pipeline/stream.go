package pipeline

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Pipeline func(ctx context.Context) <-chan error

type Stream struct {
	pipelines []Pipeline
	sc        *sync.WaitGroup
}

func NewStream() *Stream {
	return &Stream{
		sc: &sync.WaitGroup{},
	}
}

func (s *Stream) Add(pipeline Pipeline) {
	s.pipelines = append(s.pipelines, pipeline)
}

func (s *Stream) Run() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	signalStream := make(chan os.Signal)
	signal.Notify(signalStream, os.Interrupt, os.Kill, syscall.SIGTERM)

	s.sc.Add(len(s.pipelines))

	go func() {
		for _, line := range s.pipelines {
			go func(p Pipeline) {
				defer cancel()
				defer s.sc.Done()
				for err = range p(ctx) {
					logrus.WithError(err).Error("pipline error")
					cancel()
				}
			}(line)
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				err = fmt.Errorf("context done")
			case <-signalStream:
				logrus.Error("got interrupt signal canceling context")
				err = fmt.Errorf("got interrupt signal canceling context")
				cancel()
			}
		}
	}()
	s.sc.Wait()
	return err
}
