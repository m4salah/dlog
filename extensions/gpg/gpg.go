package gpg

import (
	"flag"

	"github.com/m4salah/dlog"
)

const EXT = ".md.pgp"

var gpgId string

func init() {
	flag.StringVar(&gpgId, "gpg", "", "PGP key ID to decrypt and edit .md.pgp files using gpg. if empty encryption will be off")
	dlog.RegisterPageSource(new(encryptedPages))
	dlog.RegisterQuickCommand(quickCommands)
	dlog.Post(`/\+/gpg/encrypt/{page:.+}`, encryptHandler)
	dlog.Post(`/\+/gpg/decrypt/{page:.+}`, decryptHandler)
}
