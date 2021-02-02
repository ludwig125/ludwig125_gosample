package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/net/context"
)

func main() {
	// いろいろなio.reader

	var r io.Reader
	var err error

	// ファイルから読み込み
	r, err = os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	print(r)

	// 文字列から読み込み
	r = strings.NewReader("Read will return these bytes")
	print(r)

	// bytes.Bufferもreader
	var buf *bytes.Buffer
	b := bytes.NewBufferString("This is string")
	buf = b
	r = buf
	print(r)

	// 文字列から読み込み
	r = strings.NewReader("Read will return these bytes")
	// d, err := gzipData(ioReaderToBytes(r))
	// if err != nil {
	// 	log.Fatalf("failed to gzip: %v", err)
	// }

	if err := makeGzip(buf, ioReaderToBytes(r)); err != nil {
		log.Fatalf("failed to gzip: %v", err)
	}
	d := buf.Bytes()
	d2, err := gunzipData(d)
	if err != nil {
		log.Fatalf("failed to gunzip: %v", err)
	}
	r2 := bytes.NewReader(d)
	print(r2)
	fmt.Println()
	fmt.Println(string(d2))

	// // 外部コマンドから読み込み
	// command := "for i in `seq 1 5`; do echo $i; sleep 1;done"
	// // command := "hoge"
	// r, errCh := commandExec(command)
	// if r != nil {
	// 	print(r)
	// }
	// fmt.Println(<-errCh) // command := "hoge" の場合はここでエラーを出力する
}

func print(r io.Reader) {
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("\nn", n)
}

func commandExec(command string) (io.ReadCloser, chan error) {
	cmdErrCh := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())

	// cmd := exec.Command("bash", "-c", command)
	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		defer cancel()
		defer close(cmdErrCh)

		cmdErrCh <- fmt.Errorf("failed to exec Command StdoutPipe: %v", err)
		return nil, cmdErrCh
	}

	cmd.Start()

	go func() {
		defer cancel()
		defer close(cmdErrCh)

		if err := cmd.Wait(); err != nil {
			cmdErrCh <- fmt.Errorf("failed to exec Command Wait: %w", err)
		}
	}()

	return stdout, cmdErrCh
}

func ioReaderToBytes(r io.Reader) []byte {
	buf := new(bytes.Buffer)
	io.Copy(buf, r)
	return buf.Bytes()
}

func gzipData(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	if err := gz.Flush(); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil

}

func gunzipData(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}

	var resB bytes.Buffer
	if _, err := resB.ReadFrom(r); err != nil {
		return nil, err
	}

	return resB.Bytes(), nil
}

func makeGzip(w io.Writer, content []byte) error {
	gz := gzip.NewWriter(w)
	defer gz.Close()
	if _, err := gz.Write(content); err != nil {
		return err
	}
	return nil
}
