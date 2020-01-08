package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type key int

const userIpKey key = 0

func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip :%q is not IP:port", req.RemoteAddr)
	}
	return net.ParseIP(ip), err
}

func NewContext(ctx context.Context, userIp net.IP) context.Context {
	return context.WithValue(ctx, userIpKey, userIp)
}

func FromContext(ctx context.Context) (net.IP, bool) {
	userIP, ok := ctx.Value(userIpKey).(net.IP)
	return userIP, ok
}
