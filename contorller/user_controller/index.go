package user_controller

func (s *Struct) ReadUserController(req interface{}) (res interface{}, err error) {
	return s.Service.UserService("get")
}
func (s *Struct) CreateUserController(req interface{}) (res interface{}, err error) {
	return s.Service.UserService("post")
}
func (s *Struct) UpdateUserController(req interface{}) (res interface{}, err error) {
	return s.Service.UserService("put")
}
func (s *Struct) DeleteUserController(req interface{}) (res interface{}, err error) {
	return s.Service.UserService("delete")
}
