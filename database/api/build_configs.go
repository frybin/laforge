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

func configBuildConfigsRouter(router *httprouter.Router) {
	router.GET("/buildconfigs", GetAllBuildConfigs)
	router.POST("/buildconfigs", AddBuildConfigs)

	router.GET("/buildconfigs/:argId", GetBuildConfigs)
	router.PUT("/buildconfigs/:argId", UpdateBuildConfigs)
	router.DELETE("/buildconfigs/:argId", DeleteBuildConfigs)
}

func configGinBuildConfigsRouter(router gin.IRoutes) {
	router.GET("/buildconfigs", ConverHttprouterToGin(GetAllBuildConfigs))
	router.POST("/buildconfigs", ConverHttprouterToGin(AddBuildConfigs))
	router.GET("/buildconfigs/:argId", ConverHttprouterToGin(GetBuildConfigs))
	router.PUT("/buildconfigs/:argId", ConverHttprouterToGin(UpdateBuildConfigs))
	router.DELETE("/buildconfigs/:argId", ConverHttprouterToGin(DeleteBuildConfigs))
}

// GetAllBuildConfigs is a function to get a slice of record(s) from build_configs table in the laforge-2 database
// @Summary Get list of BuildConfigs
// @Tags BuildConfigs
// @Description GetAllBuildConfigs is a handler to get a slice of record(s) from build_configs table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BuildConfigs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildconfigs [get]
// http "http://localhost:8080/buildconfigs?page=0&pagesize=20"
func GetAllBuildConfigs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllBuildConfigs(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetBuildConfigs is a function to get a single record from the build_configs table in the laforge-2 database
// @Summary Get record from table BuildConfigs by  argId
// @Tags BuildConfigs
// @ID argId
// @Description GetBuildConfigs is a function to get a single record from the build_configs table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.BuildConfigs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /buildconfigs/{argId} [get]
// http "http://localhost:8080/buildconfigs/hello world"
func GetBuildConfigs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetBuildConfigs(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddBuildConfigs add to add a single record to build_configs table in the laforge-2 database
// @Summary Add an record to build_configs table
// @Description add to add a single record to build_configs table in the laforge-2 database
// @Tags BuildConfigs
// @Accept  json
// @Produce  json
// @Param BuildConfigs body model.BuildConfigs true "Add BuildConfigs"
// @Success 200 {object} model.BuildConfigs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildconfigs [post]
// echo '{"id": "RKlvbHgm\\R\\IZGcpGESIZrc[n","provider": "WHGdkF\\a_kibJUcYbS\\CkFAAq","attrs": "FMB[[]xNmyIZeOdAoh]ZckE]g","competitions_id": "wXgZPmHKHajpVvEIqg_^dJc\\O"}' | http POST "http://localhost:8080/buildconfigs"
func AddBuildConfigs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	buildconfigs := &model.BuildConfigs{}

	if err := readJSON(r, buildconfigs); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := buildconfigs.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	buildconfigs.Prepare()

	if err := buildconfigs.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	buildconfigs, _, err = dao.AddBuildConfigs(r.Context(), buildconfigs)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, buildconfigs)
}

// UpdateBuildConfigs Update a single record from build_configs table in the laforge-2 database
// @Summary Update an record in table build_configs
// @Description Update a single record from build_configs table in the laforge-2 database
// @Tags BuildConfigs
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  BuildConfigs body model.BuildConfigs true "Update BuildConfigs record"
// @Success 200 {object} model.BuildConfigs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildconfigs/{argId} [patch]
// echo '{"id": "RKlvbHgm\\R\\IZGcpGESIZrc[n","provider": "WHGdkF\\a_kibJUcYbS\\CkFAAq","attrs": "FMB[[]xNmyIZeOdAoh]ZckE]g","competitions_id": "wXgZPmHKHajpVvEIqg_^dJc\\O"}' | http PUT "http://localhost:8080/buildconfigs/hello world"
func UpdateBuildConfigs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	buildconfigs := &model.BuildConfigs{}
	if err := readJSON(r, buildconfigs); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := buildconfigs.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	buildconfigs.Prepare()

	if err := buildconfigs.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	buildconfigs, _, err = dao.UpdateBuildConfigs(r.Context(),
		argId,
		buildconfigs)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, buildconfigs)
}

// DeleteBuildConfigs Delete a single record from build_configs table in the laforge-2 database
// @Summary Delete a record from build_configs
// @Description Delete a single record from build_configs table in the laforge-2 database
// @Tags BuildConfigs
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.BuildConfigs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /buildconfigs/{argId} [delete]
// http DELETE "http://localhost:8080/buildconfigs/hello world"
func DeleteBuildConfigs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBuildConfigs(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
