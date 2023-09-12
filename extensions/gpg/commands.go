package gpg

import (
	"fmt"
	"html/template"
	"net/url"
	"path"

	"github.com/m4salah/dlog"
)

const decryptableExt = ".pgp"

func quickCommands(p dlog.Page) []dlog.Command {
	if len(gpgId) == 0 {
		return nil
	}

	if path.Ext(p.FileName()) == decryptableExt {
		return []dlog.Command{
			&decryptCommand{page: p},
		}
	} else {
		return []dlog.Command{
			&encryptCommand{page: p},
		}
	}
}

type encryptCommand struct {
	page dlog.Page
}

func (e *encryptCommand) Icon() string         { return "fa-solid fa-lock" }
func (e *encryptCommand) Name() string         { return "Make private" }
func (e *encryptCommand) Link() string         { return "" }
func (e *encryptCommand) OnClick() template.JS { return "encrypt(event)" }
func (e *encryptCommand) Widget() template.HTML {
	action := "/+/gpg/encrypt/" + url.PathEscape(e.page.Name())
	return template.HTML(fmt.Sprintf(`
	  <script>
	  function encrypt(event) {
		 event.preventDefault();

		 const data = new FormData()
		 data.append('csrf', document.querySelector('input[name=csrf]').value);

		 let method = 'POST'

		 fetch("%s", {method: method, body: data})
			 .then( () => location.reload() );
	  }
	  </script>
`, action))
}

type decryptCommand struct {
	page dlog.Page
}

func (e *decryptCommand) Icon() string         { return "fa-solid fa-lock-open has-text-danger" }
func (e *decryptCommand) Name() string         { return "Make public" }
func (e *decryptCommand) Link() string         { return "" }
func (e *decryptCommand) OnClick() template.JS { return "decrypt(event)" }
func (e *decryptCommand) Widget() template.HTML {
	action := "/+/gpg/decrypt/" + url.PathEscape(e.page.Name())
	return template.HTML(fmt.Sprintf(`
	  <script>
	  function decrypt(event) {
		 event.preventDefault();

		 const data = new FormData()
		 data.append('csrf', document.querySelector('input[name=csrf]').value);

		 let method = 'POST'

		 fetch("%s", {method: method, body: data})
			 .then( () => location.reload() );
	  }
	  </script>
`, action))
}
