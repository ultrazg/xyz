package utils

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	port int
	doc  bool
)

func InitFlag() (int, bool) {
	flag.IntVar(&port, "p", 23020, "指定服务监听的端口")
	flag.BoolVar(&doc, "d", false, "打开 Api 文档")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]\n", "xyz")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	return port, doc
}

func CheckPort(port int) error {
	address := ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", address)
	if err == nil {
		listener.Close()
		return nil
	}

	return fmt.Errorf("端口 %d 不可用", port)
}
