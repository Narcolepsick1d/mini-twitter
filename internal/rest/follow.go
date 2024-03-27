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

func (h *HandlerConfig) follow(c *gin.Context) {
	var (
		rq models.Follow
	)
	const (
		fn = "internal.rest.follow.follow"
	)
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error binding %v", fn, err), sl.Err(err))
	}
	token := c.Request.Header.Get("Authorization")
	rq.FollowerId, err = h.ParseToken(token[7:])
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error parse token %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: err.Error()})
		return
	}
	if rq.LeadId == rq.FollowerId {
		c.JSON(404, models.BaseResponse{Error: "You cant follow yourself"})
	}
	err = model.FollowUser(c.Request.Context(), h.Dependencies.DB, rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error parse token %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: err.Error()})
		return
	}
	c.JSON(200, "following")
}

func (h *HandlerConfig) unfollow(c *gin.Context) {
	var (
		rq models.Follow
	)
	const (
		fn = "internal.rest.follow.follow"
	)
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error binding %v", fn, err), sl.Err(err))
	}
	token := c.Request.Header.Get("Authorization")
	rq.FollowerId, err = h.ParseToken(token[7:])
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error parse token %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: err.Error()})
		return
	}
	if rq.LeadId == rq.FollowerId {
		c.JSON(404, models.BaseResponse{Error: "You cant unfollow yourself"})
	}
	err = model.UnFollowUser(c.Request.Context(), h.Dependencies.DB, goqu.Ex{"lead_id": rq.LeadId, "follower": rq.FollowerId})
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error db unfollow %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: err.Error()})
		return
	}
	c.JSON(200, "unfollowing")
}
