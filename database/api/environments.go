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

func configEnvironmentsRouter(router *httprouter.Router) {
	router.GET("/environments", GetAllEnvironments)
	router.POST("/environments", AddEnvironments)

	router.GET("/environments/:argId", GetEnvironments)
	router.PUT("/environments/:argId", UpdateEnvironments)
	router.DELETE("/environments/:argId", DeleteEnvironments)
}

func configGinEnvironmentsRouter(router gin.IRoutes) {
	router.GET("/environments", ConverHttprouterToGin(GetAllEnvironments))
	router.POST("/environments", ConverHttprouterToGin(AddEnvironments))
	router.GET("/environments/:argId", ConverHttprouterToGin(GetEnvironments))
	router.PUT("/environments/:argId", ConverHttprouterToGin(UpdateEnvironments))
	router.DELETE("/environments/:argId", ConverHttprouterToGin(DeleteEnvironments))
}

// GetAllEnvironments is a function to get a slice of record(s) from environments table in the laforge-2 database
// @Summary Get list of Environments
// @Tags Environments
// @Description GetAllEnvironments is a handler to get a slice of record(s) from environments table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Environments}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /environments [get]
// http "http://localhost:8080/environments?page=0&pagesize=20"
func GetAllEnvironments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllEnvironments(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetEnvironments is a function to get a single record from the environments table in the laforge-2 database
// @Summary Get record from table Environments by  argId
// @Tags Environments
// @ID argId
// @Description GetEnvironments is a function to get a single record from the environments table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.Environments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /environments/{argId} [get]
// http "http://localhost:8080/environments/hello world"
func GetEnvironments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetEnvironments(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddEnvironments add to add a single record to environments table in the laforge-2 database
// @Summary Add an record to environments table
// @Description add to add a single record to environments table in the laforge-2 database
// @Tags Environments
// @Accept  json
// @Produce  json
// @Param Environments body model.Environments true "Add Environments"
// @Success 200 {object} model.Environments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /environments [post]
// echo '{"name": "TB`Ke[ZQ[eGdyRlC^Aglv[Jow","attrs": "jdJ^ZCxGXiIY_q`lWiKwtKjyh","planned_checksum": "xuPNYquDBWyXU`NwCDdhMIBlC","current_checksum": "MgYBtJpIivnh[OHnCCNjTjkjI","previous_checksum": "tTI`xHge[piJOhKKMddMcaOj_","id": "OmYjkkI_OHBk\\PFRmNPHQvRsN","competitions_id": "[WmhsNWYQ]]UlWhcCGCTnrcHD","roles_id": "Cx]r`YxMY^uHmlieNhKvV^OPA"}' | http POST "http://localhost:8080/environments"
func AddEnvironments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	environments := &model.Environments{}

	if err := readJSON(r, environments); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := environments.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	environments.Prepare()

	if err := environments.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	environments, _, err = dao.AddEnvironments(r.Context(), environments)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, environments)
}

// UpdateEnvironments Update a single record from environments table in the laforge-2 database
// @Summary Update an record in table environments
// @Description Update a single record from environments table in the laforge-2 database
// @Tags Environments
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  Environments body model.Environments true "Update Environments record"
// @Success 200 {object} model.Environments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /environments/{argId} [patch]
// echo '{"name": "TB`Ke[ZQ[eGdyRlC^Aglv[Jow","attrs": "jdJ^ZCxGXiIY_q`lWiKwtKjyh","planned_checksum": "xuPNYquDBWyXU`NwCDdhMIBlC","current_checksum": "MgYBtJpIivnh[OHnCCNjTjkjI","previous_checksum": "tTI`xHge[piJOhKKMddMcaOj_","id": "OmYjkkI_OHBk\\PFRmNPHQvRsN","competitions_id": "[WmhsNWYQ]]UlWhcCGCTnrcHD","roles_id": "Cx]r`YxMY^uHmlieNhKvV^OPA"}' | http PUT "http://localhost:8080/environments/hello world"
func UpdateEnvironments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	environments := &model.Environments{}
	if err := readJSON(r, environments); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := environments.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	environments.Prepare()

	if err := environments.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	environments, _, err = dao.UpdateEnvironments(r.Context(),
		argId,
		environments)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, environments)
}

// DeleteEnvironments Delete a single record from environments table in the laforge-2 database
// @Summary Delete a record from environments
// @Description Delete a single record from environments table in the laforge-2 database
// @Tags Environments
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.Environments
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /environments/{argId} [delete]
// http DELETE "http://localhost:8080/environments/hello world"
func DeleteEnvironments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteEnvironments(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
