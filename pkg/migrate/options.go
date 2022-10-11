package migrate

import (
	"context"
	"embed"
	"github.com/ppxb/go-fiber/pkg/utils"
)

type Options struct {
	ctx         context.Context
	driver      string
	uri         string
	lockName    string
	before      func(ctx context.Context) error
	changeTable string
	fs          embed.FS
	fsRoot      string
}

func WithCtx(ctx context.Context) func(*Options) {
	return func(options *Options) {
		if !utils.InterfaceIsNil(ctx) {
			getOptions(options).ctx = ctx
		}
	}
}

func WithDriver(s string) func(*Options) {
	return func(options *Options) {
		getOptions(options).driver = s
	}
}

func WithUri(s string) func(*Options) {
	return func(options *Options) {
		getOptions(options).uri = s
	}
}

func WithLockName(s string) func(*Options) {
	return func(options *Options) {
		getOptions(options).lockName = s
	}
}

func WithBefore(f func(ctx context.Context) error) func(*Options) {
	return func(options *Options) {
		if f != nil {
			getOptions(options).before = f
		}
	}
}

func WithChangeTable(s string) func(*Options) {
	return func(options *Options) {
		getOptions(options).changeTable = s
	}
}

func WithFs(fs embed.FS) func(*Options) {
	return func(options *Options) {
		getOptions(options).fs = fs
	}
}

func WithFsRoot(s string) func(*Options) {
	return func(options *Options) {
		getOptions(options).fsRoot = s
	}
}

func getOptions(options *Options) *Options {
	if options == nil {
		return &Options{
			driver:      "mysql",
			lockName:    "MigrationLock",
			changeTable: "schema_migrations",
		}
	}
	return options
}
