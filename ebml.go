package ebml

import (
	"fmt"
	"io"
)

// NewReader creates ...
func NewReader(rs io.ReaderAt) (*Reader, error) {
	r := &Reader{
		rs: rs,
	}
	err := r.init()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Reader describes ...
type Reader struct {
	rs io.ReaderAt
}

func (r *Reader) init() error {
	return fmt.Errorf("not implemented")
}

// ReadAt reads len(buf) bytes into buf starting at offset off in the source.
func (r *Reader) ReadAt(buf []byte, offset int64) (int, error) {
	return r.rs.ReadAt(buf, offset)
}

// Version returns ...
func (r *Reader) Version() uint32 {
	return 0
}

// Root returns ...
func (r *Reader) Root() *Element {
	return nil
}

// Element describes ...
type Element struct {
}

// Size returns ...
func (e *Element) Size() uint64 {
	return 0
}
