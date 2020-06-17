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

func configHostDefinitionsRouter(router *httprouter.Router) {
	router.GET("/hostdefinitions", GetAllHostDefinitions)
	router.POST("/hostdefinitions", AddHostDefinitions)

	router.GET("/hostdefinitions/:argId", GetHostDefinitions)
	router.PUT("/hostdefinitions/:argId", UpdateHostDefinitions)
	router.DELETE("/hostdefinitions/:argId", DeleteHostDefinitions)
}

func configGinHostDefinitionsRouter(router gin.IRoutes) {
	router.GET("/hostdefinitions", ConverHttprouterToGin(GetAllHostDefinitions))
	router.POST("/hostdefinitions", ConverHttprouterToGin(AddHostDefinitions))
	router.GET("/hostdefinitions/:argId", ConverHttprouterToGin(GetHostDefinitions))
	router.PUT("/hostdefinitions/:argId", ConverHttprouterToGin(UpdateHostDefinitions))
	router.DELETE("/hostdefinitions/:argId", ConverHttprouterToGin(DeleteHostDefinitions))
}

// GetAllHostDefinitions is a function to get a slice of record(s) from host_definitions table in the laforge-2 database
// @Summary Get list of HostDefinitions
// @Tags HostDefinitions
// @Description GetAllHostDefinitions is a handler to get a slice of record(s) from host_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.HostDefinitions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /hostdefinitions [get]
// http "http://localhost:8080/hostdefinitions?page=0&pagesize=20"
func GetAllHostDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllHostDefinitions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetHostDefinitions is a function to get a single record from the host_definitions table in the laforge-2 database
// @Summary Get record from table HostDefinitions by  argId
// @Tags HostDefinitions
// @ID argId
// @Description GetHostDefinitions is a function to get a single record from the host_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.HostDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /hostdefinitions/{argId} [get]
// http "http://localhost:8080/hostdefinitions/hello world"
func GetHostDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetHostDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddHostDefinitions add to add a single record to host_definitions table in the laforge-2 database
// @Summary Add an record to host_definitions table
// @Description add to add a single record to host_definitions table in the laforge-2 database
// @Tags HostDefinitions
// @Accept  json
// @Produce  json
// @Param HostDefinitions body model.HostDefinitions true "Add HostDefinitions"
// @Success 200 {object} model.HostDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /hostdefinitions [post]
// echo '{"hostname": "]aEvCmCjgif`bRAZ[sttmhUUi","os": "jqXLaqxQeNwIkUwjuIsrASgKZ","id": "qVmkWoUWWIfepXgJEf[nlpHPb"}' | http POST "http://localhost:8080/hostdefinitions"
func AddHostDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	hostdefinitions := &model.HostDefinitions{}

	if err := readJSON(r, hostdefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := hostdefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	hostdefinitions.Prepare()

	if err := hostdefinitions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	hostdefinitions, _, err = dao.AddHostDefinitions(r.Context(), hostdefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, hostdefinitions)
}

// UpdateHostDefinitions Update a single record from host_definitions table in the laforge-2 database
// @Summary Update an record in table host_definitions
// @Description Update a single record from host_definitions table in the laforge-2 database
// @Tags HostDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  HostDefinitions body model.HostDefinitions true "Update HostDefinitions record"
// @Success 200 {object} model.HostDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /hostdefinitions/{argId} [patch]
// echo '{"hostname": "]aEvCmCjgif`bRAZ[sttmhUUi","os": "jqXLaqxQeNwIkUwjuIsrASgKZ","id": "qVmkWoUWWIfepXgJEf[nlpHPb"}' | http PUT "http://localhost:8080/hostdefinitions/hello world"
func UpdateHostDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	hostdefinitions := &model.HostDefinitions{}
	if err := readJSON(r, hostdefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := hostdefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	hostdefinitions.Prepare()

	if err := hostdefinitions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	hostdefinitions, _, err = dao.UpdateHostDefinitions(r.Context(),
		argId,
		hostdefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, hostdefinitions)
}

// DeleteHostDefinitions Delete a single record from host_definitions table in the laforge-2 database
// @Summary Delete a record from host_definitions
// @Description Delete a single record from host_definitions table in the laforge-2 database
// @Tags HostDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.HostDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /hostdefinitions/{argId} [delete]
// http DELETE "http://localhost:8080/hostdefinitions/hello world"
func DeleteHostDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteHostDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
