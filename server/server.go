package server

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/mailgun/vulcand/Godeps/_workspace/src/github.com/mailgun/metrics"
	"github.com/mailgun/vulcand/backend"
)

type Server interface {
	UpsertHost(host *backend.Host) error
	DeleteHost(hostname string) error
	UpdateHostKeyPair(hostname string, keyPair *backend.KeyPair) error

	AddHostListener(host *backend.Host, l *backend.Listener) error
	DeleteHostListener(host *backend.Host, listenerId string) error

	UpsertLocation(host *backend.Host, loc *backend.Location) error
	DeleteLocation(host *backend.Host, locationId string) error

	UpdateLocationUpstream(host *backend.Host, loc *backend.Location) error
	UpdateLocationPath(host *backend.Host, loc *backend.Location, path string) error
	UpdateLocationOptions(host *backend.Host, loc *backend.Location) error

	UpsertLocationMiddleware(host *backend.Host, loc *backend.Location, mi *backend.MiddlewareInstance) error
	DeleteLocationMiddleware(host *backend.Host, loc *backend.Location, mType, mId string) error

	UpsertEndpoint(upstream *backend.Upstream, e *backend.Endpoint, affectedLocations []*backend.Location) error
	DeleteEndpoint(upstream *backend.Upstream, endpointId string, affectedLocations []*backend.Location) error

	GetStats(hostname, locationId string, e *backend.Endpoint) *backend.EndpointStats

	TakeFiles([]*FileDescriptor) error
	GetFiles() ([]*FileDescriptor, error)

	Start() error
	Stop(wait bool)
}

type Options struct {
	MetricsClient   metrics.Client
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	MaxHeaderBytes  int
	DefaultListener *backend.Listener
	Files           []*FileDescriptor
}

type NewServerFn func(id int) (Server, error)

type FileDescriptor struct {
	Address backend.Address
	File    *os.File
}

func (fd *FileDescriptor) ToListener() (net.Listener, error) {
	listener, err := net.FileListener(fd.File)
	if err != nil {
		return nil, err
	}
	fd.File.Close()
	return listener, nil
}

func (fd *FileDescriptor) String() string {
	return fmt.Sprintf("FileDescriptor(%s, %d)", fd.Address, fd.File.Fd())
}
