package repo

import "go-demo/entity"

// &entity.User{Name: "user-demo-01", Pwd: "user-demo-pwd-01"}
func (s *Struct) UpsertUser(user *entity.User) (entity.User, error) {
	if user.ID == nil {
		s.db.Create(&user)
		return *user, nil
	} else {
		s.db.Model(&user).Updates(&user)
		return *user, nil
	}
}

func (s *Struct) ReadUsers(name *string) (entity.User, error) {
	var user entity.User
	if name != nil {
		s.db.First(&user, "name = ?", name)
	}
	s.db.First(&user)
	return user, nil
}

func (s *Struct) DeleteUsers(id *string, name *string) (interface{}, error) {
	tx := s.db.Delete(&entity.User{}, "id = ? and name = ?", id, name)
	return nil, tx.Error
}

// func (s *Struct) UpsertUsers(users []entity.User) {
// 	s.db.Create(users)
// }
