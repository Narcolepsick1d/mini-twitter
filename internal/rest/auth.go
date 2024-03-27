package rest

import (
	"Narcolepsick1d/mini-twitter/internal/model"
	"Narcolepsick1d/mini-twitter/internal/models"
	"Narcolepsick1d/mini-twitter/pkg/sl"
	"Narcolepsick1d/mini-twitter/pkg/validator"
	"context"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"math/rand"
	"time"
)

// signUp is the handler for user sign up.
// @Summary User sign up
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param query_params query models.UserSignUp true "Query parameters"
// @Success 200 {string} string "OK"
// @Failure 400 {object} models.BaseResponse "Wrong data"
// @Router /signup [post]
func (h *HandlerConfig) signUp(c *gin.Context) {
	var (
		rq models.UserSignUp
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
	if !validator.ValidateUserSignUp(rq) {
		c.JSON(404, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
		return
	}
	rq.Password, err = h.Dependencies.Hash.Hash(rq.Password)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error Hash %v", fn, err), sl.Err(err))
	}

	err = model.SignUp(c.Request.Context(), h.Dependencies.DB, rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error Sign up db %v", fn, err), sl.Err(err))
	}
	c.JSON(200, gin.H{})
}

func (h *HandlerConfig) signIn(c *gin.Context) {
	var (
		rq models.User
	)
	const (
		fn = "internal.rest.auth.signUp"
	)
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error binding %v", fn, err), sl.Err(err))
	}
	if !validator.ValidateUserSignIn(rq) {
		c.JSON(404, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
		return
	}
	rq.Password, err = h.Dependencies.Hash.Hash(rq.Password)
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error Hash %v", fn, err), sl.Err(err))
	}

	user, err := model.GetCredential(c.Request.Context(), h.Dependencies.DB, &goqu.Ex{"password": rq.Password, "email": rq.Email})
	if err != nil {
		slog.Error(fmt.Sprintf("%s Error Sign in db %v", fn, err), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: "db error", ErrorCode: 404})
		return
	}
	if len(user) == 0 {
		c.JSON(404, models.BaseResponse{Error: "non sing up", ErrorCode: 404})
		return
	}
	accessToken, refreshToken, err := h.GenerateTokens(user[0].Id)

	if err != nil {
		slog.Error(fmt.Sprintf("%s Error token in db %v", fn, err), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: err.Error()})
		return
	}

	if err != nil {
		slog.Error(fmt.Sprintf("%s Error unmarshal  %v", fn, err), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: err.Error()})
		return
	}

	c.Writer.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.Writer.Header().Add("Content-Type", "application/json")
	c.JSON(200, accessToken)
}

func (h *HandlerConfig) GenerateTokens(userId string) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	})

	accessToken, err := t.SignedString([]byte(h.Dependencies.Secret))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	if err := model.CreateToken(context.TODO(), h.Dependencies.DB, models.RefreshSession{
		UserID:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
func newRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (h *HandlerConfig) ParseToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(h.Dependencies.Secret), nil
	})
	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid subject")
	}

	return subject, nil
}
