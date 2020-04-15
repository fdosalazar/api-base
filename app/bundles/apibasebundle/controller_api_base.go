package apibasebundle

import (
	"api-base/app/core"
	"net/http"
	//"strconv"
)

// KittiesController struct
type ApiBaseController struct {
	core.Controller
	abm ApiBaseMapper
}

// NewKittiesController instance
func NewApiBaseController(abm ApiBaseMapper) *ApiBaseController {
	return &ApiBaseController{
		Controller: core.Controller{},
		abm:        abm,
	}
}

func (c *ApiBaseController) Alive(w http.ResponseWriter, r *http.Request) {
	response := c.abm.Alive()
	c.SendJSON(w, &response, http.StatusOK)
}

func (c *ApiBaseController) Env(w http.ResponseWriter, r *http.Request) {
	response := c.abm.Env()
	c.SendJSON(w, &response, http.StatusOK)
}
