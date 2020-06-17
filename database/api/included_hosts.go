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

func configIncludedHostsRouter(router *httprouter.Router) {
	router.GET("/includedhosts", GetAllIncludedHosts)
	router.POST("/includedhosts", AddIncludedHosts)

	router.GET("/includedhosts/:argId", GetIncludedHosts)
	router.PUT("/includedhosts/:argId", UpdateIncludedHosts)
	router.DELETE("/includedhosts/:argId", DeleteIncludedHosts)
}

func configGinIncludedHostsRouter(router gin.IRoutes) {
	router.GET("/includedhosts", ConverHttprouterToGin(GetAllIncludedHosts))
	router.POST("/includedhosts", ConverHttprouterToGin(AddIncludedHosts))
	router.GET("/includedhosts/:argId", ConverHttprouterToGin(GetIncludedHosts))
	router.PUT("/includedhosts/:argId", ConverHttprouterToGin(UpdateIncludedHosts))
	router.DELETE("/includedhosts/:argId", ConverHttprouterToGin(DeleteIncludedHosts))
}

// GetAllIncludedHosts is a function to get a slice of record(s) from included_hosts table in the laforge-2 database
// @Summary Get list of IncludedHosts
// @Tags IncludedHosts
// @Description GetAllIncludedHosts is a handler to get a slice of record(s) from included_hosts table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IncludedHosts}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedhosts [get]
// http "http://localhost:8080/includedhosts?page=0&pagesize=20"
func GetAllIncludedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllIncludedHosts(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetIncludedHosts is a function to get a single record from the included_hosts table in the laforge-2 database
// @Summary Get record from table IncludedHosts by  argId
// @Tags IncludedHosts
// @ID argId
// @Description GetIncludedHosts is a function to get a single record from the included_hosts table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.IncludedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /includedhosts/{argId} [get]
// http "http://localhost:8080/includedhosts/hello world"
func GetIncludedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetIncludedHosts(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddIncludedHosts add to add a single record to included_hosts table in the laforge-2 database
// @Summary Add an record to included_hosts table
// @Description add to add a single record to included_hosts table in the laforge-2 database
// @Tags IncludedHosts
// @Accept  json
// @Produce  json
// @Param IncludedHosts body model.IncludedHosts true "Add IncludedHosts"
// @Success 200 {object} model.IncludedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedhosts [post]
// echo '{"attrs": "KPSnydV\\HgAEBRl`Ja[diamB]","id": "\\mHkwmvtJ^yRTQvmtxYmwaKrn","included_networks_id": "GDqVHJa[BDLkaLoDNwMfL_mv_","host_definitions_id": "[PBtvpwcR]X[HjmgIn[^dvLWq","planned_checksum": "hfICaX^cVmDbBlcb^JRltfFML","current_checksum": "m_tW`EBeXX^msAr[lBsxE[wZL","previous_checksum": "LsEuEb_en`qXIL[r`^ZL_UiEw"}' | http POST "http://localhost:8080/includedhosts"
func AddIncludedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	includedhosts := &model.IncludedHosts{}

	if err := readJSON(r, includedhosts); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includedhosts.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includedhosts.Prepare()

	if err := includedhosts.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	includedhosts, _, err = dao.AddIncludedHosts(r.Context(), includedhosts)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includedhosts)
}

// UpdateIncludedHosts Update a single record from included_hosts table in the laforge-2 database
// @Summary Update an record in table included_hosts
// @Description Update a single record from included_hosts table in the laforge-2 database
// @Tags IncludedHosts
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  IncludedHosts body model.IncludedHosts true "Update IncludedHosts record"
// @Success 200 {object} model.IncludedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedhosts/{argId} [patch]
// echo '{"attrs": "KPSnydV\\HgAEBRl`Ja[diamB]","id": "\\mHkwmvtJ^yRTQvmtxYmwaKrn","included_networks_id": "GDqVHJa[BDLkaLoDNwMfL_mv_","host_definitions_id": "[PBtvpwcR]X[HjmgIn[^dvLWq","planned_checksum": "hfICaX^cVmDbBlcb^JRltfFML","current_checksum": "m_tW`EBeXX^msAr[lBsxE[wZL","previous_checksum": "LsEuEb_en`qXIL[r`^ZL_UiEw"}' | http PUT "http://localhost:8080/includedhosts/hello world"
func UpdateIncludedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	includedhosts := &model.IncludedHosts{}
	if err := readJSON(r, includedhosts); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includedhosts.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includedhosts.Prepare()

	if err := includedhosts.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	includedhosts, _, err = dao.UpdateIncludedHosts(r.Context(),
		argId,
		includedhosts)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includedhosts)
}

// DeleteIncludedHosts Delete a single record from included_hosts table in the laforge-2 database
// @Summary Delete a record from included_hosts
// @Description Delete a single record from included_hosts table in the laforge-2 database
// @Tags IncludedHosts
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.IncludedHosts
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /includedhosts/{argId} [delete]
// http DELETE "http://localhost:8080/includedhosts/hello world"
func DeleteIncludedHosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIncludedHosts(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
