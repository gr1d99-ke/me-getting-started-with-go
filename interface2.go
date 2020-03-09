package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type MyWriterCloser interface {
	Writer
	Closer
}

type MyWriter struct {}

func (mw MyWriter) Write(data []byte) (int, error)  {
	return 0, nil
}

func (mw *MyWriter) Close() error {
	fmt.Println("aaaaa")
	return nil
}

func main()  {
	var c MyWriterCloser = MyWriter{}
	fmt.Println(c.Close())
}