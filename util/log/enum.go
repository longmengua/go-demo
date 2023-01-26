package log

var color = struct {
	red    string
	green  string
	yellow string
	blue   string
}{
	// color code reference: https://en.wikipedia.org/wiki/ANSI_escape_code
	red:    "31",
	green:  "32",
	yellow: "33",
	blue:   "34",
}

var level = struct {
	debug string
	info  string
	warn  string
	err   string
}{
	debug: "debug",
	info:  "info",
	warn:  "warn",
	err:   "err",
}
