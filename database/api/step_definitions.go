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

func configStepDefinitionsRouter(router *httprouter.Router) {
	router.GET("/stepdefinitions", GetAllStepDefinitions)
	router.POST("/stepdefinitions", AddStepDefinitions)

	router.GET("/stepdefinitions/:argId", GetStepDefinitions)
	router.PUT("/stepdefinitions/:argId", UpdateStepDefinitions)
	router.DELETE("/stepdefinitions/:argId", DeleteStepDefinitions)
}

func configGinStepDefinitionsRouter(router gin.IRoutes) {
	router.GET("/stepdefinitions", ConverHttprouterToGin(GetAllStepDefinitions))
	router.POST("/stepdefinitions", ConverHttprouterToGin(AddStepDefinitions))
	router.GET("/stepdefinitions/:argId", ConverHttprouterToGin(GetStepDefinitions))
	router.PUT("/stepdefinitions/:argId", ConverHttprouterToGin(UpdateStepDefinitions))
	router.DELETE("/stepdefinitions/:argId", ConverHttprouterToGin(DeleteStepDefinitions))
}

// GetAllStepDefinitions is a function to get a slice of record(s) from step_definitions table in the laforge-2 database
// @Summary Get list of StepDefinitions
// @Tags StepDefinitions
// @Description GetAllStepDefinitions is a handler to get a slice of record(s) from step_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.StepDefinitions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /stepdefinitions [get]
// http "http://localhost:8080/stepdefinitions?page=0&pagesize=20"
func GetAllStepDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllStepDefinitions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetStepDefinitions is a function to get a single record from the step_definitions table in the laforge-2 database
// @Summary Get record from table StepDefinitions by  argId
// @Tags StepDefinitions
// @ID argId
// @Description GetStepDefinitions is a function to get a single record from the step_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.StepDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /stepdefinitions/{argId} [get]
// http "http://localhost:8080/stepdefinitions/hello world"
func GetStepDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetStepDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddStepDefinitions add to add a single record to step_definitions table in the laforge-2 database
// @Summary Add an record to step_definitions table
// @Description add to add a single record to step_definitions table in the laforge-2 database
// @Tags StepDefinitions
// @Accept  json
// @Produce  json
// @Param StepDefinitions body model.StepDefinitions true "Add StepDefinitions"
// @Success 200 {object} model.StepDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /stepdefinitions [post]
// echo '{"id": "`^EHvF\\KEUAoQbLoTP[JrXRRy","type": "]g_fx^QlwPBOhYKSAtZKfmLAQ"}' | http POST "http://localhost:8080/stepdefinitions"
func AddStepDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	stepdefinitions := &model.StepDefinitions{}

	if err := readJSON(r, stepdefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := stepdefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	stepdefinitions.Prepare()

	if err := stepdefinitions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	stepdefinitions, _, err = dao.AddStepDefinitions(r.Context(), stepdefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, stepdefinitions)
}

// UpdateStepDefinitions Update a single record from step_definitions table in the laforge-2 database
// @Summary Update an record in table step_definitions
// @Description Update a single record from step_definitions table in the laforge-2 database
// @Tags StepDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  StepDefinitions body model.StepDefinitions true "Update StepDefinitions record"
// @Success 200 {object} model.StepDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /stepdefinitions/{argId} [patch]
// echo '{"id": "`^EHvF\\KEUAoQbLoTP[JrXRRy","type": "]g_fx^QlwPBOhYKSAtZKfmLAQ"}' | http PUT "http://localhost:8080/stepdefinitions/hello world"
func UpdateStepDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	stepdefinitions := &model.StepDefinitions{}
	if err := readJSON(r, stepdefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := stepdefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	stepdefinitions.Prepare()

	if err := stepdefinitions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	stepdefinitions, _, err = dao.UpdateStepDefinitions(r.Context(),
		argId,
		stepdefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, stepdefinitions)
}

// DeleteStepDefinitions Delete a single record from step_definitions table in the laforge-2 database
// @Summary Delete a record from step_definitions
// @Description Delete a single record from step_definitions table in the laforge-2 database
// @Tags StepDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.StepDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /stepdefinitions/{argId} [delete]
// http DELETE "http://localhost:8080/stepdefinitions/hello world"
func DeleteStepDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteStepDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
