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

type FileWithLineNumOptions struct {
	skipGorm   bool
	skipHelper bool
}

func getFileWithLineNumOptions(options *FileWithLineNumOptions) *FileWithLineNumOptions {
	if options == nil {
		return &FileWithLineNumOptions{}
	}
	return options
}

func getOptions(options *Options) *Options {
	if options == nil {
		return &Options{
			level:          DebugLevel,
			lineNum:        false,
			lineNumLevel:   1,
			lineNumVersion: true,
		}
	}
	return options
}
