package netbox_dns

// From https://github.com/netbox-community/go-netbox/blob/18aea9a4ac2c7a995e1ed15baa105f9032b4239e/netbox/netbox.go

import (
	"fmt"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ProtonMail/go-netbox-dns/netbox_dns/client"
)

// NewNetboxAt returns a client which will connect to the given
// hostname, which can optionally include a port, e.g. localhost:8000
func NewNetboxDNSAt(host string) *client.NetBoxAPI {
	t := client.DefaultTransportConfig().WithHost(host)
	return client.NewHTTPClientWithConfig(strfmt.Default, t)
}

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

// NewNetboxWithAPIKey returns a client which will connect to the given
// hostname (and optionally port), and will set the expected Authorization
// header on each request
func NewNetboxDNSWithAPIKey(host string, apiToken string) *client.NetBoxAPI {
	t := runtimeclient.New(host, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	return client.New(t, strfmt.Default)
}
