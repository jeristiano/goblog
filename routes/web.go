package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeristiano/goblog/app/http/controllers"
)

// Initialize 初始化路由
func Initialize() {
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	// 静态页面
	r.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	r.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")
}
