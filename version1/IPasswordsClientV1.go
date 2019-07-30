package version1

type IPasswordsClientV1 interface {
	GetPasswordInfo(correlationId string,
		userId string) (result *UserPasswordInfoV1, err error)

	SetTempPassword(correlationId string,
		userId string) (password string, err error)

	SetPassword(correlationId string, userId string, password string) error

	DeletePassword(correlationId string, userId string) error

	Authenticate(correlationId string, userId string,
		password string) (authenticated bool, err error)

	ChangePassword(correlationId string, userId string,
		oldPassword string, newPassword string) error

	ValidateCode(correlationId string, userId string,
		code string) (valid bool, err error)

	ResetPassword(correlationId string, userId string,
		code string, password string) error

	RecoverPassword(correlationId string, userId string) error
}
