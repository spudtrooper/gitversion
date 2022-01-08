package main

import (
	"flag"

	"github.com/spudtrooper/gitversion/gen"
	"github.com/spudtrooper/gitversion/gitversion"
	"github.com/spudtrooper/goutil/check"
)

var (
	updateDir   = flag.String("update_dir", ".", "directory for update")
	pkg         = flag.String("pkg", "gitversion", "the go package")
	versionFlag = flag.String("version_flag", "version", "the version flag that is generated")
	tag         = flag.String("tag", "", "the new tag")
	inc         = flag.Bool("inc", false, "simple increment the current tag, e.g. v0.0.1 -> v0.0.2")
)

func realMain() error {
	if gitversion.CheckVersionFlag() {
		return nil
	}
	if err := gen.Main(*updateDir, *pkg, *versionFlag, gen.MainTag(*tag), gen.MainIncTag(*inc)); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	check.Err(realMain())
}
