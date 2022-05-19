package routes

import (
	"net/http"

	"goblog/app/http/controllers"

	"github.com/gorilla/mux"
)

// Initialize 初始化路由
func Initialize() {
	Router := mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
}
