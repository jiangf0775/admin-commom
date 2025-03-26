package {{.pkgName}}

import (
    "common/web"
	{{.imports}}
)

type {{.logic}} struct {
    webs.LogicResult
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) (resp *web.Result, err error) {
	// todo: add your logic here and delete this line

	return l.ResultMsg("", web.DBError, err)
}
