package datastruct

type BrowserInfo struct {
	Browser string
	Version string
	Count int
}

type ExportedData struct {
	Project string
	Date string
	Data []BrowserInfo
}
