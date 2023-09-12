package pandoc

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark/ast"

	"github.com/m4salah/dlog"
)

var SUPPORTED_EXT = []string{".org", ".rst", ".rtf", ".odt"}
var pandoc_support bool

func init() {
	flag.BoolVar(&pandoc_support, "pandoc", false, "Use pandoc to render "+strings.Join(SUPPORTED_EXT, ", "))
	dlog.RegisterPageSource(new(pandoc))
}

type pandoc struct{}

func (p *pandoc) Page(name string) dlog.Page {
	if !pandoc_support {
		return nil
	}

	for _, ext := range SUPPORTED_EXT {
		pg := page{
			name: name,
			ext:  ext,
		}
		if pg.Exists() {
			return &pg
		}
	}

	return nil
}

func (p *pandoc) Each(ctx context.Context, f func(dlog.Page)) {
	if !pandoc_support {
		return
	}

	filepath.WalkDir(".", func(name string, d fs.DirEntry, err error) error {
		select {

		case <-ctx.Done():
			return errors.New("context stopped")

		default:
			ext := path.Ext(name)
			basename := name[:len(name)-len(ext)]

			for _, supported_ext := range SUPPORTED_EXT {
				if supported_ext == ext {
					f(&page{
						name: basename,
						ext:  ext,
					})
					break
				}
			}

		}

		return nil
	})
}

type page struct {
	name string
	ext  string
}

func (p *page) Name() string {
	return p.name
}

func (p *page) FileName() string {
	return filepath.FromSlash(p.name) + p.ext
}

func (p *page) Exists() bool {
	_, err := os.Stat(p.FileName())
	return err == nil
}

func (p *page) Render() template.HTML {
	cmd := exec.Command("pandoc", "-f", p.ext[1:], "-t", "html", p.FileName(), "--output", "-")
	html, err := cmd.Output()
	if err != nil {
		log.Printf("Couldn't parse %s to markdown using pandoc: %s", p.FileName(), err.Error())
	}

	return template.HTML(html)
}

func (p *page) Content() dlog.Markdown {
	dat, err := os.ReadFile(p.FileName())
	if err != nil {
		return ""
	}
	return dlog.Markdown(dat)
}

func (p *page) ModTime() time.Time {
	s, err := os.Stat(p.FileName())
	if err != nil {
		return time.Time{}
	}

	return s.ModTime()
}

func (p *page) Delete() bool {
	defer dlog.Trigger(dlog.AfterDelete, p)

	if p.Exists() {
		err := os.Remove(p.FileName())
		if err != nil {
			fmt.Printf("Can't delete `%s`, err: %s\n", p.Name(), err)
			return false
		}
	}
	return true
}

func (p *page) Write(content dlog.Markdown) bool {
	dlog.Trigger(dlog.BeforeWrite, p)
	defer dlog.Trigger(dlog.AfterWrite, p)

	name := p.FileName()
	os.MkdirAll(filepath.Dir(name), 0700)

	content = dlog.Markdown(strings.ReplaceAll(string(content), "\r\n", "\n"))
	if err := os.WriteFile(name, []byte(content), 0644); err != nil {
		fmt.Printf("Can't write `%s`, err: %s\n", p.Name(), err)
		return false
	}
	return true
}

func (p *page) AST() ast.Node { return ast.NewDocument() }
func (p *page) Emoji() string { return "" }
