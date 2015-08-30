package logexport

import (
	"github.com/ivan-uskov/go-course/log_parser/datastruct"
	"encoding/json"
	"net"
)

type Message struct {
	host string
	data datastruct.ExportedData
}

func NewMessage(project string, host string) Message {
	return Message {
		host: host,
		data: datastruct.ExportedData {
			Project: project,
		},
	}
}

func (this *Message) AddClientInfo(ci datastruct.ClientInfo) {
	this.data.Data = append(this.data.Data, ci)
}

func (this *Message) Send() error {
	data, err := json.Marshal(this.data)
	if err != nil {
		return err
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", this.host)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return err
}

