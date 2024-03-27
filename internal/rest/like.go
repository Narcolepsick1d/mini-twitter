package rest

import (
	"Narcolepsick1d/mini-twitter/internal/model"
	"Narcolepsick1d/mini-twitter/internal/models"
	"Narcolepsick1d/mini-twitter/pkg/sl"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (h *HandlerConfig) like(c *gin.Context) {
	var (
		rq models.LikeReq
	)
	const (
		fn = "internal.rest.auth.signUp"
	)
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error binding %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
		return
	}
	token := c.Request.Header.Get("Authorization")
	rq.UserId, err = h.ParseToken(token[7:])
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error parse token %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: err.Error()})
		return
	}
	err = model.AddLike(c.Request.Context(), h.Dependencies.DB, rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error db %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
		return
	}
	c.JSON(200, gin.H{})
	return
}

func (h *HandlerConfig) unlike(c *gin.Context) {
	var (
		rq models.LikeReq
	)
	const (
		fn = "internal.rest.auth.signUp"
	)
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error binding %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
		return
	}
	token := c.Request.Header.Get("Authorization")
	rq.UserId, err = h.ParseToken(token[7:])
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error parse token %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: err.Error()})
		return
	}
	err = model.RemoveLike(c.Request.Context(), h.Dependencies.DB, goqu.Ex{"user_id": rq.UserId, "tweet_id": rq.TweetId})
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error db %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
		return
	}
	c.JSON(200, gin.H{})
	return
}
