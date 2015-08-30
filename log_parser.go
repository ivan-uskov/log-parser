package main

import (
	"fmt"
	"flag"
	"time"
	"github.com/ivan-uskov/log-parser/textio"
	"github.com/ivan-uskov/log-parser/logparser"
	"github.com/ivan-uskov/log-parser/logexport"
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

func parseData(ss *textio.StringStream, msgCollection *logexport.MessageCollection) {
	var logString string
	for ss.ReadString(&logString) {
		ci, err := logparser.ParseClientInfo(logString)
		if err != nil {
			fmt.Println(err)
			continue
		}
		msgCollection.AddClientInfo(ci)
	}

	if !ss.IsSuccess() {
		panic("Read error")
	}
}

func sendMessage(msg *logexport.MessageCollection){
	err := msg.SendAll()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	parameters := parseFlags()
	for {
		fmt.Println("Start sending")
		ss := initStream(*parameters.ProcessFile)
		defer ss.Close()
		msgCollection := logexport.NewMessageCollection(*parameters.Project, *parameters.Host)

		parseData(&ss, &msgCollection)
		sendMessage(&msgCollection)
		fmt.Println("Sended -> go to sleep")
		time.Sleep(5000 * time.Millisecond)
	}
}
