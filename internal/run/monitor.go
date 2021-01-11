package run

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/frankgreco/aviation/internal/log"
)

type process struct {
	ctx     context.Context
	cancel  context.CancelFunc
	signals chan os.Signal
	logger  log.Logger
}

func NewMonitor(parent context.Context, logger log.Logger) Runnable {
	ctx, cancel := context.WithCancel(parent)
	signals := make(chan os.Signal, 2)

	return &process{
		ctx:     ctx,
		cancel:  cancel,
		signals: signals,
		logger:  logger,
	}
}

func (p *process) Run() error {
	p.logger.Info("starting parent process monitor")

	signal.Notify(p.signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		firstSignal := <-p.signals
		p.logger.Info(fmt.Sprintf("received %s signal", firstSignal.String()))
		p.cancel()
		secondSignal := <-p.signals
		p.logger.Info(fmt.Sprintf("received %s signal", secondSignal.String()))
		os.Exit(1) // second signal. Exit directly.
	}()

	<-p.ctx.Done()
	return nil
}

func (p *process) Close(error) error {
	p.logger.Info("closing parent process monitor")
	p.cancel()
	return nil
}
