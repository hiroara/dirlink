package controller

import (
	"os"
	"os/signal"

	"github.com/hiroara/dirlink/config"
	"github.com/hiroara/dirlink/link"
)

type Controller struct {
	Links []*link.Link
}

func New(links []*link.Link) *Controller {
	return &Controller{Links: links}
}

func FromEntries(entries []*config.BindEntry, verbose bool) (*Controller, error) {
	links := make([]*link.Link, 0, len(entries))
	for _, entry := range entries {
		ls, err := link.FromEntry(entry, verbose)
		if err != nil {
			return nil, err
		}
		links = append(links, ls...)
	}
	return New(links), nil
}

func (ctl *Controller) Mount() error {
	defer ctl.Unmount()
	ctl.signalFinalize(os.Interrupt)
	for _, l := range ctl.Links {
		err := l.Mount()
		if err != nil {
			return err
		}
	}
	for _, l := range ctl.Links {
		l.Wait()
	}
	return nil
}

func (ctl *Controller) Unmount() (err error) {
	errors := make([]error, 0)
	for _, l := range ctl.Links {
		if err := l.Unmount(); err != nil {
			errors = append(errors, err)
		}
	}
	return aggregateErrors(errors)
}

func (ctl *Controller) signalFinalize(sigs ...os.Signal) {
	sig := make(chan os.Signal)
	for _, s := range sigs {
		signal.Notify(sig, s)
	}
	go func() {
		<-sig
		ctl.Unmount()
	}()
}
