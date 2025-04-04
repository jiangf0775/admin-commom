package {{.pkgName}}

import (
    "common/web"
	{{.imports}}
)

type {{.logic}} struct {
    web.LogicResult
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

/** TODO 返回类型：
 * 新增的方法 返回类型 ( InsertResult, error )
 * 更新的方法 返回类型 ( UpdateResult, error )
 * 删除的方法 返回类型 ( DeleteResult, error )
 * 分页的方法 返回类型 ( PageResult, error )
 * 其他方法 应返回Model方法对应的类型！
 */
func (l *{{.logic}}) {{.function}}({{.request}}) (*web.Result, error) {
	// todo: add your logic here and delete this line

    //todo: 需要返回数据则使用【 ResultData 】方法
	return l.Result(web.DBError, err)
}
