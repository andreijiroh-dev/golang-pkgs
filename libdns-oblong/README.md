# `obl.ong` for [`libdns`](https://github.com/libdns/libdns)

> Built as part of [Hack Club Arcade](https://hackclub.com/arcade)

[![Go Reference](https://pkg.go.dev/badge/test.svg)](https://pkg.go.dev/packages.andreijiroh.xyz/golang/libdns-oblong)

This package implements the [libdns interfaces](https://github.com/libdns/libdns) for
[`obl.ong` subdomains](https://obl.ong) using its API, allowing you to manage DNS records programmatically.

## Caveats

* To generate an API token, one must [request an OAuth client](https://admin.obl.ong/developers/applications),
get your request approved and go through the OAuth flow. You can also opt to generate one via
[this website](https://libdns.obl.ong/get-token) (coming soon).
* To use this library, your `obl.ong` subdomain must be approved for use.
