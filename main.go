package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pjmp/lsgo/cli"

	"github.com/cheynewallace/tabby"
	"github.com/dustin/go-humanize"
)

func main() {
	app := cli.New()

	contents, err := app.GetFileInfoSlice()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var total int64

	t := tabby.New()

	t.AddHeader("Size", "Mode", "Time", "Name")

	for _, d := range contents {
		s := d.Size()

		if app.Options.Recursive {
			if d.IsDir() {
				filepath.Walk(filepath.Join(app.Path, d.Name()), func(path string, info os.FileInfo, err error) error {

					if err != nil {
						return err
					}

					s += info.Size()

					return nil
				})
			}
		}

		total = total + s

		t.AddLine(humanize.Bytes(uint64(s)), d.Mode(), d.ModTime().Format("Mon Jan 2 3:04PM 2006"), d.Name())
	}

	t.AddLine("total", humanize.Bytes(uint64(total)))

	t.Print()
}
