// Generated by `wit-bindgen-wrpc-go` 0.1.1. DO NOT EDIT!
package atomics

import (
	bytes "bytes"
	context "context"
	binary "encoding/binary"
	errors "errors"
	fmt "fmt"
	wrpc__keyvalue__store "github.com/wrpc/keyvalue/imports/wrpc/keyvalue/store"
	wrpc "github.com/wrpc/wrpc/go"
	io "io"
	slog "log/slog"
	math "math"
	utf8 "unicode/utf8"
)

type Error = wrpc__keyvalue__store.Error

// Atomically increment the value associated with the key in the store by the given delta. It
// returns the new value.
//
// If the key does not exist in the store, it creates a new key-value pair with the value set
// to the given delta.
//
// If any other error occurs, it returns an `Err(error)`.
func Increment(ctx__ context.Context, wrpc__ wrpc.Invoker, bucket string, key string, delta uint64) (r0__ *wrpc.Result[uint64, Error], close__ func() error, err__ error) {
	if err__ = wrpc__.Invoke(ctx__, "wrpc:keyvalue/atomics@0.2.0-draft", "increment", func(w__ wrpc.IndexWriter, r__ wrpc.IndexReadCloser) error {
		close__ = r__.Close
		var buf__ bytes.Buffer
		writes__ := make(map[uint32]func(wrpc.IndexWriter) error, 3)
		write0__, err__ := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
			n := len(v)
			if n > math.MaxUint32 {
				return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
			}
			if err = func(v int, w io.Writer) error {
				b := make([]byte, binary.MaxVarintLen32)
				i := binary.PutUvarint(b, uint64(v))
				slog.Debug("writing string byte length", "len", n)
				_, err = w.Write(b[:i])
				return err
			}(n, w); err != nil {
				return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
			}
			slog.Debug("writing string bytes")
			_, err = w.Write([]byte(v))
			if err != nil {
				return fmt.Errorf("failed to write string bytes: %w", err)
			}
			return nil
		}(bucket, &buf__)
		if err__ != nil {
			return fmt.Errorf("failed to write `bucket` parameter: %w", err__)
		}
		if write0__ != nil {
			writes__[0] = write0__
		}
		write1__, err__ := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
			n := len(v)
			if n > math.MaxUint32 {
				return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
			}
			if err = func(v int, w io.Writer) error {
				b := make([]byte, binary.MaxVarintLen32)
				i := binary.PutUvarint(b, uint64(v))
				slog.Debug("writing string byte length", "len", n)
				_, err = w.Write(b[:i])
				return err
			}(n, w); err != nil {
				return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
			}
			slog.Debug("writing string bytes")
			_, err = w.Write([]byte(v))
			if err != nil {
				return fmt.Errorf("failed to write string bytes: %w", err)
			}
			return nil
		}(key, &buf__)
		if err__ != nil {
			return fmt.Errorf("failed to write `key` parameter: %w", err__)
		}
		if write1__ != nil {
			writes__[1] = write1__
		}
		write2__, err__ := (func(wrpc.IndexWriter) error)(nil), func(v uint64, w io.Writer) (err error) {
			b := make([]byte, binary.MaxVarintLen64)
			i := binary.PutUvarint(b, uint64(v))
			slog.Debug("writing u64")
			_, err = w.Write(b[:i])
			return err
		}(delta, &buf__)
		if err__ != nil {
			return fmt.Errorf("failed to write `delta` parameter: %w", err__)
		}
		if write2__ != nil {
			writes__[2] = write2__
		}
		_, err__ = w__.Write(buf__.Bytes())
		if err__ != nil {
			return fmt.Errorf("failed to write parameters: %w", err__)
		}
		r0__, err__ = func(r wrpc.IndexReader, path ...uint32) (*wrpc.Result[uint64, Error], error) {
			slog.Debug("reading result status byte")
			status, err := r.ReadByte()
			if err != nil {
				return nil, fmt.Errorf("failed to read result status byte: %w", err)
			}
			switch status {
			case 0:
				slog.Debug("reading `result::ok` payload")
				v, err := func(r io.ByteReader) (uint64, error) {
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
					return nil, fmt.Errorf("failed to read `result::ok` value: %w", err)
				}
				return &wrpc.Result[uint64, Error]{Ok: &v}, nil
			case 1:
				slog.Debug("reading `result::err` payload")
				v, err := func() (*Error, error) {
					v, err := func(r wrpc.IndexReader, path ...uint32) (*wrpc__keyvalue__store.Error, error) {
						v := &wrpc__keyvalue__store.Error{}
						n, err := func(r io.ByteReader) (uint8, error) {
							var x uint8
							var s uint
							for i := 0; i < 2; i++ {
								slog.Debug("reading u8 discriminant byte", "i", i)
								b, err := r.ReadByte()
								if err != nil {
									if i > 0 && err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									return x, fmt.Errorf("failed to read u8 discriminant byte: %w", err)
								}
								if s == 7 && b > 0x01 {
									return x, errors.New("discriminant overflows an 8-bit integer")
								}
								if b < 0x80 {
									return x | uint8(b)<<s, nil
								}
								x |= uint8(b&0x7f) << s
								s += 7
							}
							return x, errors.New("discriminant overflows an 8-bit integer")
						}(r)
						if err != nil {
							return nil, fmt.Errorf("failed to read discriminant: %w", err)
						}
						switch wrpc__keyvalue__store.ErrorDiscriminant(n) {
						case wrpc__keyvalue__store.ErrorNoSuchStore:
							return v.SetNoSuchStore(), nil
						case wrpc__keyvalue__store.ErrorAccessDenied:
							return v.SetAccessDenied(), nil
						case wrpc__keyvalue__store.ErrorOther:
							payload, err := func(r interface {
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
								return nil, fmt.Errorf("failed to read `other` payload: %w", err)
							}
							return v.SetOther(payload), nil
						default:
							return nil, fmt.Errorf("unknown discriminant value %d", n)
						}
					}(r, path...)
					return (*Error)(v), err
				}()

				if err != nil {
					return nil, fmt.Errorf("failed to read `result::err` value: %w", err)
				}
				return &wrpc.Result[uint64, Error]{Err: v}, nil
			default:
				return nil, fmt.Errorf("invalid result status byte %d", status)
			}
		}(r__, []uint32{0}...)
		if err__ != nil {
			return fmt.Errorf("failed to read result 0: %w", err__)
		}
		return nil
	}); err__ != nil {
		err__ = fmt.Errorf("failed to invoke `increment`: %w", err__)
		return
	}
	return
}
