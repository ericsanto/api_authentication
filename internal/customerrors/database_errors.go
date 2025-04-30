package customerrors

import "errors"

var ErrDupliateKey = errors.New("já existe")
var ErrNotFound = errors.New("")

func DupliateKeyUniqueConstraint() string {
	return "ERROR: duplicate key value violates unique constraint"
}

func IsDuplicateKeyUniqueConstraint(err error) bool {

	return errors.Is(err, ErrDupliateKey)

}
