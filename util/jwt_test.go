package util_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/lanwupark/blog-api/data"
	"github.com/lanwupark/blog-api/util"
	"github.com/stretchr/testify/assert"
)

func TestJWTSuccess(t *testing.T) {
	assert := assert.New(t)
	user := &data.TokenClaimsSubject{
		UserID:    123456,
		UserLogin: "eanson023",
		IsAdmin:   true,
	}
	token, err := util.CreateToken(user)
	t.Log(token)
	assert.NoError(err)
	parseUser, err := util.ParseToken(token)
	assert.NoError(err)
	assert.NotNil(parseUser)
	t.Log(parseUser)
	assert.Equal("eanson023", parseUser.UserLogin)
}

// TestJWTFailed 测试错误JWT解析
func TestJWTFailed(t *testing.T) {
	assert := assert.New(t)
	// 错误的token
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.OjE2MDgwMzc1NTQsImlzcyI6ImVhbnNvbiIsIm5iZiI6MTYwODAzNzQ5NCwic3ViIjoie1wiVXNlcklkXCI6MTIzNDU2LFwiTG9naW5cIjpcImVhbnNvbjAyM1wiLFwiSXNBZG1pblwiOnRydWV9In0.jUtGQMykZKlHChsYzghst_l-ynWAo0AjP-XYkvwNf5E"
	user, err := util.ParseToken(token)
	assert.Error(err)
	assert.Nil(user)
}

func TestExpired(t *testing.T) {
	assert := assert.New(t)
	// 过期的token
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDgwMzc1NTQsImlzcyI6ImVhbnNvbiIsIm5iZiI6MTYwODAzNzQ5NCwic3ViIjoie1wiVXNlcklkXCI6MTIzNDU2LFwiTG9naW5cIjpcImVhbnNvbjAyM1wiLFwiSXNBZG1pblwiOnRydWV9In0.jUtGQMykZKlHChsYzghst_l-ynWAo0AjP-XYkvwNf5E"
	user, err := util.ParseToken(token)
	assert.EqualError(err, "Token is expired")
	assert.Nil(user)
}

func TestUUID(t *testing.T) {
	t.Log(uuid.New().String())
}
