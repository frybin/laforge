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

func configProvisionedHostsRouter(router *httprouter.Router) {
	router.GET("/provisionedhosts", GetAllProvisionedHosts)
	router.POST("/provisionedhosts", AddProvisionedHosts)

	router.GET("/provisionedhosts/:argId", GetProvisionedHosts)
	router.PUT("/provisionedhosts/:argId", UpdateProvisionedHosts)
	router.DELETE("/provisionedhosts/:argId", DeleteProvisionedHosts)
}

func configGinProvisionedHostsRouter(router gin.IRoutes) {
	router.GET("/provisionedhosts", ConverHttprouterToGin(GetAllProvisionedHosts))
	router.POST("/provisionedhosts", ConverHttprouterToGin(AddProvisionedHosts))
	router.GET("/provisionedhosts/:argId", ConverHttprouterToGin(GetProvisionedHosts))
	router.PUT("/provisionedhosts/:argId", ConverHttprouterToGin(UpdateProvisionedHosts))
	router.DELETE("/provisionedhosts/:argId", ConverHttprouterToGin(DeleteProvisionedHosts))
}

// GetAllProvisionedHosts is a function to get a slice of record(s) from provisioned_hosts table in the laforge-2 database
// @Summary Get list of ProvisionedHosts
// @Tags ProvisionedHosts
// @Description GetAllProvisionedHosts is a handler to get a slice of record(s) from provisioned_hosts table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ProvisionedHosts}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedhosts [get]
// http "http://localhost:8080/provisionedhosts?page=0&pagesize=20"
func GetAllProvisionedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllProvisionedHosts(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetProvisionedHosts is a function to get a single record from the provisioned_hosts table in the laforge-2 database
// @Summary Get record from table ProvisionedHosts by  argId
// @Tags ProvisionedHosts
// @ID argId
// @Description GetProvisionedHosts is a function to get a single record from the provisioned_hosts table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.ProvisionedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /provisionedhosts/{argId} [get]
// http "http://localhost:8080/provisionedhosts/hello world"
func GetProvisionedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetProvisionedHosts(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddProvisionedHosts add to add a single record to provisioned_hosts table in the laforge-2 database
// @Summary Add an record to provisioned_hosts table
// @Description add to add a single record to provisioned_hosts table in the laforge-2 database
// @Tags ProvisionedHosts
// @Accept  json
// @Produce  json
// @Param ProvisionedHosts body model.ProvisionedHosts true "Add ProvisionedHosts"
// @Success 200 {object} model.ProvisionedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedhosts [post]
// echo '{"id": "JAAvXC`B]tiEiOYNXvcgs^A^d","included_hosts_id": "juDJNVMZcksaYydIcv^hRDAhB","provisioned_networks_id": "XrkiEOjmmqZ[`inFxeuvvmXQA","previous_checksum": "`HAJUrOvxFjdvtXbu[ZROHaoT","state": "fpA]Q^CLhKtLmor[Vkne]_T]K","ip_address": "cMVuIoRpQWk_souciv^uyVVCI","conn_info": "aSwedjiP^uELwJ`WixTiJmcBX","planned_checksum": "ajPvtCkw_SRjqFuBgU`q`YXwg","current_checksum": "j\\KYKmiFkgf`T\\rPgGrx]nEfl"}' | http POST "http://localhost:8080/provisionedhosts"
func AddProvisionedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	provisionedhosts := &model.ProvisionedHosts{}

	if err := readJSON(r, provisionedhosts); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionedhosts.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionedhosts.Prepare()

	if err := provisionedhosts.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	provisionedhosts, _, err = dao.AddProvisionedHosts(r.Context(), provisionedhosts)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionedhosts)
}

// UpdateProvisionedHosts Update a single record from provisioned_hosts table in the laforge-2 database
// @Summary Update an record in table provisioned_hosts
// @Description Update a single record from provisioned_hosts table in the laforge-2 database
// @Tags ProvisionedHosts
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  ProvisionedHosts body model.ProvisionedHosts true "Update ProvisionedHosts record"
// @Success 200 {object} model.ProvisionedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedhosts/{argId} [patch]
// echo '{"id": "JAAvXC`B]tiEiOYNXvcgs^A^d","included_hosts_id": "juDJNVMZcksaYydIcv^hRDAhB","provisioned_networks_id": "XrkiEOjmmqZ[`inFxeuvvmXQA","previous_checksum": "`HAJUrOvxFjdvtXbu[ZROHaoT","state": "fpA]Q^CLhKtLmor[Vkne]_T]K","ip_address": "cMVuIoRpQWk_souciv^uyVVCI","conn_info": "aSwedjiP^uELwJ`WixTiJmcBX","planned_checksum": "ajPvtCkw_SRjqFuBgU`q`YXwg","current_checksum": "j\\KYKmiFkgf`T\\rPgGrx]nEfl"}' | http PUT "http://localhost:8080/provisionedhosts/hello world"
func UpdateProvisionedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	provisionedhosts := &model.ProvisionedHosts{}
	if err := readJSON(r, provisionedhosts); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionedhosts.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionedhosts.Prepare()

	if err := provisionedhosts.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	provisionedhosts, _, err = dao.UpdateProvisionedHosts(r.Context(),
		argId,
		provisionedhosts)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionedhosts)
}

// DeleteProvisionedHosts Delete a single record from provisioned_hosts table in the laforge-2 database
// @Summary Delete a record from provisioned_hosts
// @Description Delete a single record from provisioned_hosts table in the laforge-2 database
// @Tags ProvisionedHosts
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.ProvisionedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /provisionedhosts/{argId} [delete]
// http DELETE "http://localhost:8080/provisionedhosts/hello world"
func DeleteProvisionedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProvisionedHosts(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
