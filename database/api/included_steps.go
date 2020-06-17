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

func configIncludedStepsRouter(router *httprouter.Router) {
	router.GET("/includedsteps", GetAllIncludedSteps)
	router.POST("/includedsteps", AddIncludedSteps)

	router.GET("/includedsteps/:argId", GetIncludedSteps)
	router.PUT("/includedsteps/:argId", UpdateIncludedSteps)
	router.DELETE("/includedsteps/:argId", DeleteIncludedSteps)
}

func configGinIncludedStepsRouter(router gin.IRoutes) {
	router.GET("/includedsteps", ConverHttprouterToGin(GetAllIncludedSteps))
	router.POST("/includedsteps", ConverHttprouterToGin(AddIncludedSteps))
	router.GET("/includedsteps/:argId", ConverHttprouterToGin(GetIncludedSteps))
	router.PUT("/includedsteps/:argId", ConverHttprouterToGin(UpdateIncludedSteps))
	router.DELETE("/includedsteps/:argId", ConverHttprouterToGin(DeleteIncludedSteps))
}

// GetAllIncludedSteps is a function to get a slice of record(s) from included_steps table in the laforge-2 database
// @Summary Get list of IncludedSteps
// @Tags IncludedSteps
// @Description GetAllIncludedSteps is a handler to get a slice of record(s) from included_steps table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IncludedSteps}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedsteps [get]
// http "http://localhost:8080/includedsteps?page=0&pagesize=20"
func GetAllIncludedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllIncludedSteps(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetIncludedSteps is a function to get a single record from the included_steps table in the laforge-2 database
// @Summary Get record from table IncludedSteps by  argId
// @Tags IncludedSteps
// @ID argId
// @Description GetIncludedSteps is a function to get a single record from the included_steps table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.IncludedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /includedsteps/{argId} [get]
// http "http://localhost:8080/includedsteps/hello world"
func GetIncludedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetIncludedSteps(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddIncludedSteps add to add a single record to included_steps table in the laforge-2 database
// @Summary Add an record to included_steps table
// @Description add to add a single record to included_steps table in the laforge-2 database
// @Tags IncludedSteps
// @Accept  json
// @Produce  json
// @Param IncludedSteps body model.IncludedSteps true "Add IncludedSteps"
// @Success 200 {object} model.IncludedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedsteps [post]
// echo '{"id": "ZbsJoApNi_dPMJKkxseydFDaO","included_hosts_id": "nltoF\\ouLdOksk^sqcXuT_jop","step_definitions_id": "kZKvNwLkgjVKFcDgebfRYPsJF","step_definition_type": "UWlKS[Zw`sH[wkIiSeVmlHjht","step_offset": 15}' | http POST "http://localhost:8080/includedsteps"
func AddIncludedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	includedsteps := &model.IncludedSteps{}

	if err := readJSON(r, includedsteps); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includedsteps.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includedsteps.Prepare()

	if err := includedsteps.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	includedsteps, _, err = dao.AddIncludedSteps(r.Context(), includedsteps)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includedsteps)
}

// UpdateIncludedSteps Update a single record from included_steps table in the laforge-2 database
// @Summary Update an record in table included_steps
// @Description Update a single record from included_steps table in the laforge-2 database
// @Tags IncludedSteps
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  IncludedSteps body model.IncludedSteps true "Update IncludedSteps record"
// @Success 200 {object} model.IncludedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedsteps/{argId} [patch]
// echo '{"id": "ZbsJoApNi_dPMJKkxseydFDaO","included_hosts_id": "nltoF\\ouLdOksk^sqcXuT_jop","step_definitions_id": "kZKvNwLkgjVKFcDgebfRYPsJF","step_definition_type": "UWlKS[Zw`sH[wkIiSeVmlHjht","step_offset": 15}' | http PUT "http://localhost:8080/includedsteps/hello world"
func UpdateIncludedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	includedsteps := &model.IncludedSteps{}
	if err := readJSON(r, includedsteps); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includedsteps.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includedsteps.Prepare()

	if err := includedsteps.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	includedsteps, _, err = dao.UpdateIncludedSteps(r.Context(),
		argId,
		includedsteps)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includedsteps)
}

// DeleteIncludedSteps Delete a single record from included_steps table in the laforge-2 database
// @Summary Delete a record from included_steps
// @Description Delete a single record from included_steps table in the laforge-2 database
// @Tags IncludedSteps
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.IncludedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /includedsteps/{argId} [delete]
// http DELETE "http://localhost:8080/includedsteps/hello world"
func DeleteIncludedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIncludedSteps(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
