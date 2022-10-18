package ms

import (
	"context"
	"embed"
	"fmt"
	"github.com/ppxb/go-fiber/pkg/log"
	"os"
)

type ConfBox struct {
	Ctx context.Context
	Fs  embed.FS
	Dir string
}

func (c ConfBox) Get(filename string) (bs []byte) {
	if filename == "" {
		return
	}
	f := fmt.Sprintf("%s%s%s", c.Dir, string(os.PathSeparator), filename)
	var err error
	bs, err = os.ReadFile(f)
	if err != nil {
		log.WithContext(c.Ctx).WithError(err).Warn("[configs box]read file %s from system failed", f)
		err = nil
	}
	if len(bs) == 0 {
		bs, err = c.Fs.ReadFile(f)
		if err != nil {
			log.WithContext(c.Ctx).WithError(err).Warn("[configs box]read file %s from embed failed", f)
		}
	}
	return
}
