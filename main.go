package main

import (
	"flag"

	"github.com/spudtrooper/gitversion/gitversion"
	"github.com/spudtrooper/goutil/check"
)

var (
	updateDir   = flag.String("update_dir", ".", "directory for update")
	pkg         = flag.String("pkg", "gitversion", "the go package")
	versionFlag = flag.String("version_flag", "version", "the version flag that is generated")
	tag         = flag.String("tag", "", "the new tag")
)

func realMain() error {
	if err := gitversion.Main(*updateDir, *pkg, *versionFlag, gitversion.MainTag(*tag)); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	check.Err(realMain())
}
