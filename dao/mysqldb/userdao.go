package mysqldb

import (
	"strings"

	"cloudfreexiao/ant-graphql/backend-go/dao/schemas"
	"cloudfreexiao/ant-graphql/backend-go/lib/logapi"
)

type MysqlUserDao struct {
}

const usertb = "user"

func (dao MysqlUserDao) CreateUser(u *schemas.User) (err error) {
	logapi.DEBUG("MysqlUserImpl CreateUser user: ", u)

	engine := get()
	sess := engine.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}
	u.Email = strings.ToLower(u.Email)
	isExist, err := sess.
		Where("email=?", u.Email).
		Get(new(schemas.User))
	if err != nil {
		return err
	} else if isExist {
		return ErrEmailAlreadyUsed{u.Email}
	}

	if _, err = sess.Insert(u); err != nil {
		return err
	}

	return sess.Commit()
}

func (dao MysqlUserDao) GetUserByEmail(email string) (*schemas.User, error) {
	if len(email) == 0 {
		return nil, ErrUserNotExist{0, email, 0}
	}

	engine := get()
	email = strings.ToLower(email)
	// First try to find the user by primary email
	user := &schemas.User{Email: email}
	has, err := engine.Get(user)
	if err != nil {
		return nil, err
	}
	if has {
		return user, nil
	}

	return nil, ErrUserNotExist{0, email, 0}
}

func (dao MysqlUserDao) UpdateUser(u *schemas.User) error {
	engine := get()
	_, err := engine.ID(u.ID).AllCols().Update(u)
	return err
}
