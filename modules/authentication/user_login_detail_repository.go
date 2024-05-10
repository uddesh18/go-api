package authentication

type UserLoginDetailRepository interface {
	Create(user UserLoginDetail)
	GetById(userid UserLoginDetail) (user UserLoginDetail, err error)
}
