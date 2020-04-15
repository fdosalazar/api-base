package apibasebundle

import (
	"api-base/app/core"
	"github.com/jinzhu/gorm"
	"net/http"
)

// ParametrosBundle handle kitties resources
type ApiBaseBundle struct {
	routes []core.Route
}

// NewParametrosBundle instance
func NewApiBaseBundle(db *gorm.DB) core.Bundle {
	apiBaseSQLMapper := NewApiBaseSQLMapper(db)
	apiBaseController := NewApiBaseController(apiBaseSQLMapper)

	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/alive",
			Handler: apiBaseController.Alive,
		},
		core.Route{
			Method:  http.MethodGet,
			Path:    "/env",
			Handler: apiBaseController.Env,
		},
	}

	return &ApiBaseBundle{
		routes: r,
	}
}

// GetRoutes implement interface core.Bundle
func (b *ApiBaseBundle) GetRoutes() []core.Route {
	return b.routes
}
