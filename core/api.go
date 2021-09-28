package core

import (
	"com.lierda.wsn.vc/util"
	"encoding/json"
	"log"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func Login(request LoginRequest) *LoginResponse {
	res, code := util.Post("login", request)
	response := &LoginResponse{}
	if code == 200 {
		json.Unmarshal(res, response)
	} else {
		return nil
	}
	return response
}

type FirmwareTreeRequest struct {
}

type FirmwareTree struct {
	Id           int            `json:"id"`
	ParentId     int            `json:"parentId"`
	Type         int            `json:"type"`
	FirmwareType int            `json:"firmwareType"`
	Name         string         `json:"name"`
	Model        string         `json:"model"`
	Src          int            `json:"src"`
	Version      string         `json:"version"`
	CreateTime   time.Time      `json:"createTime"`
	CreateBy     int            `json:"createBy"`
	Children     []FirmwareTree `json:"children"`
}
type FirmwareTreeResponse struct {
	Data struct {
		Tree []FirmwareTree `json:"tree"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func GetFirmwareTree() FirmwareTreeResponse {
	res, code := util.Get("firmware/tree", nil)
	firmwareTreeResponse := FirmwareTreeResponse{}
	json.Unmarshal(res, &firmwareTreeResponse)
	log.Println(string(res), code)
	return firmwareTreeResponse
}
