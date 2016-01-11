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
	Prepare(w http.ResponseWriter, r *http.Request, g *G)
	Get(w http.ResponseWriter, r *http.Request, g *G)
	Put(w http.ResponseWriter, r *http.Request, g *G)
	Post(w http.ResponseWriter, r *http.Request, g *G)
	Options(w http.ResponseWriter, r *http.Request, g *G)
	Head(w http.ResponseWriter, r *http.Request, g *G)
	Delete(w http.ResponseWriter, r *http.Request, g *G)
	Connect(w http.ResponseWriter, r *http.Request, g *G)
	Finish(w http.ResponseWriter, r *http.Request, g *G)
}

/*全局控制对象*/
type Global struct {
	//settings 配置信息
	Settings map[string]interface{}
}

func (self *Global) Go(handler HandleInterface) http.HandlerFunc {
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
type Web struct {
}

func (self *Web) Prepare(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Get(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Put(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Post(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Options(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Head(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Delete(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Connect(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}

func (self *Web) Finish(w http.ResponseWriter, r *http.Request, g *G) {
	fmt.Fprintln(w, "sorry 404 not found")
}
