package log

import "io"

type Options struct {
	level          Level
	output         io.Writer
	category       string
	json           bool
	lineNum        bool
	lineNumPrefix  string
	lineNumLevel   int
	lineNumSource  bool
	lineNumVersion bool
}

func getOptions(options *Options) *Options {
	if options == nil {
		return &Options{
			level:          DebugLevel,
			lineNum:        true,
			lineNumLevel:   1,
			lineNumVersion: true,
		}
	}
	return options
}
