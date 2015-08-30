package logparser

const IP_ADDRESS = "ipaddress"
const DATE = "date"
const TIME = "time"
const METHOD = "method"
const PROTOCOL = "protocol"
const URL = "url"
const REFERER = "referer"
const HTTP_VERSION = "httpver"
const STATUS = "status_code"
const BYTES_SENT = "bytes_sent"
const USERAGENT = "useragent"
const REMOTE_USER = `remote_user`

const IP_ADDRESS_FORMAT = `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`
const DATE_FORMAT = `[0-9]{2}\/[A-Za-z]{3}\/[0-9]{4}`
const TIME_FORMAT = `[0-9]{2}:[0-9]{2}:[0-9]{2} [\+|\-][0-9]{4}`;
const HTTP_VERSIONS = `HTTP\/[0-9]\.[0-9]`
const METHODS = `GET|POST|PUT|HEAD|DELETE|OPTIONS|PROPFIND`


const REGEX_STRING = `^(?P<` + IP_ADDRESS + `>` + IP_ADDRESS_FORMAT + `) - (?P<` + REMOTE_USER + `>.*) \[(?P<` + DATE + `>` + DATE_FORMAT + `):(?P<` + TIME + `>` + TIME_FORMAT + `)\] "(?P<` + METHOD + `>` + METHODS + `) (?P<`+ URL + `>.+) (?P<` + HTTP_VERSION + `>` + HTTP_VERSIONS + `)" (?P<` + STATUS + `>[0-9]{3}) (?P<` + BYTES_SENT +`>[0-9]+) "(?P<` + REFERER + `>\-|.+)" "(?P<` + USERAGENT + `>.+)" "(` + IP_ADDRESS_FORMAT + `|-)"$`
