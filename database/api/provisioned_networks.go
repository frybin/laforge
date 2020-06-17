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

func configProvisionedNetworksRouter(router *httprouter.Router) {
	router.GET("/provisionednetworks", GetAllProvisionedNetworks)
	router.POST("/provisionednetworks", AddProvisionedNetworks)

	router.GET("/provisionednetworks/:argId", GetProvisionedNetworks)
	router.PUT("/provisionednetworks/:argId", UpdateProvisionedNetworks)
	router.DELETE("/provisionednetworks/:argId", DeleteProvisionedNetworks)
}

func configGinProvisionedNetworksRouter(router gin.IRoutes) {
	router.GET("/provisionednetworks", ConverHttprouterToGin(GetAllProvisionedNetworks))
	router.POST("/provisionednetworks", ConverHttprouterToGin(AddProvisionedNetworks))
	router.GET("/provisionednetworks/:argId", ConverHttprouterToGin(GetProvisionedNetworks))
	router.PUT("/provisionednetworks/:argId", ConverHttprouterToGin(UpdateProvisionedNetworks))
	router.DELETE("/provisionednetworks/:argId", ConverHttprouterToGin(DeleteProvisionedNetworks))
}

// GetAllProvisionedNetworks is a function to get a slice of record(s) from provisioned_networks table in the laforge-2 database
// @Summary Get list of ProvisionedNetworks
// @Tags ProvisionedNetworks
// @Description GetAllProvisionedNetworks is a handler to get a slice of record(s) from provisioned_networks table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ProvisionedNetworks}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionednetworks [get]
// http "http://localhost:8080/provisionednetworks?page=0&pagesize=20"
func GetAllProvisionedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllProvisionedNetworks(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetProvisionedNetworks is a function to get a single record from the provisioned_networks table in the laforge-2 database
// @Summary Get record from table ProvisionedNetworks by  argId
// @Tags ProvisionedNetworks
// @ID argId
// @Description GetProvisionedNetworks is a function to get a single record from the provisioned_networks table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.ProvisionedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /provisionednetworks/{argId} [get]
// http "http://localhost:8080/provisionednetworks/hello world"
func GetProvisionedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetProvisionedNetworks(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddProvisionedNetworks add to add a single record to provisioned_networks table in the laforge-2 database
// @Summary Add an record to provisioned_networks table
// @Description add to add a single record to provisioned_networks table in the laforge-2 database
// @Tags ProvisionedNetworks
// @Accept  json
// @Produce  json
// @Param ProvisionedNetworks body model.ProvisionedNetworks true "Add ProvisionedNetworks"
// @Success 200 {object} model.ProvisionedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionednetworks [post]
// echo '{"planned_checksum": "malUpigVnBYER\\`ml^C_BY`Ad","current_checksum": "TXmnq]_kjlgn`BeLtihA^HNv`","previous_checksum": "OIIL]LkAqonGkDyIII`RBUolt","id": "PDjYHKOu`DSqZJGoVidQHkku`","included_networks_id": "eN]qDgmljmFolCDpISjjFOyQM","included_teams_id": "KtDxvesPEj`mxJhDGjxwRhaXE","state": "UEWpmD\\^GgRaqbdMmauyZPHsq"}' | http POST "http://localhost:8080/provisionednetworks"
func AddProvisionedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	provisionednetworks := &model.ProvisionedNetworks{}

	if err := readJSON(r, provisionednetworks); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionednetworks.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionednetworks.Prepare()

	if err := provisionednetworks.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	provisionednetworks, _, err = dao.AddProvisionedNetworks(r.Context(), provisionednetworks)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionednetworks)
}

// UpdateProvisionedNetworks Update a single record from provisioned_networks table in the laforge-2 database
// @Summary Update an record in table provisioned_networks
// @Description Update a single record from provisioned_networks table in the laforge-2 database
// @Tags ProvisionedNetworks
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  ProvisionedNetworks body model.ProvisionedNetworks true "Update ProvisionedNetworks record"
// @Success 200 {object} model.ProvisionedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionednetworks/{argId} [patch]
// echo '{"planned_checksum": "malUpigVnBYER\\`ml^C_BY`Ad","current_checksum": "TXmnq]_kjlgn`BeLtihA^HNv`","previous_checksum": "OIIL]LkAqonGkDyIII`RBUolt","id": "PDjYHKOu`DSqZJGoVidQHkku`","included_networks_id": "eN]qDgmljmFolCDpISjjFOyQM","included_teams_id": "KtDxvesPEj`mxJhDGjxwRhaXE","state": "UEWpmD\\^GgRaqbdMmauyZPHsq"}' | http PUT "http://localhost:8080/provisionednetworks/hello world"
func UpdateProvisionedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	provisionednetworks := &model.ProvisionedNetworks{}
	if err := readJSON(r, provisionednetworks); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionednetworks.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionednetworks.Prepare()

	if err := provisionednetworks.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	provisionednetworks, _, err = dao.UpdateProvisionedNetworks(r.Context(),
		argId,
		provisionednetworks)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionednetworks)
}

// DeleteProvisionedNetworks Delete a single record from provisioned_networks table in the laforge-2 database
// @Summary Delete a record from provisioned_networks
// @Description Delete a single record from provisioned_networks table in the laforge-2 database
// @Tags ProvisionedNetworks
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.ProvisionedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /provisionednetworks/{argId} [delete]
// http DELETE "http://localhost:8080/provisionednetworks/hello world"
func DeleteProvisionedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProvisionedNetworks(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
