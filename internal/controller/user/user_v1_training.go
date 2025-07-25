package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/user/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/internal/thirdparty/codeforces"
)

func (c *ControllerV1) Training(ctx context.Context, req *v1.TrainingReq) (res *v1.TrainingRes, err error) {
	r := g.RequestFromCtx(ctx)
	md := dao.UserBase.Ctx(ctx)
	user := &entity.UserBase{}
	username := gconv.String(r.Get("username").Val())
	err = md.Where("username", username).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	err = gconv.Scan(user, &res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if user.CfId != "" {
		data := codeforces.GetUserInfo(ctx, user.CfId)
		res.CFRating = data.Rating
	}
	return res, nil
}
