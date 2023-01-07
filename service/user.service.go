package service

import "errors"

func (s *Struct) UserService(op string) (interface{}, error) {
	if op == "get" {
		return s.repo.ReadUsers(nil)
	}
	if op == "post" {
		return s.repo.UpsertUser(nil)
	}
	if op == "put" {
		return s.repo.UpsertUser(nil)
	}
	if op == "delete" {
		return s.repo.DeleteUsers(nil, nil)
	}
	return nil, errors.New("Not supported service")
}
