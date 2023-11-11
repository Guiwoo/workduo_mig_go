package wd_streams

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type ApiServer struct{}

func (a *ApiServer) Run(ctx context.Context) <-chan error {
	stream := make(chan error)
	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
				stream <- fmt.Errorf("api server error")
				return
			}
		}
	}()
	return stream
}

func Test_Stream(t *testing.T) {
	stream := &Stream{}
	stream.Add((&ApiServer{}).Run)
	stream.Run()
}
