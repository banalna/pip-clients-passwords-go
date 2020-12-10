package version1

import (
	"reflect"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cclients "github.com/pip-services3-go/pip-services3-rpc-go/clients"
)

type PasswordsHttpCommandableClientV1 struct {
	*cclients.CommandableHttpClient
	userPasswordInfoV1Type reflect.Type
	mapType                reflect.Type
}

func NewPasswordsHttpCommandableClientV1() *PasswordsHttpCommandableClientV1 {
	c := &PasswordsHttpCommandableClientV1{
		CommandableHttpClient:  cclients.NewCommandableHttpClient("v1/passwords"),
		userPasswordInfoV1Type: reflect.TypeOf(&UserPasswordInfoV1{}),
		mapType:                reflect.TypeOf(make(map[string]bool)),
	}
	return c
}

func (c *PasswordsHttpCommandableClientV1) GetPasswordInfo(correlationId string,
	userId string) (result *UserPasswordInfoV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	res, err := c.CallCommand(c.userPasswordInfoV1Type, "get_password_info", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*UserPasswordInfoV1)
	return result, nil
}

func (c *PasswordsHttpCommandableClientV1) SetTempPassword(correlationId string,
	userId string) (password string, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	res, err := c.CallCommand(nil, "set_temp_password", correlationId, params)
	if err != nil {
		return "", err
	}

	result, _ := res.(string)
	return result, nil
}

func (c *PasswordsHttpCommandableClientV1) SetPassword(correlationId string, userId string, password string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"password", password,
	)

	_, err := c.CallCommand(nil, "set_password", correlationId, params)
	return err
}

func (c *PasswordsHttpCommandableClientV1) DeletePassword(correlationId string, userId string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	_, err := c.CallCommand(nil, "delete_password", correlationId, params)
	return err
}

func (c *PasswordsHttpCommandableClientV1) Authenticate(correlationId string, userId string,
	password string) (authenticated bool, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"password", password,
	)

	res, err := c.CallCommand(c.mapType, "authenticate", correlationId, params)
	if err != nil {
		return false, err
	}

	result, ok := res.(*map[string]bool)
	if !ok {
		return false, nil
	}

	val, valOk := (*result)["authenticated"]
	if !valOk {
		return false, nil
	}
	return val, nil
}

func (c *PasswordsHttpCommandableClientV1) ChangePassword(correlationId string, userId string,
	oldPassword string, newPassword string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"old_password", oldPassword,
		"new_password", newPassword,
	)

	_, err := c.CallCommand(nil, "change_password", correlationId, params)
	return err
}

func (c *PasswordsHttpCommandableClientV1) ValidateCode(correlationId string, userId string,
	code string) (valid bool, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"code", code,
	)

	res, err := c.CallCommand(c.mapType, "validate_code", correlationId, params)
	if err != nil {
		return false, err
	}

	result, ok := res.(map[string]bool)
	if !ok {
		return false, nil
	}

	val, valOk := result["valid"]
	if !valOk {
		return false, nil
	}
	return val, nil
}

func (c *PasswordsHttpCommandableClientV1) ResetPassword(correlationId string, userId string,
	code string, password string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"code", code,
		"password", password,
	)
	_, err := c.CallCommand(nil, "reset_password", correlationId, params)
	return err
}

func (c *PasswordsHttpCommandableClientV1) RecoverPassword(correlationId string, userId string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)
	_, err := c.CallCommand(nil, "recover_password", correlationId, params)
	return err

}
