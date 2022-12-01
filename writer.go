package nio

import "io"

// nWriter
type nWriter struct {
	w io.Writer
	f NCallBack
}

func NWriter(w io.Writer, f NCallBack) *nWriter {
	return &nWriter{
		w: w,
		f: f,
	}
}

func (w *nWriter) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	if n > 0 && w.f != nil {
		w.f(n)
	}
	return
}

type nWriteCloser struct {
	wc io.WriteCloser
	f  NCallBack
}

func NWriteCloser(w io.WriteCloser, f NCallBack) *nWriteCloser {
	return &nWriteCloser{
		wc: w,
		f:  f,
	}
}

func (w *nWriteCloser) Write(p []byte) (n int, err error) {
	n, err = w.wc.Write(p)
	if n > 0 && w.f != nil {
		w.f(n)
	}
	return
}

func (w *nWriteCloser) Close() error {
	return w.wc.Close()
}
