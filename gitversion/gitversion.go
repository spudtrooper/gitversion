package gitversion

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	goutilio "github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/goutil/or"
)

func writeGitversion(outFile, pkg, versionFlag string) error {
	var buf bytes.Buffer
	tmpl := `
package {{.Pkg}}

import (
	"flag"
	"fmt"
)

var (
	{{.VersionFlag}} = flag.Bool("{{.VersionFlag}}", false, "print version")
)

// Prints the big version if --{{.VersionFlag}} is set and returns true, otherwise reutrns false
func CheckVersionFlag() bool {
	if *{{.VersionFlag}} {
		fmt.Printf("Version: %s\n", theGitVersion)
		return true
	}
	return false
}		
`
	if err := renderTemplate(&buf, tmpl, "gitversion.go", struct {
		VersionFlag string
		Pkg         string
	}{
		VersionFlag: versionFlag,
		Pkg:         pkg,
	}); err != nil {
		return err
	}

	if err := writeFile(outFile, buf); err != nil {
		return err
	}
	return nil
}
func writeGitversionTest(outFile, pkg, versionFlag string) error {
	var buf bytes.Buffer
	tmpl := `
package {{.Pkg}}

import "testing"

func TestCheckVersionFlagTrue(t *testing.T) {
	*{{.VersionFlag}} = true
	if !CheckVersionFlag() {
		t.Fatalf("expected true, got false")
	}
}

func TestCheckVersionFlagFalse(t *testing.T) {
	*{{.VersionFlag}} = false
	if CheckVersionFlag() {
		t.Fatalf("expected false, got true")
	}
}		
`
	if err := renderTemplate(&buf, tmpl, "gitversion.go", struct {
		VersionFlag string
		Pkg         string
	}{
		VersionFlag: versionFlag,
		Pkg:         pkg,
	}); err != nil {
		return err
	}

	if err := writeFile(outFile, buf); err != nil {
		return err
	}
	return nil
}

func writeThegitversion(outFile, pkg, tag string) error {
	var buf bytes.Buffer
	tmpl := `
package {{.Pkg}}

const theGitVersion = "{{.Tag}}"		
`
	if err := renderTemplate(&buf, tmpl, "thegitversion.go", struct {
		Pkg string
		Tag string
	}{
		Pkg: pkg,
		Tag: tag,
	}); err != nil {
		return err
	}

	if err := writeFile(outFile, buf); err != nil {
		return err
	}
	return nil
}

func writeFile(outFile string, buf bytes.Buffer) error {
	if err := ioutil.WriteFile(outFile, buf.Bytes(), 0755); err != nil {
		return err
	}
	log.Printf("wrote to %s", outFile)
	return nil
}

func Main(dir, pkg, versionFlag string, mOpts ...MainOption) error {
	opts := MakeMainOptions(mOpts...)

	if opts.Tag() != "" {
		if output, err := run(dir, "git", "tag", "-a", opts.Tag(), "-m", fmt.Sprintf("updating tag to %s", opts.Tag())); err != nil {
			return err
		} else {
			fmt.Print(output)
		}
	}

	tag, err := run(dir, "git", "describe", "--tags")
	if err != nil {
		return err
	}
	tag = strings.TrimSpace(tag)
	log.Printf("have tag: %s", tag)

	versionFlag = or.String(versionFlag, "version")
	pkg = or.String(pkg, "gitversion")
	outDir, err := goutilio.MkdirAll(dir, pkg)
	if err != nil {
		return err
	}

	thegitversion := path.Join(outDir, "thegitversion.go")
	if err := writeThegitversion(thegitversion, pkg, tag); err != nil {
		return err
	}
	gitversion := path.Join(outDir, "gitversion.go")
	if err := writeGitversion(gitversion, pkg, versionFlag); err != nil {
		return err
	}
	gitversionTest := path.Join(outDir, "gitversion_test.go")
	if err := writeGitversionTest(gitversionTest, pkg, versionFlag); err != nil {
		return err
	}
	if output, err := run(dir, "go", "fmt", thegitversion, gitversion, gitversionTest); err != nil {
		return err
	} else {
		fmt.Print(output)
	}
	if output, err := run(dir, "go", "test", thegitversion, gitversion, gitversionTest); err != nil {
		return err
	} else {
		fmt.Print(output)
	}

	fmt.Println(`


*** Now add the following to your main:

	if gitversion.CheckVersionFlag() {
		return nil
	}

`)

	return nil
}

func run(dir, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return stdout.String(), nil
}

func renderTemplate(buf io.Writer, t string, name string, data interface{}) error {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return err
	}
	if err := tmpl.Execute(buf, data); err != nil {
		return err
	}
	return nil
}
