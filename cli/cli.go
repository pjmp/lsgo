package cli

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
)

type CliOptions struct {
	Options flagOpts
	Path    string
}

type flagOpts struct {
	Version   func() `short:"v" long:"version" description:"print version info for lsgo"`
	Order     string `short:"o" long:"order"  default:"desc" choice:"asc" choice:"desc" description:"sort order"`
	By        string `short:"b" long:"by" default:"n" choice:"n" choice:"s" choice:"t" choice:"x" description:"sort by: name (n), size (s), time (t), extension (x)"`
	Recursive bool   `short:"R" long:"recursive" description:"print total size of subdirectories recursively"`
}

func New() CliOptions {
	var opts flagOpts

	opts.Version = func() {
		fmt.Println("lsgo v0.0.1")
		os.Exit(0)
	}

	parser := flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	args, err := parser.Parse()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var path string

	if len(args) == 0 {
		path = "."
	} else {
		path = args[0]
	}

	return CliOptions{
		Options: opts,
		Path:    path,
	}
}

func (c *CliOptions) GetFileInfoSlice() ([]os.FileInfo, error) {
	path, err := os.Stat(c.Path)

	if err != nil {
		return nil, err
	}

	var contents []os.FileInfo

	if path.IsDir() {
		dir, err := ioutil.ReadDir(c.Path)

		if err != nil {
			return nil, err
		}

		contents = dir
	} else {
		contents = []os.FileInfo{path}
	}

	return contents, nil
}
