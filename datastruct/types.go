package datastruct

type ClientInfo struct {
	Ip string
	Date string
	Browser string
}

type ExportedData struct {
	Project string
	Data []ClientInfo
}
