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

func configDNSRecordDefinitionsRouter(router *httprouter.Router) {
	router.GET("/dnsrecorddefinitions", GetAllDNSRecordDefinitions)
	router.POST("/dnsrecorddefinitions", AddDNSRecordDefinitions)

	router.GET("/dnsrecorddefinitions/:argId", GetDNSRecordDefinitions)
	router.PUT("/dnsrecorddefinitions/:argId", UpdateDNSRecordDefinitions)
	router.DELETE("/dnsrecorddefinitions/:argId", DeleteDNSRecordDefinitions)
}

func configGinDNSRecordDefinitionsRouter(router gin.IRoutes) {
	router.GET("/dnsrecorddefinitions", ConverHttprouterToGin(GetAllDNSRecordDefinitions))
	router.POST("/dnsrecorddefinitions", ConverHttprouterToGin(AddDNSRecordDefinitions))
	router.GET("/dnsrecorddefinitions/:argId", ConverHttprouterToGin(GetDNSRecordDefinitions))
	router.PUT("/dnsrecorddefinitions/:argId", ConverHttprouterToGin(UpdateDNSRecordDefinitions))
	router.DELETE("/dnsrecorddefinitions/:argId", ConverHttprouterToGin(DeleteDNSRecordDefinitions))
}

// GetAllDNSRecordDefinitions is a function to get a slice of record(s) from dns_record_definitions table in the laforge-2 database
// @Summary Get list of DNSRecordDefinitions
// @Tags DNSRecordDefinitions
// @Description GetAllDNSRecordDefinitions is a handler to get a slice of record(s) from dns_record_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.DNSRecordDefinitions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dnsrecorddefinitions [get]
// http "http://localhost:8080/dnsrecorddefinitions?page=0&pagesize=20"
func GetAllDNSRecordDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllDNSRecordDefinitions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetDNSRecordDefinitions is a function to get a single record from the dns_record_definitions table in the laforge-2 database
// @Summary Get record from table DNSRecordDefinitions by  argId
// @Tags DNSRecordDefinitions
// @ID argId
// @Description GetDNSRecordDefinitions is a function to get a single record from the dns_record_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.DNSRecordDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /dnsrecorddefinitions/{argId} [get]
// http "http://localhost:8080/dnsrecorddefinitions/hello world"
func GetDNSRecordDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetDNSRecordDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddDNSRecordDefinitions add to add a single record to dns_record_definitions table in the laforge-2 database
// @Summary Add an record to dns_record_definitions table
// @Description add to add a single record to dns_record_definitions table in the laforge-2 database
// @Tags DNSRecordDefinitions
// @Accept  json
// @Produce  json
// @Param DNSRecordDefinitions body model.DNSRecordDefinitions true "Add DNSRecordDefinitions"
// @Success 200 {object} model.DNSRecordDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dnsrecorddefinitions [post]
// echo '{"id": "VMe`B^LRjVLstXodfy_xpJVHr","type": "ChZTAbTTEwxuKbK]uqQYPSIic","name": "_ei\\^kWyFUnmLuaWPwQBZBFup"}' | http POST "http://localhost:8080/dnsrecorddefinitions"
func AddDNSRecordDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dnsrecorddefinitions := &model.DNSRecordDefinitions{}

	if err := readJSON(r, dnsrecorddefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := dnsrecorddefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	dnsrecorddefinitions.Prepare()

	if err := dnsrecorddefinitions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	dnsrecorddefinitions, _, err = dao.AddDNSRecordDefinitions(r.Context(), dnsrecorddefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, dnsrecorddefinitions)
}

// UpdateDNSRecordDefinitions Update a single record from dns_record_definitions table in the laforge-2 database
// @Summary Update an record in table dns_record_definitions
// @Description Update a single record from dns_record_definitions table in the laforge-2 database
// @Tags DNSRecordDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  DNSRecordDefinitions body model.DNSRecordDefinitions true "Update DNSRecordDefinitions record"
// @Success 200 {object} model.DNSRecordDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dnsrecorddefinitions/{argId} [patch]
// echo '{"id": "VMe`B^LRjVLstXodfy_xpJVHr","type": "ChZTAbTTEwxuKbK]uqQYPSIic","name": "_ei\\^kWyFUnmLuaWPwQBZBFup"}' | http PUT "http://localhost:8080/dnsrecorddefinitions/hello world"
func UpdateDNSRecordDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	dnsrecorddefinitions := &model.DNSRecordDefinitions{}
	if err := readJSON(r, dnsrecorddefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := dnsrecorddefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	dnsrecorddefinitions.Prepare()

	if err := dnsrecorddefinitions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	dnsrecorddefinitions, _, err = dao.UpdateDNSRecordDefinitions(r.Context(),
		argId,
		dnsrecorddefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, dnsrecorddefinitions)
}

// DeleteDNSRecordDefinitions Delete a single record from dns_record_definitions table in the laforge-2 database
// @Summary Delete a record from dns_record_definitions
// @Description Delete a single record from dns_record_definitions table in the laforge-2 database
// @Tags DNSRecordDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.DNSRecordDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dnsrecorddefinitions/{argId} [delete]
// http DELETE "http://localhost:8080/dnsrecorddefinitions/hello world"
func DeleteDNSRecordDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDNSRecordDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
