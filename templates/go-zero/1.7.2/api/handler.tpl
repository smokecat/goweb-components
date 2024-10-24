package {{.PkgName}}

import (
	"context"
	"net/http"

	"github.com/smokecat/goweb-components/pkg/framework/go-zero/fw"
	"github.com/smokecat/goweb-components/pkg/xvalidator"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest/httpx"

	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasResp}}var resp any
		{{end}}

		// Inject request into the context
		ctx := fw.CtxWithHttpRequest(r.Context(), r)
		r = r.WithContext(ctx)

		err := fw.WithNewTx(r.Context(), func(ctx context.Context, conn sqlx.SqlConn) error {
			{{if .HasRequest}}var req types.{{.RequestType}}
			err := httpx.Parse(r, &req)
			if err != nil {
				return xvalidator.WrapErr(err)
			}

			// validate by validator
			err = xvalidator.Default().Struct(req)
			if err != nil {
				return xvalidator.WrapErr(err)
			}

			{{end}}l := {{.LogicName}}.New{{.LogicType}}(ctx, svcCtx)
			{{if .HasResp}}resp, {{end}}err = l.{{.Call}}({{if .HasRequest}}&req{{end}})
			return err
		})

		fw.WriteJsonResponse(ctx, w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
	}
}
