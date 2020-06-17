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

func configScriptDefinitionsRouter(router *httprouter.Router) {
	router.GET("/scriptdefinitions", GetAllScriptDefinitions)
	router.POST("/scriptdefinitions", AddScriptDefinitions)

	router.GET("/scriptdefinitions/:argId", GetScriptDefinitions)
	router.PUT("/scriptdefinitions/:argId", UpdateScriptDefinitions)
	router.DELETE("/scriptdefinitions/:argId", DeleteScriptDefinitions)
}

func configGinScriptDefinitionsRouter(router gin.IRoutes) {
	router.GET("/scriptdefinitions", ConverHttprouterToGin(GetAllScriptDefinitions))
	router.POST("/scriptdefinitions", ConverHttprouterToGin(AddScriptDefinitions))
	router.GET("/scriptdefinitions/:argId", ConverHttprouterToGin(GetScriptDefinitions))
	router.PUT("/scriptdefinitions/:argId", ConverHttprouterToGin(UpdateScriptDefinitions))
	router.DELETE("/scriptdefinitions/:argId", ConverHttprouterToGin(DeleteScriptDefinitions))
}

// GetAllScriptDefinitions is a function to get a slice of record(s) from script_definitions table in the laforge-2 database
// @Summary Get list of ScriptDefinitions
// @Tags ScriptDefinitions
// @Description GetAllScriptDefinitions is a handler to get a slice of record(s) from script_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ScriptDefinitions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scriptdefinitions [get]
// http "http://localhost:8080/scriptdefinitions?page=0&pagesize=20"
func GetAllScriptDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllScriptDefinitions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetScriptDefinitions is a function to get a single record from the script_definitions table in the laforge-2 database
// @Summary Get record from table ScriptDefinitions by  argId
// @Tags ScriptDefinitions
// @ID argId
// @Description GetScriptDefinitions is a function to get a single record from the script_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.ScriptDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /scriptdefinitions/{argId} [get]
// http "http://localhost:8080/scriptdefinitions/hello world"
func GetScriptDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetScriptDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddScriptDefinitions add to add a single record to script_definitions table in the laforge-2 database
// @Summary Add an record to script_definitions table
// @Description add to add a single record to script_definitions table in the laforge-2 database
// @Tags ScriptDefinitions
// @Accept  json
// @Produce  json
// @Param ScriptDefinitions body model.ScriptDefinitions true "Add ScriptDefinitions"
// @Success 200 {object} model.ScriptDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scriptdefinitions [post]
// echo '{"id": "lq]MlxhrOUqrqVgvWHCVgqkZA","type": "xOiZUrS]gtSH]SWZcPxWKKwc_","source_files": "oFVdYHEPigrEyUgtYedadQryI","runtime": "oRicBq\\x^eLriQA]l^Flg\\ajT"}' | http POST "http://localhost:8080/scriptdefinitions"
func AddScriptDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	scriptdefinitions := &model.ScriptDefinitions{}

	if err := readJSON(r, scriptdefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := scriptdefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	scriptdefinitions.Prepare()

	if err := scriptdefinitions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	scriptdefinitions, _, err = dao.AddScriptDefinitions(r.Context(), scriptdefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, scriptdefinitions)
}

// UpdateScriptDefinitions Update a single record from script_definitions table in the laforge-2 database
// @Summary Update an record in table script_definitions
// @Description Update a single record from script_definitions table in the laforge-2 database
// @Tags ScriptDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  ScriptDefinitions body model.ScriptDefinitions true "Update ScriptDefinitions record"
// @Success 200 {object} model.ScriptDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scriptdefinitions/{argId} [patch]
// echo '{"id": "lq]MlxhrOUqrqVgvWHCVgqkZA","type": "xOiZUrS]gtSH]SWZcPxWKKwc_","source_files": "oFVdYHEPigrEyUgtYedadQryI","runtime": "oRicBq\\x^eLriQA]l^Flg\\ajT"}' | http PUT "http://localhost:8080/scriptdefinitions/hello world"
func UpdateScriptDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	scriptdefinitions := &model.ScriptDefinitions{}
	if err := readJSON(r, scriptdefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := scriptdefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	scriptdefinitions.Prepare()

	if err := scriptdefinitions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	scriptdefinitions, _, err = dao.UpdateScriptDefinitions(r.Context(),
		argId,
		scriptdefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, scriptdefinitions)
}

// DeleteScriptDefinitions Delete a single record from script_definitions table in the laforge-2 database
// @Summary Delete a record from script_definitions
// @Description Delete a single record from script_definitions table in the laforge-2 database
// @Tags ScriptDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.ScriptDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /scriptdefinitions/{argId} [delete]
// http DELETE "http://localhost:8080/scriptdefinitions/hello world"
func DeleteScriptDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteScriptDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
