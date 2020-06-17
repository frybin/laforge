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

func configIncludedNetworksRouter(router *httprouter.Router) {
	router.GET("/includednetworks", GetAllIncludedNetworks)
	router.POST("/includednetworks", AddIncludedNetworks)

	router.GET("/includednetworks/:argId", GetIncludedNetworks)
	router.PUT("/includednetworks/:argId", UpdateIncludedNetworks)
	router.DELETE("/includednetworks/:argId", DeleteIncludedNetworks)
}

func configGinIncludedNetworksRouter(router gin.IRoutes) {
	router.GET("/includednetworks", ConverHttprouterToGin(GetAllIncludedNetworks))
	router.POST("/includednetworks", ConverHttprouterToGin(AddIncludedNetworks))
	router.GET("/includednetworks/:argId", ConverHttprouterToGin(GetIncludedNetworks))
	router.PUT("/includednetworks/:argId", ConverHttprouterToGin(UpdateIncludedNetworks))
	router.DELETE("/includednetworks/:argId", ConverHttprouterToGin(DeleteIncludedNetworks))
}

// GetAllIncludedNetworks is a function to get a slice of record(s) from included_networks table in the laforge-2 database
// @Summary Get list of IncludedNetworks
// @Tags IncludedNetworks
// @Description GetAllIncludedNetworks is a handler to get a slice of record(s) from included_networks table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IncludedNetworks}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includednetworks [get]
// http "http://localhost:8080/includednetworks?page=0&pagesize=20"
func GetAllIncludedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllIncludedNetworks(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetIncludedNetworks is a function to get a single record from the included_networks table in the laforge-2 database
// @Summary Get record from table IncludedNetworks by  argId
// @Tags IncludedNetworks
// @ID argId
// @Description GetIncludedNetworks is a function to get a single record from the included_networks table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.IncludedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /includednetworks/{argId} [get]
// http "http://localhost:8080/includednetworks/hello world"
func GetIncludedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetIncludedNetworks(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddIncludedNetworks add to add a single record to included_networks table in the laforge-2 database
// @Summary Add an record to included_networks table
// @Description add to add a single record to included_networks table in the laforge-2 database
// @Tags IncludedNetworks
// @Accept  json
// @Produce  json
// @Param IncludedNetworks body model.IncludedNetworks true "Add IncludedNetworks"
// @Success 200 {object} model.IncludedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includednetworks [post]
// echo '{"planned_checksum": "xuDnNkXYUmmD[QKAUnrxJCcUb","current_checksum": "JWX]NeyEXDMebKkToyJ_k^xbG","previous_checksum": "DHckdmLqkNWZMfmicSn^qBK\\P","attrs": "\\esKsZB^EuNTkIHYRpgpnTb`E","id": "RsPOaReogFfQ^UH[jVDHKlSwA","environments_id": "rfAGFBeyqfn_dI[KAedpfCqyl","network_defintions_id": "Gom_giUubNnFopKtmTxkMJ]OR"}' | http POST "http://localhost:8080/includednetworks"
func AddIncludedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	includednetworks := &model.IncludedNetworks{}

	if err := readJSON(r, includednetworks); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includednetworks.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includednetworks.Prepare()

	if err := includednetworks.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	includednetworks, _, err = dao.AddIncludedNetworks(r.Context(), includednetworks)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includednetworks)
}

// UpdateIncludedNetworks Update a single record from included_networks table in the laforge-2 database
// @Summary Update an record in table included_networks
// @Description Update a single record from included_networks table in the laforge-2 database
// @Tags IncludedNetworks
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  IncludedNetworks body model.IncludedNetworks true "Update IncludedNetworks record"
// @Success 200 {object} model.IncludedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includednetworks/{argId} [patch]
// echo '{"planned_checksum": "xuDnNkXYUmmD[QKAUnrxJCcUb","current_checksum": "JWX]NeyEXDMebKkToyJ_k^xbG","previous_checksum": "DHckdmLqkNWZMfmicSn^qBK\\P","attrs": "\\esKsZB^EuNTkIHYRpgpnTb`E","id": "RsPOaReogFfQ^UH[jVDHKlSwA","environments_id": "rfAGFBeyqfn_dI[KAedpfCqyl","network_defintions_id": "Gom_giUubNnFopKtmTxkMJ]OR"}' | http PUT "http://localhost:8080/includednetworks/hello world"
func UpdateIncludedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	includednetworks := &model.IncludedNetworks{}
	if err := readJSON(r, includednetworks); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includednetworks.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includednetworks.Prepare()

	if err := includednetworks.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	includednetworks, _, err = dao.UpdateIncludedNetworks(r.Context(),
		argId,
		includednetworks)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includednetworks)
}

// DeleteIncludedNetworks Delete a single record from included_networks table in the laforge-2 database
// @Summary Delete a record from included_networks
// @Description Delete a single record from included_networks table in the laforge-2 database
// @Tags IncludedNetworks
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.IncludedNetworks
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /includednetworks/{argId} [delete]
// http DELETE "http://localhost:8080/includednetworks/hello world"
func DeleteIncludedNetworks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIncludedNetworks(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
