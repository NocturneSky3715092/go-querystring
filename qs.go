// Package qs parses/builds URL query strings, using only the standard library.
package qs

import (
	"net/url"
	"strings"
)

func Parse(qs string) map[string]string {
	out := map[string]string{}
	for _, pair := range strings.Split(strings.TrimPrefix(qs, "?"), "&") {
		if pair == "" {
			continue
		}
		k, v, _ := strings.Cut(pair, "=")
		dk, _ := url.QueryUnescape(k)
		dv, _ := url.QueryUnescape(v)
		out[dk] = dv
	}
	return out
}
