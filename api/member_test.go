package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router = setupRouter()

func setupRouter() *gin.Engine {
	r := gin.Default()
	MemberRouting(r)
	return r
}

func TestCreateMember(t *testing.T) {

	form := url.Values{}
	form.Add("name", "abc")
	form.Add("email", "abc")
	form.Add("secret", "abc")

	req, _ := http.NewRequest("PUT", "/create/member", nil)
	req.PostForm = form
	res := httptest.NewRecorder()

	t.Log(req.PostForm)
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	t.Log(res.Body)
}

func TestReadMember(t *testing.T) {

	req, _ := http.NewRequest("GET", "/read/member/abc", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	t.Log(res.Body)
}