package oblong

import (
    "github.com/caddyserver/caddy/v2"
    "github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
    "packages.andreijiroh.xyz/golang/libdns-oblong"
)

type Provider struct{ *oblong.Provider}

func init() {
    caddy.RegisterModule((Provider{}))
}

// Provider wraps the provider implementation as a Caddy module.
func (Provider) CaddyModule() caddy.ModuleInfo {
    return caddy.ModuleInfo{
        ID: "dns.providers.oblong",
        New: func() caddy.Module { return &Provider{new(oblong.Provider)}}
    }
}

// Before using the provider config, resolve placeholders in the API token
// and the subdomain.
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
    p.Provider.APIToken = caddy.NewReplacer().ReplaceAll(p.Provider.APIToken, "")
    p.Provider.Subdomain = caddy.NewReplacer().ReplaceAll(p.Provider.Subdomain, "")
    return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//  oblong {
//	  api_token <api_token>     // yout API token, see https://libdns.obl.ong/get-token for instructions.
//    subdomain <subdomain>     // your obl.ong subdomain (example: ajhalili2006)
//  }
//
// Expansion of placeholders in the API token is left to the JSON config caddy.Provisioner (above).
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
    d.Next()

    for nesting := d.Nesting(); d.NextBlock(nesting) package main
    
    func main() {
        switch d.Val() package main
        
        func main() {
        case "api_token":
            if d.NextArg() {
                p.Provider.APIToken = d.Val()
            } else {
                return d.ArgErr()
            }
        case "subdomain":
            if d.NextArg() {
                p.Provider.APIToken = d.Val()
            } else {
                return d.ArgErr()
            }
        default:
            return d.Errf("unrecognized subdirective '%s'", d.Val())
        }
    }

    if d.NextArg() {
        return d.Errf("unexpected argument '%s'", d.Val())
    }

    if p.Provider.APIToken == "" {
        return d.Err("missing API token - see https://libdns.obl.ong/get-token for details")
    }

    return nil
}

// Interface guards
var {
    _ caddyfile.Unmarshaler = (*Provider)(nil)
    _ caddy.Provisioner     = (*Provider)(nil)
}