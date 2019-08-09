package cmd

import (
	"bytes"
	"fmt"
	"github.com/devkanro/eol/support"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode"
)

var rootCmd = &cobra.Command{
	Use:     "eol",
	Short:   "eol is using for changing text's the end of line.",
	Long:    `A line separators changer for text, it supports crlf(\r\n), lf(\n) and cr(\r).`,
	Version: "v1.0",
}

var output string

var AllEncoding = []support.Encoding{support.UTF8, support.UTF16LE, support.UTF16BE, support.GBK, support.GB18030, support.Big5, support.ShiftJIS, support.EUCKR}

var AllEol = []*support.Eol{support.LF, support.CRLF, support.CR}

var EncodingMap = map[string][]support.Encoding{
	"all":     AllEncoding,
	"utf8":    {support.UTF8},
	"utf16":   {support.UTF16LE, support.UTF16BE},
	"utf16l":  {support.UTF16LE},
	"utf16b":  {support.UTF16BE},
	"unicode": {support.UTF8, support.UTF16LE, support.UTF16BE},
	"gbk":     {support.GBK},
	"gb18030": {support.GB18030},
	"big5":    {support.Big5},
	"shiftj":  {support.ShiftJIS},
	"euckr":   {support.EUCKR},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleArgs(args []string, eol []rune, encoding string) error {
	for _, path := range args {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			err := handleFolder(path, eol, encoding)
			if err != nil {
				fmt.Printf("Skip convert folder '%s' due to error: %s\n", path, err)
				continue
			}
		} else {
			inputFilePath := filepath.Join(path)
			err := handleFile(inputFilePath, eol, encoding)
			if err != nil {
				fmt.Printf("Skip convert file '%s' due to error: %s\n", inputFilePath, err)
				continue
			}
		}
	}

	return nil
}

func handleFolder(path string, eol []rune, encoding string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, fileInfo := range files {
		inputFilePath := filepath.Join(path, fileInfo.Name())
		if fileInfo.IsDir() {
			err := handleFolder(inputFilePath, eol, encoding)
			if err != nil {
				fmt.Printf("Skip convert folder '%s' due to error: %s\n", inputFilePath, err)
				continue
			}
		} else {
			err := handleFile(inputFilePath, eol, encoding)
			if err != nil {
				fmt.Printf("Skip convert file '%s' due to error: %s\n", inputFilePath, err)
				continue
			}
		}
	}

	return nil
}

func handleFile(path string, eol []rune, encoding string) error {
	encodings, exist := EncodingMap[encoding]
	if !exist {
		encodings = AllEncoding
	}

	for _, encoding := range encodings {
		ok, err := handleFileWithEncoding(path, eol, encoding)
		if err != nil {
			return err
		}

		if ok {
			return nil
		}
	}

	fmt.Printf("Skip convert file '%s' due to binary char\n", path)

	return nil
}

func handleFileWithEncoding(path string, eol []rune, encoding support.Encoding) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = file.Close()
	}()

	reader := encoding.Reader(file)
	output := new(bytes.Buffer)
	writer := encoding.Writer(output)

	var prev rune = 0
	for char, _, err := reader.ReadRune(); ; char, _, err = reader.ReadRune() {
		if err != nil {
			if err == io.EOF {
				if prev == '\r' {
					_, err = writer.WriteRunes(eol)
					if err != nil {
						return false, err
					}
				}
				break
			}
			return false, err
		}

		switch char {
		case '\r':
			if prev == '\r' {
				_, err = writer.WriteRunes(eol)
				if err != nil {
					return false, err
				}
			}
		case '\n':
			if prev == '\r' {
				_, err = writer.WriteRunes(eol)
				if err != nil {
					return false, err
				}
			} else {
				_, err = writer.WriteRunes(eol)
				if err != nil {
					return false, err
				}
			}
		default:
			if (unicode.IsControl(char) && char != '\t') || char == unicode.ReplacementChar{
				return false, nil
			}

			if prev == '\r' {
				_, err = writer.WriteRunes(eol)
				if err != nil {
					return false, err
				}
			}
			_, err = writer.WriteRune(char)
			if err != nil {
				return false, err
			}
		}
		prev = char
	}

	file.Close()
	file, err = os.Create(path)
	if err != nil {
		return false, err
	}

	err = writer.Flush()
	if err != nil {
		return false, err
	}

	io.Copy(file, output)
	return true, nil
}