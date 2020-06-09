package gor

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type fileInputReader struct {
	reader    *bufio.Reader
	data      [][]byte
	file      *os.File
	timestamp int64
}

var (
	ByteHttp          = []byte("HTTP/1.")
	ByteHost          = []byte("Host")
	ByteAgent         = []byte("User-Agent")
	ByteContentType   = []byte("Content-Type")
	ByteRemoteIp      = []byte("RemoteIp")
	ByteXForwardedFor = []byte("X-Forwarded-For")
	ByteAuthorization = []byte("Authorization")
	ByteDid           = []byte("did")
	ByteConnection    = []byte("Connection:")
)

func (f *fileInputReader) parseNext() error {
	payloadSeparatorAsBytes := []byte(payloadSeparator)
	var buffer bytes.Buffer
	var body bytes.Buffer
	var header bytes.Buffer
	list := [7][]byte{}
	isFirst := true
	isBody := false
	for {
		line, err := f.reader.ReadBytes('\n')
		//strings.Join(list,"||")
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return err
			}

			if err == io.EOF {
				f.file.Close()
				f.file = nil
				return err
			}
		}

		if isFirst {
			meta := bytes.Split(line, []byte{' '})
			// time
			//list[0] =  strings.TrimSpace(string(meta[2]))
			list[0] = bytes.TrimSpace(meta[2])
			isFirst = false
			continue
		}
		contains := bytes.Contains(line, ByteHttp)
		if contains {
			meta := bytes.Split(line, []byte{' '})
			// method
			list[1] = bytes.TrimSpace(meta[0])
			// url
			list[2] = bytes.TrimSpace(meta[1])
			continue
		}
		contains = bytes.Contains(line, ByteHost)
		if contains {
			meta := bytes.Split(line, []byte{':'})
			// host
			list[3] = bytes.TrimSpace(meta[1])
			// port
			if len(meta) == 3 {
				list[4] = bytes.TrimSpace(meta[2])
			} else {
				list[4] = bytes.TrimSpace([]byte("80"))
			}
			continue
		}
		//contains = bytes.Contains(line, ByteAgent)
		//if contains {
		//	//meta := bytes.Split(line, []byte{' '})
		//	// User-Agent
		//	list[5] = bytes.TrimSpace(bytes.TrimLeft(line, "User-Agent:"))
		//	continue
		//}
		//
		//contains = bytes.Contains(line, ByteContentType)
		//if contains {
		//
		//	//meta := bytes.Split(line, []byte{':'})
		//	// User-Agent
		//	list[6] = bytes.TrimSpace(bytes.TrimLeft(line, "Content-Type:"))
		//	continue
		//}
		//
		//contains = bytes.Contains(line, ByteRemoteIp)
		//if contains {
		//	//RemoteIp
		//	list[7] = bytes.TrimSpace(bytes.TrimLeft(line, "RemoteIp:"))
		//	continue
		//}
		//contains = bytes.Contains(line, ByteXForwardedFor)
		//if contains {
		//	//X-Forwarded-For
		//	list[8] = bytes.TrimSpace(bytes.TrimLeft(line, "X-Forwarded-For:"))
		//	continue
		//}
		//contains = bytes.Contains(line, ByteAuthorization)
		//if contains {
		//	//Authorization
		//	list[9] = bytes.TrimSpace(bytes.TrimLeft(line, "Authorization:"))
		//	continue
		//}
		//contains = bytes.Contains(line, ByteDid)
		//if contains && !isBody {
		//	//did
		//	list[10] = bytes.TrimSpace(bytes.TrimLeft(line, "did:"))
		//	continue
		//}
		// 全是header
		if !isBody && !bytes.Equal([]byte("\r\n"), line) {
			contains = bytes.Contains(line, ByteConnection)
			if contains {
				continue
			}
			header.Write(bytes.TrimSpace(line))
			header.Write([]byte("@@"))
		}
		if bytes.Equal([]byte("\r\n"), line) {
			//fmt.Println(payloadSeparatorAsBytes[1:],line)
			isBody = true
			continue
		}

		if bytes.Equal(payloadSeparatorAsBytes[1:], line) {
			//all := bytes.ReplaceAll(bytes.TrimSpace(body.Bytes()), []byte("\r\n"), []byte(""))
			//all = bytes.ReplaceAll(all,[]byte("\r"), []byte(""))
			all := bytes.ReplaceAll(bytes.TrimSpace(body.Bytes()), []byte("\n"), []byte(""))
			//headAll := bytes.ReplaceAll(bytes.TrimSpace(header.Bytes()), []byte("\n"), []byte(""))
			list[5] = bytes.TrimRight(header.Bytes(), "@@")
			list[6] = all
			slice := make([][]byte, len(list))
			copy(slice, list[:])
			f.data = slice
			return nil
		}
		if isBody {
			body.Write(line)
		}
		buffer.Write(line)
	}

}

func (f *fileInputReader) ReadPayload() [][]byte {
	defer func() {
		if f.data == nil {
			return
		}
		if err := f.parseNext(); err == io.EOF {
			f.data = nil
		}
	}()

	return f.data
}
func (f *fileInputReader) Close() error {
	if f.file != nil {
		f.file.Close()
	}

	return nil
}

func NewFileInputReader(path string) (*fileInputReader, error) {
	matches, err := filepath.Glob(path)
	if err != nil {
		log.Println("Wrong file pattern", path, err)
		return nil, err
	}

	if len(matches) == 0 {
		log.Println("No files match pattern: ", path)
		return nil, errors.New("no matching files")
	}

	file, err := os.Open(matches[0])

	if err != nil {
		log.Println(err)
		return nil, err
	}

	r := &fileInputReader{file: file}
	if strings.HasSuffix(path, ".gz") {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		r.reader = bufio.NewReader(gzReader)
	} else {
		r.reader = bufio.NewReader(file)
	}
	r.parseNext()
	return r, nil

}

type FileOutPutWrite struct {
	write *bufio.Writer
	file  *os.File
}

func NewFileOutPutWrite(fileName string) (*FileOutPutWrite, error) {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return nil, err
	}
	write := new(FileOutPutWrite)
	write.file = file
	write.write = bufio.NewWriter(file)
	return write, nil
}

func (c *FileOutPutWrite) Write(data []byte) (err error) {
	_, err = c.write.Write(data)
	return
}

func (c *FileOutPutWrite) Close() error {
	_ = c.write.Flush()
	if c.file != nil {
		return c.file.Close()
	}
	return nil
}
