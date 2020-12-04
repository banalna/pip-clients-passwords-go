package build

import (
	clients1 "github.com/pip-services-users/pip-clients-passwords-go/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type PasswordsClientFactory struct {
	cbuild.Factory
}

func NewPasswordsClientFactory() *PasswordsClientFactory {
	c := &PasswordsClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	// nullClientDescriptor := cref.NewDescriptor("pip-services-sasswords", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("pip-services-sasswords", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("pip-services-passwords", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("pip-services-passwords", "client", "grpc", "*", "1.0")

	// c.RegisterType(nullClientDescriptor, clients1.NewPasswordsNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewPasswordsDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewPasswordsHttpCommandableClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewPasswordGrpcClientV1)

	return c
}
