package link

import (
	"log"
	"sync"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/hiroara/dirlink/config"
)

type Link struct {
	Src     string
	Dest    string
	Verbose bool
	server  *fuse.Server
	mutex   sync.Mutex
}

func New(src, dest string, verbose bool) (*Link, error) {
	s, err := homedir.Expand(src)
	if err != nil {
		return nil, err
	}
	d, err := homedir.Expand(dest)
	if err != nil {
		return nil, err
	}
	return &Link{Src: s, Dest: d, Verbose: verbose}, nil
}

func FromEntry(entry *config.BindEntry, verbose bool) ([]*Link, error) {
	links := make([]*Link, 0, len(entry.Links))
	for _, link := range entry.Links {
		l, err := New(entry.Src, link, verbose)
		if err != nil {
			return nil, err
		}
		links = append(links, l)
	}
	return links, nil
}

func (l *Link) Mount() error {
	l.lock()
	defer l.unlock()
	root, err := fs.NewLoopbackRoot(l.Src)
	if err != nil {
		return &linkFailure{link: l, cause: err}
	}
	srv, err := fs.Mount(l.Dest, root, nil)
	if err != nil {
		return &linkFailure{link: l, cause: err}
	}
	l.server = srv
	if l.Verbose {
		log.Printf("Mount: %s -> %s", l.Src, l.Dest)
	}
	return nil
}

func (l *Link) Wait() {
	l.server.Wait()
}

func (l *Link) Unmount() error {
	l.lock()
	defer l.unlock()
	if l.server == nil {
		return nil
	}
	err := l.server.Unmount()
	if err != nil {
		return err
	}
	l.server = nil
	if l.Verbose {
		log.Printf("Unmount: %s -> %s", l.Src, l.Dest)
	}
	return nil
}

func (l *Link) lock() {
	l.mutex.Lock()
}

func (l *Link) unlock() {
	l.mutex.Unlock()
}
