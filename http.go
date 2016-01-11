/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: http中间件，采用接口思想，restfull，开发者只要继承基类，即可处理http请求
*/

package utils

import (
	"fmt"
	"log"
	"net/http"
)

func PutError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type HandleInterface interface {
	Prepare(w http.ResponseWriter, r *http.Request, g *Global)
	Get(w http.ResponseWriter, r *http.Request, g *Global)
	Put(w http.ResponseWriter, r *http.Request, g *Global)
	Post(w http.ResponseWriter, r *http.Request, g *Global)
	Options(w http.ResponseWriter, r *http.Request, g *Global)
	Head(w http.ResponseWriter, r *http.Request, g *Global)
	Delete(w http.ResponseWriter, r *http.Request, g *Global)
	Connect(w http.ResponseWriter, r *http.Request, g *Global)
	Finish(w http.ResponseWriter, r *http.Request, g *Global)
}

/*全局控制对象*/
type Web struct {
	//settings 配置信息
	Settings map[string]interface{}
}

func NewWeb() *Web {
	return &Web{}
}

func (self *Web) Go(handler HandleInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//为了保证程序不会异常退出，增加recover
		debug, err := GetSetting(self.Settings, "DEBUG")
		PutError(err)
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
type BaseHandle struct {
}

func (self *BaseHandle) Prepare(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Get(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Put(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Post(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Options(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Head(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Delete(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Connect(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *BaseHandle) Finish(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}
