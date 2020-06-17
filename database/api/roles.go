package api

import (
	"net/http"

	"github.com/frybin/laforge/database/dao"
	"github.com/frybin/laforge/database/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configRolesRouter(router *httprouter.Router) {
	router.GET("/roles", GetAllRoles)
	router.POST("/roles", AddRoles)

	router.GET("/roles/:argId", GetRoles)
	router.PUT("/roles/:argId", UpdateRoles)
	router.DELETE("/roles/:argId", DeleteRoles)
}

func configGinRolesRouter(router gin.IRoutes) {
	router.GET("/roles", ConverHttprouterToGin(GetAllRoles))
	router.POST("/roles", ConverHttprouterToGin(AddRoles))
	router.GET("/roles/:argId", ConverHttprouterToGin(GetRoles))
	router.PUT("/roles/:argId", ConverHttprouterToGin(UpdateRoles))
	router.DELETE("/roles/:argId", ConverHttprouterToGin(DeleteRoles))
}

// GetAllRoles is a function to get a slice of record(s) from roles table in the laforge-2 database
// @Summary Get list of Roles
// @Tags Roles
// @Description GetAllRoles is a handler to get a slice of record(s) from roles table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Roles}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /roles [get]
// http "http://localhost:8080/roles?page=0&pagesize=20"
func GetAllRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	records, totalRows, err := dao.GetAllRoles(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetRoles is a function to get a single record from the roles table in the laforge-2 database
// @Summary Get record from table Roles by  argId
// @Tags Roles
// @ID argId
// @Description GetRoles is a function to get a single record from the roles table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.Roles
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /roles/{argId} [get]
// http "http://localhost:8080/roles/hello world"
func GetRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetRoles(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddRoles add to add a single record to roles table in the laforge-2 database
// @Summary Add an record to roles table
// @Description add to add a single record to roles table in the laforge-2 database
// @Tags Roles
// @Accept  json
// @Produce  json
// @Param Roles body model.Roles true "Add Roles"
// @Success 200 {object} model.Roles
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /roles [post]
// echo '{"name": "BhGSMFiupcplZYpEGkhGGXWQB","id": "jcjAJtLSAeDQO`bsfeT\\eh[vO"}' | http POST "http://localhost:8080/roles"
func AddRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	roles := &model.Roles{}

	if err := readJSON(r, roles); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := roles.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	roles.Prepare()

	if err := roles.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	roles, _, err = dao.AddRoles(r.Context(), roles)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, roles)
}

// UpdateRoles Update a single record from roles table in the laforge-2 database
// @Summary Update an record in table roles
// @Description Update a single record from roles table in the laforge-2 database
// @Tags Roles
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  Roles body model.Roles true "Update Roles record"
// @Success 200 {object} model.Roles
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /roles/{argId} [patch]
// echo '{"name": "BhGSMFiupcplZYpEGkhGGXWQB","id": "jcjAJtLSAeDQO`bsfeT\\eh[vO"}' | http PUT "http://localhost:8080/roles/hello world"
func UpdateRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	roles := &model.Roles{}
	if err := readJSON(r, roles); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := roles.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	roles.Prepare()

	if err := roles.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	roles, _, err = dao.UpdateRoles(r.Context(),
		argId,
		roles)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, roles)
}

// DeleteRoles Delete a single record from roles table in the laforge-2 database
// @Summary Delete a record from roles
// @Description Delete a single record from roles table in the laforge-2 database
// @Tags Roles
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.Roles
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /roles/{argId} [delete]
// http DELETE "http://localhost:8080/roles/hello world"
func DeleteRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRoles(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
