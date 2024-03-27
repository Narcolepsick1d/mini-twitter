package scope

import (
	"Narcolepsick1d/mini-twitter/pkg/hash"
	"github.com/doug-martin/goqu/v9"
)

// Dependencies is meta object which contains db connect, api client etc.
type Dependencies struct {
	DB     *goqu.Database
	Hash   *hash.SHA1Hasher
	Secret string
}
