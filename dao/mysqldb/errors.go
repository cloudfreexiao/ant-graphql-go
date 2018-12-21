package mysqldb

import "fmt"

type ErrEmailAlreadyUsed struct {
	Email string
}

func (err ErrEmailAlreadyUsed) Error() string {
	return fmt.Sprintf("e-mail has been used [email: %s]", err.Email)
}

type ErrUserNotExist struct {
	UID   int64
	Name  string
	KeyID int64
}

func (err ErrUserNotExist) Error() string {
	return fmt.Sprintf("user does not exist [uid: %d, name: %s, keyid: %d]", err.UID, err.Name, err.KeyID)
}
