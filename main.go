package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type arg struct {
	fhost string
	thost string
}

func ParseConfLine(data []byte) arg {
	a := new(arg)
	tmp := make([]byte, 0)
	for _, c := range data {
		if c == 0x20 { //换到下一个内容
			a.fhost = string(tmp)
			tmp = make([]byte, 0)
		} else if c == '\r' {
			break
		} else {
			tmp = append(tmp, c)
		}
	}
	return *a
}

func ParseConfigFile(conf string) ([]arg, error) {
	args := make([]arg, 0)
	fp, err := os.Open(conf)
	if err != nil {
		return args, err
	}
	bfp := bufio.NewReader(fp)
	for true {
		ldata, _, err := bfp.ReadLine()
		if err != nil {
			return args, err
		}
		args = append(args, ParseConfLine(ldata))
	}
	return args, nil
}

func main() {
	args := flag.String("c", "", "configure file")
	flag.Parse()
	if *args == "" {
		fmt.Println("configure not set,please use -c")
		return
	}
	fmt.Println("use configure file ", *args)
	conf, err := ParseConfigFile(*args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)
}
