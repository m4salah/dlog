package gpg

import (
	"errors"

	"github.com/m4salah/dlog"
)

var (
	deleteFailedErr     = errors.New("Couldn't delete original page")
	encryptionFailedErr = errors.New("Couldn't encrypt page")
)

func encryptHandler(w dlog.Response, r dlog.Request) dlog.Output {
	if dlog.READONLY {
		return dlog.Unauthorized("read only mode")
	}

	vars := dlog.Vars(r)
	p := dlog.NewPage(vars["page"])
	if !p.Exists() {
		return dlog.NotFound("page not found")
	}

	encryptedPage := page{name: p.Name()}
	if !encryptedPage.Write(p.Content()) {
		return dlog.InternalServerError(encryptionFailedErr)
	}

	if !p.Delete() {
		return dlog.InternalServerError(deleteFailedErr)
	}

	return dlog.NoContent()
}

func decryptHandler(w dlog.Response, r dlog.Request) dlog.Output {
	if dlog.READONLY {
		return dlog.Unauthorized("read only mode")
	}

	vars := dlog.Vars(r)
	p := dlog.NewPage(vars["page"])
	if !p.Exists() {
		return dlog.NotFound("page not found")
	}

	content := p.Content()
	if !p.Delete() {
		return dlog.InternalServerError(deleteFailedErr)
	}

	decryptedPage := dlog.NewPage(p.Name())
	if !decryptedPage.Write(content) {
		return dlog.InternalServerError(encryptionFailedErr)
	}

	return dlog.NoContent()
}
