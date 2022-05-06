package efficient

var Routers = &routers{gRouters: map[groupRelativePath]*gRouter{}}

type groupRelativePath = string

type routers struct {
	gRouters map[groupRelativePath]*gRouter
}

type gRouter struct {
	routers      *routers
	relativePath string
	handlers     GroupInterface
	gRouters     []router
}

type router struct {
	httpMethod   string
	relativePath string
	handler      ControllerInterface
}

func (r *routers) Group(relativePath string, group GroupInterface) *gRouter {
	r.gRouters[relativePath] = &gRouter{
		routers:      nil,
		relativePath: relativePath,
		handlers:     group,
		gRouters:     nil,
	}
	return &gRouter{
		routers:      r,
		relativePath: relativePath,
		handlers:     group,
	}
}

func (g *gRouter) Add(relativePath string, controller ControllerInterface, httpMethod ...string) *gRouter {
	g.routers.add(g.relativePath, httpMethod, relativePath, controller)
	return g
}

func (r *routers) add(groupRelativePath string, httpMethods []string, relativePath string, controller ControllerInterface) {
	if _, ok := r.gRouters[groupRelativePath]; !ok {
		r.gRouters[groupRelativePath] = &gRouter{
			routers:      r,
			relativePath: groupRelativePath,
			handlers:     nil,
			gRouters:     nil,
		}
	}
	for _, httpMethod := range httpMethods {
		r.gRouters[groupRelativePath].gRouters = append(r.gRouters[groupRelativePath].gRouters, router{
			httpMethod:   httpMethod,
			relativePath: relativePath,
			handler:      controller,
		})
	}
}

func (r *routers) Add(relativePath string, controller ControllerInterface, httpMethod ...string) *routers {
	r.add("/", httpMethod, relativePath, controller)
	return r
}

func (r *routers) handle(groupRelativePath string, httpMethod string, relativePath string, controller ControllerInterface) {
	gRouters := r.gRouters[groupRelativePath]
	gRouters.gRouters = append(r.gRouters[groupRelativePath].gRouters, router{
		httpMethod:   httpMethod,
		relativePath: relativePath,
		handler:      controller,
	})
}

func registerRouters(e Engine) {
	for gname, gRouter := range Routers.gRouters {
		if gRouter.handlers != nil {
			g := e.Group(gname, func(cxt Context) {
				gRouter.handlers.Execute(cxt)
			})
			for _, router := range gRouter.gRouters {
				method := router.httpMethod
				g.Handle(method, router.relativePath, func(cxt Context) {
					router.handler.Prepare(cxt)
					router.handler.Execute(cxt, router.handler)
					router.handler.Finish(cxt)
				})
			}
			continue
		}

		for _, router := range gRouter.gRouters {
			method := router.httpMethod
			e.Handle(method, router.relativePath, func(cxt Context) {
				router.handler.Prepare(cxt)
				router.handler.Execute(cxt, router.handler)
				router.handler.Finish(cxt)
			})
		}
	}
}