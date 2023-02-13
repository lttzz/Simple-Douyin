// Code generated by hertz generator.

package api

import (
	"context"
	"encoding/json"
	"log"

	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/cmd/api/biz/mw"
	"Simple-Douyin/cmd/api/rpc"
	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// resp := new(api.FavoriteActionResponse)
	v, _ := c.Get(constants.IdentityKey)
	resp, err := rpc.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
		UserId:     v.(*api.User).ID,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	uid := int64(0)
	if token, err := mw.JwtMiddleware.ParseTokenString(req.Token); err == nil {
		claims := jwt.ExtractClaimsFromToken(token)
		userid, _ := claims[constants.IdentityKey].(json.Number).Int64()
		uid = userid
	}
	log.Println("[ypx debug] api favorite userid", uid)
	// v, _ := c.Get(constants.IdentityKey)
	resp, err := rpc.FavoriteList(context.Background(), &favorite.FavoriteListRequest{
		UserId: req.UserID,
		// MUserId: v.(*api.User).ID,
		MUserId: uid,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}
