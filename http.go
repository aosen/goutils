/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: http中间件，采用接口思想，restfull，开发者只要继承基类，即可处理http请求
举个例子：https://github.com/aosen/novel
*/

package utils

import (
	"errors"
	"log"
	"net/http"
)

func PutError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type HandleInterface interface {
	Prepare(w http.ResponseWriter, r *http.Request, web *Web)
	Get(w http.ResponseWriter, r *http.Request, web *Web)
	Put(w http.ResponseWriter, r *http.Request, web *Web)
	Post(w http.ResponseWriter, r *http.Request, web *Web)
	Options(w http.ResponseWriter, r *http.Request, web *Web)
	Head(w http.ResponseWriter, r *http.Request, web *Web)
	Delete(w http.ResponseWriter, r *http.Request, web *Web)
	Connect(w http.ResponseWriter, r *http.Request, web *Web)
	Finish(w http.ResponseWriter, r *http.Request, web *Web)
}

/*全局控制对象*/
type Web struct {
	//配置信息
	Settings map[string]string
}

func NewWeb(setting map[string]string) *Web {
	return &Web{
		Settings: setting,
	}
}

func (self *Web) Go(handler HandleInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//为了保证程序不会异常退出，增加recover
		debug, ok := GetSetting(self.Settings, "DEBUG")
		if !ok {
			PutError(errors.New("not found setting for debug"))
		}
		if debug != "True" {
			defer func() {
				if x := recover(); x != nil {
					log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
				}
			}()
		}
		//无论什么方法 都预先调用prepare方法
		handler.Prepare(w, r, self)
		//相应http方法关联处理
		switch r.Method {
		case "GET":
			handler.Get(w, r, self)
		case "PUT":
			handler.Put(w, r, self)
		case "POST":
			handler.Post(w, r, self)
		case "OPTIONS":
			handler.Options(w, r, self)
		case "HEAD":
			handler.Head(w, r, self)
		case "DELETE":
			handler.Delete(w, r, self)
		case "CONNECT":
			handler.Connect(w, r, self)
		}
		//无论什么方法 结束后都调用finish方法
		handler.Finish(w, r, self)
	}
}

//所有http处理类都继承此类
type WebHandler struct {
}

func (self *WebHandler) Prepare(w http.ResponseWriter, r *http.Request, web *Web) {
}

func (self *WebHandler) Get(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Put(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Post(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Options(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Head(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Delete(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Connect(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Finish(w http.ResponseWriter, r *http.Request, web *Web) {
}

//获取配置文件信息
func GetSetting(settings map[string]string, key string) (string, bool) {
	value, ok := settings[key]
	return value, ok
}
