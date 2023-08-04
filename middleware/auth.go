package middleware

import (
	"courseScheduling/auth"
	"github.com/beego/beego/v2/server/web/context"
)

func Auth(ctx *context.Context) {
	tokenString := ctx.GetCookie("token")
	info, err := auth.ParseToken(tokenString)
	if err != nil {
		ctx.Abort(403, "auth fail")
		return
	}
	ctx.Input.SetData("uid", info.UserID)
}
