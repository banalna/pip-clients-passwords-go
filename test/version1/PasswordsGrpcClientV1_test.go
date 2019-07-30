package test_version1

import (
	"os"
	"testing"

	"github.com/pip-services-users/pip-clients-passwords-go/version1"
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

var client *version1.PasswordGrpcClientV1
var fixture *PasswordsClientFixtureV1

func setup(t *testing.T) *PasswordsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	client = version1.NewPasswordGrpcClientV1()
	client.Configure(httpConfig)
	client.Open("")

	fixture = NewPasswordsClientFixtureV1(client)

	return fixture
}

func teardown(t *testing.T) {
	client.Close("")
}

func TestRecoverPassword(t *testing.T) {
	fixture := setup(t)
	defer teardown(t)

	fixture.TestRecoverPassword(t)
}

func TestChangePassword(t *testing.T) {
	fixture := setup(t)
	defer teardown(t)

	fixture.TestChangePassword(t)
}

func TestSigninWithWrongPassword(t *testing.T) {
	fixture := setup(t)
	defer teardown(t)

	fixture.TestSigninWithWrongPassword(t)
}
