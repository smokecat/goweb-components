package fw

import (
	"context"
	"net"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/smokecat/goweb-components/pkg/xjwt"
	"github.com/smokecat/goweb-components/pkg/xutil"
)

const (
	CtxKeyHttpRequest = "ctx_http_request"
)

func CtxWithHttpRequest(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, CtxKeyHttpRequest, req)
}

func HttpRequestFromCtx(ctx context.Context) (*http.Request, bool) {
	req, ok := xutil.CtxValue[*http.Request](ctx, CtxKeyHttpRequest)
	return req, ok
}

func RemoteAddrFromCtx(ctx context.Context) string {
	req, ok := HttpRequestFromCtx(ctx)
	if !ok {
		return ""
	}

	return httpx.GetRemoteAddr(req)
}

func HostPortFromCtx(ctx context.Context) (host string, port string) {
	host, port, err := net.SplitHostPort(RemoteAddrFromCtx(ctx))
	if err != nil {
		return "", ""
	}
	return host, port
}

func JwtDataFromCtx(ctx context.Context) (userId int64, uid string, email string, authToken string) {
	userId = xutil.CtxValueQty[int64](ctx, xjwt.JwtKeyUserId)
	uid = xutil.CtxValueQty[string](ctx, xjwt.JwtKeyUid)
	email = xutil.CtxValueQty[string](ctx, xjwt.JwtKeyEmail)
	authToken = xutil.CtxValueQty[string](ctx, xjwt.JwtKeyAuthToken)

	return userId, uid, email, authToken
}

func JwtDataUserIdFromCtx(ctx context.Context) int64 {
	return xutil.CtxValueQty[int64](ctx, xjwt.JwtKeyUserId)
}

func JwtDataUidFromCtx(ctx context.Context) string {
	return xutil.CtxValueQty[string](ctx, xjwt.JwtKeyUid)
}

func JwtDataEmailFromCtx(ctx context.Context) string {
	return xutil.CtxValueQty[string](ctx, xjwt.JwtKeyEmail)
}

func JwtDataAuthTokenFromCtx(ctx context.Context) string {
	return xutil.CtxValueQty[string](ctx, xjwt.JwtKeyAuthToken)
}
