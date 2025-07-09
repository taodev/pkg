package util

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockBuffer struct {
	buf   []byte
	stepN int
	err   error
}

func (b *mockBuffer) Write(p []byte) (n int, err error) {
	if b.stepN > 0 {
		wn := b.stepN
		if len(p) < wn {
			wn = len(p)
		}
		p = p[:wn]
	}

	b.buf = append(b.buf, p...)
	return len(p), b.err
}

func TestWriteFull(t *testing.T) {
	buf := &mockBuffer{
		stepN: 2,
	}
	data := []byte("hello world")
	n, err := WriteFull(buf, data)
	assert.NoError(t, err)
	assert.Equal(t, len(data), n)
	assert.Equal(t, data, buf.buf)

	buf = &mockBuffer{
		stepN: 2,
		err:   io.EOF,
	}
	n, err = WriteFull(buf, data)
	assert.Error(t, err)
	assert.Equal(t, buf.stepN, n)
}
