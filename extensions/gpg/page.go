package gpg

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/m4salah/dlog"
	emojiAst "github.com/yuin/goldmark-emoji/ast"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type page struct {
	name string
	ast  ast.Node
}

func (p *page) Name() string     { return p.name }
func (p *page) FileName() string { return filepath.FromSlash(p.name) + EXT }

func (p *page) Exists() bool {
	_, err := os.Stat(p.FileName())
	return err == nil
}

func (p *page) Render() template.HTML {
	content := p.Content()
	content = dlog.PreProcess(content)
	var buf bytes.Buffer
	if err := dlog.MarkDownRenderer.Convert([]byte(content), &buf); err != nil {
		return template.HTML(err.Error())
	}

	return template.HTML(buf.String())
}

func (p *page) Content() dlog.Markdown {
	cmd := exec.Command("gpg", "--decrypt", p.FileName())
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Coudln't decrypt file: %s, err: %s", p.FileName(), err.Error())
	}

	return dlog.Markdown(out)
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
	cmd := exec.Command("gpg", "-r", gpgId, "--output", p.FileName(), "--batch", "--yes", "--encrypt")
	cmd.Stdin = bytes.NewBuffer([]byte(content))

	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Can't write `%s`, out: %s, err: %s\n", p.Name(), out, err)
		return false
	}

	return true
}

func (p *page) AST() ast.Node {
	if p.ast == nil {
		p.ast = dlog.MarkDownRenderer.Parser().Parse(text.NewReader([]byte(p.Content())))
	}

	return p.ast
}
func (p *page) Emoji() string {
	if e, ok := dlog.FindInAST[*emojiAst.Emoji](p.AST()); ok {
		return string(e.Value.Unicode)
	}

	return ""
}
