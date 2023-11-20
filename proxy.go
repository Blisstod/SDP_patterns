package main

import "fmt"

type UserInfo interface {
	GetUserDetails() string
}

type RealUserInfo struct {
	userDetails string
}

func (r *RealUserInfo) GetUserDetails() string {
	return r.userDetails
}

func NewRealUserInfo(details string) *RealUserInfo {
	return &RealUserInfo{userDetails: details}
}

type UserInfoProxy struct {
	realUserInfo *RealUserInfo
	userToken    string
}

func (p *UserInfoProxy) GetUserDetails() string {
	if p.userToken == "validToken" {
		return p.realUserInfo.GetUserDetails()
	}
	return "Access Denied: Invalid Token"
}

func NewUserInfoProxy(realUserInfo *RealUserInfo, token string) *UserInfoProxy {
	return &UserInfoProxy{realUserInfo: realUserInfo, userToken: token}
}

func main() {
	realUserInfo := NewRealUserInfo("Sensitive User Details")

	// Proxy with valid token
	proxyWithAccess := NewUserInfoProxy(realUserInfo, "validToken")
	fmt.Println(proxyWithAccess.GetUserDetails())

	// Proxy with invalid token
	proxyWithoutAccess := NewUserInfoProxy(realUserInfo, "invalidToken")
	fmt.Println(proxyWithoutAccess.GetUserDetails())
}
