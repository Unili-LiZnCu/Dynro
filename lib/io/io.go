package io

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFileByLine(filename string) (ret []string ,err error){
	f, err := os.Open(filename)
	if err != nil {
		return nil,err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return ret,nil
			}
			return nil,err
		}
		ret = append(ret,line)
	}
}

