package resolver

import (
	"cloudfreexiao/ant-graphql/backend-go/graphql/model"
)

type addrResolver struct {
	addr *model.Address
}

func (r *addrResolver) IP() *string {
	return &r.addr.IP
}

func (r *addrResolver) Mask() *string {
	return &r.addr.Mask
}
