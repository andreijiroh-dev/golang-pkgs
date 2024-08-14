package obl_ong

import (
	_ "bytes"
	"context"
	_ "context"
	_ "encoding/json"
	_ "fmt"
	_ "net/http"
	_ "net/url"

	_ "github.com/libdns/libdns"
)

const baseUrl = "https://admin.obl.ong/api/v1"

func (p *Provider) getZoneInfo(ctx context.Context, subdomain string) {
	p.z
}
