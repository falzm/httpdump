package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/zenazn/goji"
)

var (
	version   string
	buildDate string

	flagHelp    bool
	flagVersion bool
)

func init() {
	flag.BoolVar(&flagHelp, "h", false, "display this help and exit")
	flag.BoolVar(&flagVersion, "version", false, "display version and exit")
	flag.Usage = func() { printUsage(os.Stderr) }

	flag.Parse()
}

func main() {
	if flagHelp {
		printUsage(os.Stdout)
		os.Exit(0)
	} else if flagVersion {
		printVersion(version, buildDate)
		os.Exit(0)
	}

	goji.Handle("/*", dump)
	goji.Serve()

	os.Exit(0)
}

func printUsage(output io.Writer) {
	fmt.Fprintf(output, "Usage: httpdump [OPTIONS]")
	fmt.Fprint(output, "\n\nOptions:\n")

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(output, "   -%s  %s\n", f.Name, f.Usage)
	})

	os.Exit(2)
}

func printVersion(version, buildDate string) {
	fmt.Printf("%s version %s, built on %s\nGo version: %s (%s)\n",
		path.Base(os.Args[0]),
		version,
		buildDate,
		runtime.Version(),
		runtime.Compiler,
	)
}

func dieOnError(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s", format), a)
	os.Exit(1)
}
