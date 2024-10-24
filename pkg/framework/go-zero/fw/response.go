package fw

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/smokecat/goweb-components/pkg/xcode"
	"github.com/smokecat/goweb-components/pkg/xerr"
	"github.com/smokecat/goweb-components/pkg/xutil"
)

var emptyOkResp = map[string]string{
	"code": "OK",
	"msg":  "OK",
}

type baseRes interface {
	IsEmpty() bool
	GetCode() string
	GetMsg() string
	SetCode(code string)
	SetMsg(msg string)
}

type skipWrapRes interface {
	IsSkipWrapRes() bool
}

func WriteJsonResponse(ctx context.Context, w http.ResponseWriter, v any, err error) {
	// TODO: write headers

	writeResponse(ctx, w, v, err)
}

func writeResponse(ctx context.Context, w http.ResponseWriter, resp any, err error) {
	if err != nil {
		httpx.ErrorCtx(ctx, w, err)
	} else {
		httpx.OkJsonCtx(ctx, w, resp)
	}
}

func OkResponseHandler(ctx context.Context, v any) any {
	if xutil.IsNil(v) {
		return emptyOkResp
	}

	if r, ok := v.(skipWrapRes); ok && r.IsSkipWrapRes() {
		return v
	}

	if r, ok := v.(baseRes); ok {
		if r.IsEmpty() {
			r.SetCode("OK")
			r.SetMsg("OK")
		}
		return v
	}

	return map[string]any{
		"code": "OK",
		"msg":  "OK",
		"data": v,
	}
}

func ErrHandler(ctx context.Context, err error) (int, any) {
	var xe xerr.XErr
	if err == nil {
		xe = xerr.New(xcode.WithCode(xcode.CodeNil, http.StatusInternalServerError), "")
	} else if e, ok := err.(xerr.XErr); ok {
		if !e.IsTypePublic() {
			// private error
			if e.Meta() == nil {
				logc.Errorf(ctx, "Private error: %+v", e.Error())
			} else {
				logc.Errorf(ctx, "Private error: %+v meta=%#v", e.Error(), e.Meta())
			}
			xe = xerr.Wrap(e.Err(), xcode.CodeInternalServerError, "Unexpected internal error")
		} else if e.Code().IsNil() || e.Code().IsOK() {
			xe = xerr.New(xcode.CodeOK, "")
		} else {
			xe = e
		}
	} else {
		xe = xerr.Wrap(err, xcode.CodeInternalServerError, "Unexpected error")
	}

	logc.Errorf(ctx, "Public error: %+v", xe.Err())

	resp := map[string]any{
		"code": xe.Code().Text(),
		"msg":  xe.Error(),
	}
	if xe.Meta() != nil {
		resp["data"] = xe.Meta()
	}

	return xe.Code().HttpStatus(), resp
}
