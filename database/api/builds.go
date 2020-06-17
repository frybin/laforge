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

func configBuildsRouter(router *httprouter.Router) {
	router.GET("/builds", GetAllBuilds)
	router.POST("/builds", AddBuilds)

	router.GET("/builds/:argId", GetBuilds)
	router.PUT("/builds/:argId", UpdateBuilds)
	router.DELETE("/builds/:argId", DeleteBuilds)
}

func configGinBuildsRouter(router gin.IRoutes) {
	router.GET("/builds", ConverHttprouterToGin(GetAllBuilds))
	router.POST("/builds", ConverHttprouterToGin(AddBuilds))
	router.GET("/builds/:argId", ConverHttprouterToGin(GetBuilds))
	router.PUT("/builds/:argId", ConverHttprouterToGin(UpdateBuilds))
	router.DELETE("/builds/:argId", ConverHttprouterToGin(DeleteBuilds))
}

// GetAllBuilds is a function to get a slice of record(s) from builds table in the laforge-2 database
// @Summary Get list of Builds
// @Tags Builds
// @Description GetAllBuilds is a handler to get a slice of record(s) from builds table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Builds}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /builds [get]
// http "http://localhost:8080/builds?page=0&pagesize=20"
func GetAllBuilds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllBuilds(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetBuilds is a function to get a single record from the builds table in the laforge-2 database
// @Summary Get record from table Builds by  argId
// @Tags Builds
// @ID argId
// @Description GetBuilds is a function to get a single record from the builds table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.Builds
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /builds/{argId} [get]
// http "http://localhost:8080/builds/hello world"
func GetBuilds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetBuilds(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddBuilds add to add a single record to builds table in the laforge-2 database
// @Summary Add an record to builds table
// @Description add to add a single record to builds table in the laforge-2 database
// @Tags Builds
// @Accept  json
// @Produce  json
// @Param Builds body model.Builds true "Add Builds"
// @Success 200 {object} model.Builds
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /builds [post]
// echo '{"current_checksum": "R_WSBQrlfohQ[AuZaq^ohwLth","previous_checksum": "GQLcZsq_XgWGdfKIajTsSxmci","id": "bhTVVMoYJFj^rsPTSMBjVJfVg","environments_id": "gXksQEatneXZd`fnL^tFTnVOG","build_configs_id": "HUKwDsjakc^nPZOXvwqC]QYbO","state": "wKOVbZ[FhVKfyJstrcieJDHwh","planned_checksum": "[e]oQu`RRIHTU]NMk]clVA\\bm"}' | http POST "http://localhost:8080/builds"
func AddBuilds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	builds := &model.Builds{}

	if err := readJSON(r, builds); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := builds.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	builds.Prepare()

	if err := builds.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	builds, _, err = dao.AddBuilds(r.Context(), builds)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, builds)
}

// UpdateBuilds Update a single record from builds table in the laforge-2 database
// @Summary Update an record in table builds
// @Description Update a single record from builds table in the laforge-2 database
// @Tags Builds
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  Builds body model.Builds true "Update Builds record"
// @Success 200 {object} model.Builds
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /builds/{argId} [patch]
// echo '{"current_checksum": "R_WSBQrlfohQ[AuZaq^ohwLth","previous_checksum": "GQLcZsq_XgWGdfKIajTsSxmci","id": "bhTVVMoYJFj^rsPTSMBjVJfVg","environments_id": "gXksQEatneXZd`fnL^tFTnVOG","build_configs_id": "HUKwDsjakc^nPZOXvwqC]QYbO","state": "wKOVbZ[FhVKfyJstrcieJDHwh","planned_checksum": "[e]oQu`RRIHTU]NMk]clVA\\bm"}' | http PUT "http://localhost:8080/builds/hello world"
func UpdateBuilds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	builds := &model.Builds{}
	if err := readJSON(r, builds); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := builds.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	builds.Prepare()

	if err := builds.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	builds, _, err = dao.UpdateBuilds(r.Context(),
		argId,
		builds)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, builds)
}

// DeleteBuilds Delete a single record from builds table in the laforge-2 database
// @Summary Delete a record from builds
// @Description Delete a single record from builds table in the laforge-2 database
// @Tags Builds
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.Builds
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /builds/{argId} [delete]
// http DELETE "http://localhost:8080/builds/hello world"
func DeleteBuilds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBuilds(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
