package main

import "os"
import "bufio"
import "log"
import "strings"
import "strconv"
import "gomiscutils"
import "fmt"
import "compress/gzip"

type NullValue struct{}

func main(){
	var pathOfFile1 string = os.Args[1]
	var pathOfFile2 string = os.Args[2]
	var in string
	func() {
		defer func() {
			if e := recover();e!=nil { in = "in" }
		}()
		in = os.Args[3]
	}()
	var fullrslist map[uint32]NullValue = make(map[uint32]NullValue,100000000)
	var checkrslist map[uint32]NullValue = make(map[uint32]NullValue,500000)
	readrslist(pathOfFile1, fullrslist, true)
	readrslist(pathOfFile2, checkrslist, false)
	for k, _ := range checkrslist {
		if _, ok := fullrslist[k]; inorout(in,ok) {
			fmt.Println(k)
		}
	}
}

func inorout(b string, ok bool) bool {
	switch b {
	case "in":
		return ok
	case "out":
		return !ok
	default:
		return ok
	}
	panic("nil")
}


func readrslist(pathOfFile string, rslist map[uint32]NullValue, gziped bool) {
	var r *gzip.Reader
	var f *os.File
	var fReader *bufio.Reader
	var e error

	f,e = os.Open(pathOfFile)
	if e!=nil { log.Fatal(e) }

	if gziped {
		r,e = gzip.NewReader(f)
		if e!=nil { log.Fatal(e) }
		fReader = bufio.NewReader(r)
	} else {
		fReader = bufio.NewReader(f)
	}
	var lines chan string = gomiscutils.Readline(fReader)
	var s string
	var i uint64
	for l := range lines {
		s = gomiscutils.TrimNewLine(l)
		if strings.HasPrefix(s,"rs") {
			s = s[2:]
		}
		i, _ = strconv.ParseUint(s, 10, 64)
		rslist[uint32(i)]= NullValue{}
	}
}
