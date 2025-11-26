package myrouter

import (
	"net/http"
)

type Router struct {
	trees []*Tree
}

func New() *Router {
	return &Router{
		// TODO: write smarter
		trees: []*Tree{
			{},
			{},
			{},
			{},
			{},
			{},
			{},
			{},
			{},
			{},
		},
	}
}

func (r *Router) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.insert("GET", path, handler)
}

func (r *Router) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.insert("POST", path, handler)
}

func (r *Router) DELETE(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.insert("DELETE", path, handler)
}

func (r *Router) insert(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	validatePath(path)

	switch method {
	case "GET":
		r.trees[0].add(path, handler)
	case "POST":
		r.trees[1].add(path, handler)
	}
}

// TODO: write more pattern, detail.
func getMethodIndexOf(method string) int {
	switch method {
	case "GET":
		return 0
	case "POST":
		return 1
	}
	return -1
}

// func (r *Router) get(method string, path string) *Node {
//
// }
//
// func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	path := req.URL
// 	method := req.Method
// 	methodIndex := getMethodIndexOf(method)
// }
