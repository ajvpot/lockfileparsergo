package reader

import (
	"io"
)

// NamedReadCloser represents a named ReadCloser interface
type NamedReadCloser interface {
	io.ReadCloser
	Name() string
}

// NewNamedReadCloser creates a NamedReadCloser for loading external streams into the V8 context.
func NewNamedReadCloser(name string, rdr io.Reader) NamedReadCloser {
	rc, ok := rdr.(io.ReadCloser)
	if !ok {
		rc = io.NopCloser(rdr)
	}
	return &namedReadCloser{
		name: name,
		cr:   rc,
	}
}

type namedReadCloser struct {
	name string
	cr   io.ReadCloser
}

func (n *namedReadCloser) Close() error {
	return n.cr.Close()
}
func (n *namedReadCloser) Read(p []byte) (int, error) {
	return n.cr.Read(p)
}
func (n *namedReadCloser) Name() string {
	return n.name
}
