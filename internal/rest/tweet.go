package rest

import (
	"Narcolepsick1d/mini-twitter/internal/model"
	"Narcolepsick1d/mini-twitter/internal/models"
	"Narcolepsick1d/mini-twitter/pkg/sl"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (h *HandlerConfig) tweet(c *gin.Context) {
	var (
		rq models.TweetReq
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
	err = model.CreateTweet(c.Request.Context(), h.Dependencies.DB, rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error db %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
		return
	}
	c.JSON(200, gin.H{})
	return
}
