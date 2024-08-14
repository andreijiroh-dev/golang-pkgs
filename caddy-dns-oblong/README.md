## oblong module for Caddy

[![Go Reference](https://pkg.go.dev/badge/test.svg)](https://pkg.go.dev/packages.andreijiroh.xyz/golang/caddy-dns-oblong)

THis package contains a DNS provider module for Caddy and can be used to
manage DNS records with `obl.ong` subdomains.

Module name: `dns.providers.oblong`

## Configuration

### With JSON

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/) like so:

```json
{
    "module": "acme",
    "challenges": {
        "dns": {
            "provider": {
                "name": "oblong",
                "api_token": "{env.OBLONG_API_TOKEN}",
                "subdomain": "your-subdomain"
            }
        }
    }
}
```

### With Caddyfile

```Caddyfile
tls {
    dns oblong {
        api_token {env.OBLONG_API_TOKEN}
        subdomain "your-subdomain"
    }
}
```

## Authenication

See [the associated README in the libdns package](../libdns-oblong/README.md) for
instructions on generating your own API token via the
OAuth flow.

## Troubleshooting

### Error: `timed out waiting for record to fully propagate`

Some environments may have trouble querying the `_acme-challenge` TXT record from `obl.ong``. Verify in the obl.ong pannel that the temporary record is being created.

If the record does exist, your DNS resolver may be caching an earlier response before the record was valid. You can instead configure Caddy to use an alternative DNS resolver such as [Cloudflare's official `1.1.1.1`](https://www.cloudflare.com/en-gb/learning/dns/what-is-1.1.1.1/).

Add a custom `resolver` to the [`tls` directive](https://caddyserver.com/docs/caddyfile/directives/tls):

```Caddyfile
tls {
    dns oblong {
        api_token {env.OBLONG_API_TOKEN}
        subdomain "your-subdomain"
    }
    resolvers 1.1.1.1
}
```

Or with Caddy JSON to the `acme` module: [`challenges.dns.provider.resolvers: ["1.1.1.1"]`](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/challenges/dns/resolvers/).