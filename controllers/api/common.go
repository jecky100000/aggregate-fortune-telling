package api

import (
	"gin/ay"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	Token string
	Appid int
	Json  *ay.Json
)

type CommonController struct {
}

func (con CommonController) A() {

}

func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

func GetToken(token string) string {
	uid := ay.AuthCode(token, "DECODE", "", 0)
	return uid
}

func DateFormat(y, m, d, h, i, s int) string {
	vm := strconv.Itoa(m)
	if len(vm) == 1 {
		vm = "0" + vm
	}

	vd := strconv.Itoa(d)
	if len(vd) == 1 {
		vd = "0" + vd
	}

	vh := strconv.Itoa(h)
	if len(vh) == 1 {
		vh = "0" + vh
	}

	vi := strconv.Itoa(i)
	if len(vi) == 1 {
		vi = "0" + vi
	}

	vs := strconv.Itoa(s)
	if len(vs) == 1 {
		vs = "0" + vs
	}

	return strconv.Itoa(y) + "-" + vm + "-" + vd + " " + vh + ":" + vi + ":" + vs
}
