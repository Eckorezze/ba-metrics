package routes

func (router *Router) V1Routes() {
	v1 := router.Router.Group("/v1")

	route := Route{Group: v1, Pg: router.Pg}

	route.HelloRoute()
}
