// Generated by `wit-bindgen-wrpc-go` 0.1.1. DO NOT EDIT!
package atomics

import (
	bytes "bytes"
	context "context"
	binary "encoding/binary"
	errors "errors"
	fmt "fmt"
	exports__wrpc__keyvalue__store "github.com/wrpc/keyvalue/server/exports/wrpc/keyvalue/store"
	wrpc "github.com/wrpc/wrpc/go"
	io "io"
	slog "log/slog"
	sync "sync"
	atomic "sync/atomic"
	utf8 "unicode/utf8"
)

type Error = exports__wrpc__keyvalue__store.Error
type Handler interface {
	// Atomically increment the value associated with the key in the store by the given delta. It
	// returns the new value.
	//
	// If the key does not exist in the store, it creates a new key-value pair with the value set
	// to the given delta.
	//
	// If any other error occurs, it returns an `Err(error)`.
	Increment(ctx__ context.Context, bucket string, key string, delta uint64) (*wrpc.Result[uint64, Error], error)
}

func ServeInterface(s wrpc.Server, h Handler) (stop func() error, err error) {
	stops := make([]func() error, 0, 1)
	stop = func() error {
		for _, stop := range stops {
			if err := stop(); err != nil {
				return err
			}
		}
		return nil
	}
	stop0, err := s.Serve("wrpc:keyvalue/atomics@0.2.0-draft", "increment", func(ctx context.Context, w wrpc.IndexWriter, r wrpc.IndexReadCloser) error {
		slog.DebugContext(ctx, "reading parameter", "i", 0)
		p0, err := func(r interface {
			io.ByteReader
			io.Reader
		}) (string, error) {
			var x uint32
			var s uint8
			for i := 0; i < 5; i++ {
				slog.Debug("reading string length byte", "i", i)
				b, err := r.ReadByte()
				if err != nil {
					if i > 0 && err == io.EOF {
						err = io.ErrUnexpectedEOF
					}
					return "", fmt.Errorf("failed to read string length byte: %w", err)
				}
				if s == 28 && b > 0x0f {
					return "", errors.New("string length overflows a 32-bit integer")
				}
				if b < 0x80 {
					x = x | uint32(b)<<s
					buf := make([]byte, x)
					slog.Debug("reading string bytes", "len", x)
					_, err = r.Read(buf)
					if err != nil {
						return "", fmt.Errorf("failed to read string bytes: %w", err)
					}
					if !utf8.Valid(buf) {
						return string(buf), errors.New("string is not valid UTF-8")
					}
					return string(buf), nil
				}
				x |= uint32(b&0x7f) << s
				s += 7
			}
			return "", errors.New("string length overflows a 32-bit integer")
		}(r)
		if err != nil {
			return fmt.Errorf("failed to read parameter 0: %w", err)
		}
		slog.DebugContext(ctx, "reading parameter", "i", 1)
		p1, err := func(r interface {
			io.ByteReader
			io.Reader
		}) (string, error) {
			var x uint32
			var s uint8
			for i := 0; i < 5; i++ {
				slog.Debug("reading string length byte", "i", i)
				b, err := r.ReadByte()
				if err != nil {
					if i > 0 && err == io.EOF {
						err = io.ErrUnexpectedEOF
					}
					return "", fmt.Errorf("failed to read string length byte: %w", err)
				}
				if s == 28 && b > 0x0f {
					return "", errors.New("string length overflows a 32-bit integer")
				}
				if b < 0x80 {
					x = x | uint32(b)<<s
					buf := make([]byte, x)
					slog.Debug("reading string bytes", "len", x)
					_, err = r.Read(buf)
					if err != nil {
						return "", fmt.Errorf("failed to read string bytes: %w", err)
					}
					if !utf8.Valid(buf) {
						return string(buf), errors.New("string is not valid UTF-8")
					}
					return string(buf), nil
				}
				x |= uint32(b&0x7f) << s
				s += 7
			}
			return "", errors.New("string length overflows a 32-bit integer")
		}(r)
		if err != nil {
			return fmt.Errorf("failed to read parameter 1: %w", err)
		}
		slog.DebugContext(ctx, "reading parameter", "i", 2)
		p2, err := func(r io.ByteReader) (uint64, error) {
			var x uint64
			var s uint8
			for i := 0; i < 10; i++ {
				slog.Debug("reading u64 byte", "i", i)
				b, err := r.ReadByte()
				if err != nil {
					if i > 0 && err == io.EOF {
						err = io.ErrUnexpectedEOF
					}
					return x, fmt.Errorf("failed to read u64 byte: %w", err)
				}
				if s == 63 && b > 0x01 {
					return x, errors.New("varint overflows a 64-bit integer")
				}
				if b < 0x80 {
					return x | uint64(b)<<s, nil
				}
				x |= uint64(b&0x7f) << s
				s += 7
			}
			return x, errors.New("varint overflows a 64-bit integer")
		}(r)
		if err != nil {
			return fmt.Errorf("failed to read parameter 2: %w", err)
		}
		slog.DebugContext(ctx, "calling `wrpc:keyvalue/atomics@0.2.0-draft.increment` handler")
		r0, err := h.Increment(ctx, p0, p1, p2)
		if err != nil {
			return fmt.Errorf("failed to handle `wrpc:keyvalue/atomics@0.2.0-draft.increment` invocation: %w", err)
		}

		var buf bytes.Buffer
		writes := make(map[uint32]func(wrpc.IndexWriter) error, 1)
		write0, err := func(v *wrpc.Result[uint64, Error], w interface {
			io.ByteWriter
			io.Writer
		}) (func(wrpc.IndexWriter) error, error) {
			switch {
			case v.Ok == nil && v.Err == nil:
				return nil, errors.New("both result variants cannot be nil")
			case v.Ok != nil && v.Err != nil:
				return nil, errors.New("exactly one result variant must non-nil")

			case v.Ok != nil:
				slog.Debug("writing `result::ok` status byte")
				if err := w.WriteByte(0); err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` status byte: %w", err)
				}
				slog.Debug("writing `result::ok` payload")
				write, err := (func(wrpc.IndexWriter) error)(nil), func(v uint64, w io.Writer) (err error) {
					b := make([]byte, binary.MaxVarintLen64)
					i := binary.PutUvarint(b, uint64(v))
					slog.Debug("writing u64")
					_, err = w.Write(b[:i])
					return err
				}(*v.Ok, w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			default:
				slog.Debug("writing `result::err` status byte")
				if err := w.WriteByte(1); err != nil {
					return nil, fmt.Errorf("failed to write `result::err` status byte: %w", err)
				}
				slog.Debug("writing `result::err` payload")
				write, err := (v.Err).WriteToIndex(w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::err` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			}
		}(r0, &buf)
		if err != nil {
			return fmt.Errorf("failed to write result value 0: %w", err)
		}
		if write0 != nil {
			writes[0] = write0
		}
		slog.DebugContext(ctx, "transmitting `wrpc:keyvalue/atomics@0.2.0-draft.increment` result")
		_, err = w.Write(buf.Bytes())
		if err != nil {
			return fmt.Errorf("failed to write result: %w", err)
		}
		if len(writes) > 0 {
			var wg sync.WaitGroup
			var wgErr atomic.Value
			for index, write := range writes {
				wg.Add(1)
				w, err := w.Index(index)
				if err != nil {
					return fmt.Errorf("failed to index writer: %w", err)
				}
				write := write
				go func() {
					defer wg.Done()
					if err := write(w); err != nil {
						wgErr.Store(err)
					}
				}()
			}
			wg.Wait()
			err := wgErr.Load()
			if err == nil {
				return nil
			}
			return err.(error)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serve `wrpc:keyvalue/atomics@0.2.0-draft.increment`: %w", err)
	}
	stops = append(stops, stop0)
	return stop, nil
}