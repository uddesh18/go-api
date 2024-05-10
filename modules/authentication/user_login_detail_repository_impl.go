package authentication

import "gorm.io/gorm"

type UserLoginDetailRepositoryImpl struct {
	Db *gorm.DB
}

// Create implements UserLoginDetailRepository.
func (u *UserLoginDetailRepositoryImpl) Create(user UserLoginDetail) {
	panic("unimplemented")
}

// GetById implements UserLoginDetailRepository.
func (u *UserLoginDetailRepositoryImpl) GetById(userid UserLoginDetail) (user UserLoginDetail, err error) {
	panic("unimplemented")
}

func NewUserLoginDetailRepositoryImpl(Db *gorm.DB) UserLoginDetailRepository {
	return &UserLoginDetailRepositoryImpl{Db: Db}
}
