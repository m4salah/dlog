package embed

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/m4salah/dlog"
	"github.com/m4salah/dlog/extensions/shortcode"
)

func init() {
	shortcode.ShortCode("embed", embedShortcode)
}

func embedShortcode(in dlog.Markdown) template.HTML {
	p := dlog.NewPage(strings.TrimSpace(string(in)))
	if p == nil || !p.Exists() {
		return template.HTML(fmt.Sprintf("Page: %s doesn't exist", in))
	}

	return p.Render()
}
