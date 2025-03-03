package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			ref := string(c.GetHeader("Referer"))
			next := "/sign-in"
			if ref != "" && utils.ValidateNext(ref) {
				next = fmt.Sprintf("/sign-in?next=%s", ref)
			}
			c.Redirect(http.StatusFound, []byte(next)) // 使用 http.StatusFound (302)
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}
