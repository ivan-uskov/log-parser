package logexport

import (
	"github.com/ivan-uskov/log-parser/logparser"
)


type MessageCollection struct {
	messages []Message
	host string
	project string
}

func NewMessageCollection(project string, host string) MessageCollection {
	return MessageCollection{
		messages: []Message{},
		host: host,
		project: project,
	}
}

func (this *MessageCollection) AddClientInfo(ci map[string]string) {
	date := ci[logparser.DATE]
	useragent := ci[logparser.USERAGENT]
	browser, version := logparser.ParseUserAgent(useragent)
	isAdded := false

	for key, val := range this.messages {
		if val.GetDate() == date {
			this.messages[key].AddBrowserInfo(browser, version)
			isAdded = true
		}
	}

	if !isAdded {
		msg := NewMessage(this.project, this.host, date)
		msg.AddBrowserInfo(browser, version)
		this.messages = append(this.messages, msg)
	}
}

func (this *MessageCollection) SendAll() error {
	for _, val := range this.messages {
		err := val.Send()
		if err != nil {
			return err
		}
	}

	return  nil
}
