package app

import "proto-example/app/proto"

func NewAccount() *proto.Account {
	return &proto.Account{
		Id:         1,
		Name:       "Abhishek Ghosh",
		IsVerified: true,
	}
}
