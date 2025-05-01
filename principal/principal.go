package principal

import (
	"context"
	"strings"
)

const PrincipalKey string = "the_principal"

type Principal struct {
	Name   string
	Scopes []string
}

func (p *Principal) HasAccess(scope string) bool {
	scope = strings.TrimSpace(scope)
	if scope == "" {
		return false
	}
	for _, _scope := range p.Scopes {
		if _scope == scope {
			return true
		}
	}
	return false
}

func NewPrincipal(name, scope string) *Principal {
	scope = strings.TrimSpace(scope)
	if scope == "" {
		return &Principal{
			Name: strings.TrimSpace(name),
		}
	}
	return &Principal{
		Name:   strings.TrimSpace(name),
		Scopes: strings.Fields(scope),
	}
}

func NewContext(c context.Context, p *Principal) context.Context {
	return context.WithValue(c, PrincipalKey, p)
}

func FromContext(c context.Context) (*Principal, bool) {
	prnpal, ok := c.Value(PrincipalKey).(*Principal)
	return prnpal, ok
}

func HasAccess(c context.Context, scope string) bool {
	principal, ok := FromContext(c)
	if !ok {
		return false
	}
	return principal.HasAccess(scope)
}
