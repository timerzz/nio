package nio

import "io"

// nReader
type nReader struct {
	r io.Reader
	f NCallBack
}

func NReader(r io.Reader, f NCallBack) *nReader {
	return &nReader{
		r: r,
		f: f,
	}
}

func (r *nReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n > 0 && r.f != nil {
		r.f(n)
	}
	return
}

type nReadCloser struct {
	rc io.ReadCloser
	f  NCallBack
}

func NReadCloser(r io.ReadCloser, f NCallBack) *nReadCloser {
	return &nReadCloser{
		rc: r,
		f:  f,
	}
}

func (r *nReadCloser) Read(p []byte) (n int, err error) {
	n, err = r.rc.Read(p)
	if n > 0 && r.f != nil {
		r.f(n)
	}
	return
}

func (r *nReadCloser) Close() error {
	return r.rc.Close()
}
