package version1

import (
	"github.com/pip-services-users/pip-clients-passwords-go/protos"
	"github.com/pip-services3-go/pip-services3-grpc-go/clients"
)

type PasswordGrpcClientV1 struct {
	clients.GrpcClient
}

func NewPasswordGrpcClientV1() *PasswordGrpcClientV1 {
	return &PasswordGrpcClientV1{
		GrpcClient: *clients.NewGrpcClient("passwords_v1.Passwords"),
	}
}

func (c *PasswordGrpcClientV1) GetPasswordInfo(correlationId string,
	userId string) (result *UserPasswordInfoV1, err error) {
	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordInfoReply)
	err = c.Call("get_password_info", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toPasswordInfo(reply.Info)

	return result, nil
}

func (c *PasswordGrpcClientV1) SetTempPassword(correlationId string,
	userId string) (password string, err error) {
	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordValueReply)
	err = c.Call("set_temp_password", correlationId, req, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return "", err
	}

	password = reply.Password

	return password, nil
}

func (c *PasswordGrpcClientV1) SetPassword(correlationId string,
	userId string, password string) error {
	req := &protos.PasswordIdAndValueRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Password:      password,
	}

	reply := new(protos.PasswordEmptyReply)
	err := c.Call("set_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) DeletePassword(correlationId string,
	userId string) error {
	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordEmptyReply)
	err := c.Call("delete_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) Authenticate(correlationId string,
	userId string, password string) (authenticated bool, err error) {
	req := &protos.PasswordIdAndValueRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Password:      password,
	}

	reply := new(protos.PasswordAuthenticateReply)
	err = c.Call("authenticate", correlationId, req, reply)
	if err != nil {
		return false, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return false, err
	}

	authenticated = reply.Authenticated

	return authenticated, nil
}

func (c *PasswordGrpcClientV1) ChangePassword(correlationId string,
	userId string, oldPassword string, newPassword string) error {
	req := &protos.PasswordIdAndValuesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		OldPassword:   oldPassword,
		NewPassword:   newPassword,
	}

	reply := new(protos.PasswordEmptyReply)
	err := c.Call("change_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) ValidateCode(correlationId string,
	userId string, code string) (valid bool, err error) {
	req := &protos.PasswordIdAndCodeRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Code:          code,
	}

	reply := new(protos.PasswordValidReply)
	err = c.Call("validate_code", correlationId, req, reply)
	if err != nil {
		return false, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return false, err
	}

	valid = reply.Valid

	return valid, nil
}

func (c *PasswordGrpcClientV1) ResetPassword(correlationId string,
	userId string, code string, password string) error {
	req := &protos.PasswordIdAndCodeAndValueRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Code:          code,
		Password:      password,
	}

	reply := new(protos.PasswordEmptyReply)
	err := c.Call("reset_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *PasswordGrpcClientV1) RecoverPassword(correlationId string,
	userId string) error {
	req := &protos.PasswordIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.PasswordEmptyReply)
	err := c.Call("recover_password", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}
