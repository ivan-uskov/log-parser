package main

import (
	"fmt"
	"flag"
	"github.com/ivan-uskov/go-course/log_parser/textio"
	"github.com/ivan-uskov/go-course/log_parser/logparser"
	"github.com/ivan-uskov/go-course/log_parser/logexport"
)

type Parameters struct  {
	Project *string
	ProcessFile *string
	Host *string
}

func parseFlags() (Parameters) {
	parameters := Parameters{
		Project: flag.String("p", "isonline|iscloud", "project name"),
		ProcessFile: flag.String("f", "access.log", "nginx success log file path"),
		Host: flag.String("h", "127.0.0.1", "consumer host"),
	}
	flag.Parse()

	return parameters
}

func initStream(fileName string) textio.StringStream {
	ss := textio.NewStringStream(fileName)
	if !ss.IsSuccess() {
		panic(ss.GetError())
	}

	return ss
}

func parseData(ss *textio.StringStream, msg *logexport.Message) {
	var logString string
	for ss.ReadString(&logString) {
		ci, err := logparser.ParseClientInfo(logString)
		if err != nil {
			fmt.Println(err)
			continue
		}
		msg.AddClientInfo(ci)
	}

	if !ss.IsSuccess() {
		panic("Read error")
	}
}

func sendMessage(msg *logexport.Message){
	err := msg.Send()
	if err != nil {
		panic(err)
	}
}

func main() {
	parameters := parseFlags()
	ss := initStream(*parameters.ProcessFile)
	defer ss.Close()
	msg := logexport.NewMessage(*parameters.Project, *parameters.Host)

	parseData(&ss, &msg)
	sendMessage(&msg)
}
