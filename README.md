# go-netbox-dns

[![GoDoc](http://godoc.org/github.com/proton/go-netbox-dns?status.svg)](http://godoc.org/github.com/proton/go-netbox-dns)

This package is largely inspired by [go-netbox](https://raw.githubusercontent.com/netbox-community/go-netbox)
but providing a client for the [netbox-dns](https://github.com/auroraresearchlab/netbox-dns) plugin API.

This package has been generated with Netbox 3.4.2 and netbox-dns 0.16.1.


## Using the client

The `github.com/proton/go-netbox-dns/netbox_dns` package has some convenience functions for creating clients with the most common
configurations you are likely to need while connecting to NetBox. `NewNetboxDNSAt` allows you to specify a hostname
(including port, if you need it), and `NewNetboxDNSWithAPIKey` allows you to specify both a hostname:port and API token.
```golang
import (
    "github.com/proton/go-netbox-dns/netbox_dns"
)
...
    c := netbox_dns.NewNetboxDNSAt("your.netbox.host:8000")
    // OR
    c := netbox_dns.NewNetboxDNSWithAPIKey("your.netbox.host:8000", "your_netbox_token")
```

If you specify the API key, you do not need to pass an additional `authInfo` to operations that need authentication, and
can pass `nil`:
```golang
    c.Plugin.PluginsNetboxDNSZonesCreate(createRequest, nil)
```

## Regenerating the clients

You can download a swagger.json (available via `/api/docs/?format=openapi`) and run
`make clean generate`.
