package test_version1

import (
	"os"
	"testing"

	"github.com/pip-services-users/pip-clients-passwords-go/version1"
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type passwordsHttpCommandableClientV1Test struct {
	client  *version1.PasswordsHttpCommandableClientV1
	fixture *PasswordsClientFixtureV1
}

func newPasswordsHttpCommandableClientV1Test() *passwordsHttpCommandableClientV1Test {
	return &passwordsHttpCommandableClientV1Test{}
}

func (c *passwordsHttpCommandableClientV1Test) setup(t *testing.T) *PasswordsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewPasswordsHttpCommandableClientV1()
	c.client.Configure(httpConfig)
	c.client.Open("")

	c.fixture = NewPasswordsClientFixtureV1(c.client)

	return c.fixture
}

func (c *passwordsHttpCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close("")
}

func TestHttpRecoverPassword(t *testing.T) {
	c := newPasswordsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestRecoverPassword(t)
}

func TestHttpChangePassword(t *testing.T) {
	c := newPasswordsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestChangePassword(t)
}

func TestHttpSigninWithWrongPassword(t *testing.T) {
	c := newPasswordsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSigninWithWrongPassword(t)
}
