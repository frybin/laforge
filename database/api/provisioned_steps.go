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

func configProvisionedStepsRouter(router *httprouter.Router) {
	router.GET("/provisionedsteps", GetAllProvisionedSteps)
	router.POST("/provisionedsteps", AddProvisionedSteps)

	router.GET("/provisionedsteps/:argId", GetProvisionedSteps)
	router.PUT("/provisionedsteps/:argId", UpdateProvisionedSteps)
	router.DELETE("/provisionedsteps/:argId", DeleteProvisionedSteps)
}

func configGinProvisionedStepsRouter(router gin.IRoutes) {
	router.GET("/provisionedsteps", ConverHttprouterToGin(GetAllProvisionedSteps))
	router.POST("/provisionedsteps", ConverHttprouterToGin(AddProvisionedSteps))
	router.GET("/provisionedsteps/:argId", ConverHttprouterToGin(GetProvisionedSteps))
	router.PUT("/provisionedsteps/:argId", ConverHttprouterToGin(UpdateProvisionedSteps))
	router.DELETE("/provisionedsteps/:argId", ConverHttprouterToGin(DeleteProvisionedSteps))
}

// GetAllProvisionedSteps is a function to get a slice of record(s) from provisioned_steps table in the laforge-2 database
// @Summary Get list of ProvisionedSteps
// @Tags ProvisionedSteps
// @Description GetAllProvisionedSteps is a handler to get a slice of record(s) from provisioned_steps table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ProvisionedSteps}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedsteps [get]
// http "http://localhost:8080/provisionedsteps?page=0&pagesize=20"
func GetAllProvisionedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllProvisionedSteps(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetProvisionedSteps is a function to get a single record from the provisioned_steps table in the laforge-2 database
// @Summary Get record from table ProvisionedSteps by  argId
// @Tags ProvisionedSteps
// @ID argId
// @Description GetProvisionedSteps is a function to get a single record from the provisioned_steps table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.ProvisionedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /provisionedsteps/{argId} [get]
// http "http://localhost:8080/provisionedsteps/hello world"
func GetProvisionedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetProvisionedSteps(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddProvisionedSteps add to add a single record to provisioned_steps table in the laforge-2 database
// @Summary Add an record to provisioned_steps table
// @Description add to add a single record to provisioned_steps table in the laforge-2 database
// @Tags ProvisionedSteps
// @Accept  json
// @Produce  json
// @Param ProvisionedSteps body model.ProvisionedSteps true "Add ProvisionedSteps"
// @Success 200 {object} model.ProvisionedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedsteps [post]
// echo '{"id": "jRqwWyqNbaBreHPO]RZPpegSS","state": "gXlBvM\\sYRWLvvfGJsuQKGHjq","planned_checksum": "nkGQlvLRyaSxcFNT`UQIwimQu","current_checksum": "XCpeyScTbtFfHxBVdhGvKxfjE","previous_checksum": "CijiVY[_I\\iPrwIiqefgxwnkv","provisioned_hosts_id": "K`xgdKF^EFsNQR]\\SFChQDBsD"}' | http POST "http://localhost:8080/provisionedsteps"
func AddProvisionedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	provisionedsteps := &model.ProvisionedSteps{}

	if err := readJSON(r, provisionedsteps); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionedsteps.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionedsteps.Prepare()

	if err := provisionedsteps.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	provisionedsteps, _, err = dao.AddProvisionedSteps(r.Context(), provisionedsteps)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionedsteps)
}

// UpdateProvisionedSteps Update a single record from provisioned_steps table in the laforge-2 database
// @Summary Update an record in table provisioned_steps
// @Description Update a single record from provisioned_steps table in the laforge-2 database
// @Tags ProvisionedSteps
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  ProvisionedSteps body model.ProvisionedSteps true "Update ProvisionedSteps record"
// @Success 200 {object} model.ProvisionedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedsteps/{argId} [patch]
// echo '{"id": "jRqwWyqNbaBreHPO]RZPpegSS","state": "gXlBvM\\sYRWLvvfGJsuQKGHjq","planned_checksum": "nkGQlvLRyaSxcFNT`UQIwimQu","current_checksum": "XCpeyScTbtFfHxBVdhGvKxfjE","previous_checksum": "CijiVY[_I\\iPrwIiqefgxwnkv","provisioned_hosts_id": "K`xgdKF^EFsNQR]\\SFChQDBsD"}' | http PUT "http://localhost:8080/provisionedsteps/hello world"
func UpdateProvisionedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	provisionedsteps := &model.ProvisionedSteps{}
	if err := readJSON(r, provisionedsteps); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionedsteps.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionedsteps.Prepare()

	if err := provisionedsteps.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	provisionedsteps, _, err = dao.UpdateProvisionedSteps(r.Context(),
		argId,
		provisionedsteps)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionedsteps)
}

// DeleteProvisionedSteps Delete a single record from provisioned_steps table in the laforge-2 database
// @Summary Delete a record from provisioned_steps
// @Description Delete a single record from provisioned_steps table in the laforge-2 database
// @Tags ProvisionedSteps
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.ProvisionedSteps
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /provisionedsteps/{argId} [delete]
// http DELETE "http://localhost:8080/provisionedsteps/hello world"
func DeleteProvisionedSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProvisionedSteps(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
