package signals

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var onlyOneSignalHandler = make(chan struct{})
var logger = logrus.WithFields(logrus.Fields{"pkg": "signals"})

// SignalHandler establishes interrupt handling. Must only be called once.
func SignalHandler() (stopCh <-chan struct{}) {
	logger.Infof("called")

	// close(channel) can only be called once. Thus cause a panic if handler is called twice
	close(onlyOneSignalHandler)

	stop := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		// wait for interrupt
		signal := <-sigint
		logger.Infof("received %s", signal)

		// close channel to indicate to caller interrupt has been received
		close(stop)

		// second signal. Exit directly.
		signal = <-sigint
		logger.Infof("received %s. Exiting with 1", signal)
		os.Exit(1)
	}()

	return stop
}
