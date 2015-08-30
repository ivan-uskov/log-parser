package logexport

import (
	"github.com/ivan-uskov/log-parser/datastruct"
	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
)

type Message struct {
	host string
	data datastruct.ExportedData
}

func NewMessage(project string, host string, date string) Message {
	return Message {
		host: host,
		data: datastruct.ExportedData {
			Project: project,
			Date: date,
		},
	}
}

func (this *Message) AddBrowserInfo(browser string, version string) {
	isAdded := false
	for key, val := range this.data.Data {
		if val.Browser == browser && val.Version == version {
			this.data.Data[key].Count += 1
			isAdded = true
		}
	}

	if !isAdded {
		this.data.Data = append(this.data.Data, datastruct.BrowserInfo{
				Browser: browser,
				Version: version,
				Count: 1,
			})
	}
}

func (this *Message) GetDate() string {
	return this.data.Date
}

func (this *Message) Send() error {
	data, err := json.Marshal(this.data)
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Add("data", string(data))

	resp, err := http.PostForm("http://" + this.host, values)
	fmt.Println(resp)
	return err
}

