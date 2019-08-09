package support

import (
	"bufio"
	"golang.org/x/text/encoding"
	"io"
)

type EncodingReader interface {
	ReadRune() (r rune, size int, err error)
}

type Encoding interface {
	Reader(reader io.Reader) EncodingReader
	Writer(writer io.Writer) EncodingWriter
}

type internalEncoding struct {
	encoding encoding.Encoding
}

func (encoding *internalEncoding) Reader(raw io.Reader) EncodingReader {
	return &internalEncodingReader{
		reader:   bufio.NewReader(encoding.encoding.NewDecoder().Reader(raw)),
		encoding: encoding.encoding,
	}
}

func (encoding *internalEncoding) Writer(raw io.Writer) EncodingWriter {
	return &internalEncodingWriter{
		writer:   bufio.NewWriter(encoding.encoding.NewEncoder().Writer(raw)),
		encoding: encoding.encoding,
	}
}

type internalEncodingReader struct {
	reader   *bufio.Reader
	encoding encoding.Encoding
}

func (encoding *internalEncodingReader) ReadRune() (r rune, size int, err error) {
	return encoding.reader.ReadRune()
}

type EncodingWriter interface {
	Flush() error
	WriteRune(r rune) (size int, err error)
	WriteRunes(r []rune) (size int, err error)
}

type internalEncodingWriter struct {
	writer   *bufio.Writer
	encoding encoding.Encoding
}

func (encoding *internalEncodingWriter) Flush() error {
	return encoding.writer.Flush()
}

func (encoding *internalEncodingWriter) WriteRune(r rune) (size int, err error) {
	return encoding.writer.WriteRune(r)
}

func (encoding *internalEncodingWriter) WriteRunes(r []rune) (size int, err error) {
	size = 0
	for _, value := range r {
		s, err := encoding.WriteRune(value)
		size += s
		if err != nil {
			return size, err
		}
	}
	return size, nil
}
