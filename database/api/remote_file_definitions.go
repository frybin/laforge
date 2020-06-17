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

func configRemoteFileDefinitionsRouter(router *httprouter.Router) {
	router.GET("/remotefiledefinitions", GetAllRemoteFileDefinitions)
	router.POST("/remotefiledefinitions", AddRemoteFileDefinitions)

	router.GET("/remotefiledefinitions/:argId", GetRemoteFileDefinitions)
	router.PUT("/remotefiledefinitions/:argId", UpdateRemoteFileDefinitions)
	router.DELETE("/remotefiledefinitions/:argId", DeleteRemoteFileDefinitions)
}

func configGinRemoteFileDefinitionsRouter(router gin.IRoutes) {
	router.GET("/remotefiledefinitions", ConverHttprouterToGin(GetAllRemoteFileDefinitions))
	router.POST("/remotefiledefinitions", ConverHttprouterToGin(AddRemoteFileDefinitions))
	router.GET("/remotefiledefinitions/:argId", ConverHttprouterToGin(GetRemoteFileDefinitions))
	router.PUT("/remotefiledefinitions/:argId", ConverHttprouterToGin(UpdateRemoteFileDefinitions))
	router.DELETE("/remotefiledefinitions/:argId", ConverHttprouterToGin(DeleteRemoteFileDefinitions))
}

// GetAllRemoteFileDefinitions is a function to get a slice of record(s) from remote_file_definitions table in the laforge-2 database
// @Summary Get list of RemoteFileDefinitions
// @Tags RemoteFileDefinitions
// @Description GetAllRemoteFileDefinitions is a handler to get a slice of record(s) from remote_file_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RemoteFileDefinitions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /remotefiledefinitions [get]
// http "http://localhost:8080/remotefiledefinitions?page=0&pagesize=20"
func GetAllRemoteFileDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllRemoteFileDefinitions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetRemoteFileDefinitions is a function to get a single record from the remote_file_definitions table in the laforge-2 database
// @Summary Get record from table RemoteFileDefinitions by  argId
// @Tags RemoteFileDefinitions
// @ID argId
// @Description GetRemoteFileDefinitions is a function to get a single record from the remote_file_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.RemoteFileDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /remotefiledefinitions/{argId} [get]
// http "http://localhost:8080/remotefiledefinitions/hello world"
func GetRemoteFileDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetRemoteFileDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddRemoteFileDefinitions add to add a single record to remote_file_definitions table in the laforge-2 database
// @Summary Add an record to remote_file_definitions table
// @Description add to add a single record to remote_file_definitions table in the laforge-2 database
// @Tags RemoteFileDefinitions
// @Accept  json
// @Produce  json
// @Param RemoteFileDefinitions body model.RemoteFileDefinitions true "Add RemoteFileDefinitions"
// @Success 200 {object} model.RemoteFileDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /remotefiledefinitions [post]
// echo '{"id": "bqNZxFuHHVwHXeyolFhnrVqaV","type": "yn`TpXkHrPdej\\ZUUgliUVBhd","source": "VH]kWpp`u`ErCZqx\\imCVA^xX"}' | http POST "http://localhost:8080/remotefiledefinitions"
func AddRemoteFileDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	remotefiledefinitions := &model.RemoteFileDefinitions{}

	if err := readJSON(r, remotefiledefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := remotefiledefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	remotefiledefinitions.Prepare()

	if err := remotefiledefinitions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	remotefiledefinitions, _, err = dao.AddRemoteFileDefinitions(r.Context(), remotefiledefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, remotefiledefinitions)
}

// UpdateRemoteFileDefinitions Update a single record from remote_file_definitions table in the laforge-2 database
// @Summary Update an record in table remote_file_definitions
// @Description Update a single record from remote_file_definitions table in the laforge-2 database
// @Tags RemoteFileDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  RemoteFileDefinitions body model.RemoteFileDefinitions true "Update RemoteFileDefinitions record"
// @Success 200 {object} model.RemoteFileDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /remotefiledefinitions/{argId} [patch]
// echo '{"id": "bqNZxFuHHVwHXeyolFhnrVqaV","type": "yn`TpXkHrPdej\\ZUUgliUVBhd","source": "VH]kWpp`u`ErCZqx\\imCVA^xX"}' | http PUT "http://localhost:8080/remotefiledefinitions/hello world"
func UpdateRemoteFileDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	remotefiledefinitions := &model.RemoteFileDefinitions{}
	if err := readJSON(r, remotefiledefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := remotefiledefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	remotefiledefinitions.Prepare()

	if err := remotefiledefinitions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	remotefiledefinitions, _, err = dao.UpdateRemoteFileDefinitions(r.Context(),
		argId,
		remotefiledefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, remotefiledefinitions)
}

// DeleteRemoteFileDefinitions Delete a single record from remote_file_definitions table in the laforge-2 database
// @Summary Delete a record from remote_file_definitions
// @Description Delete a single record from remote_file_definitions table in the laforge-2 database
// @Tags RemoteFileDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.RemoteFileDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /remotefiledefinitions/{argId} [delete]
// http DELETE "http://localhost:8080/remotefiledefinitions/hello world"
func DeleteRemoteFileDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRemoteFileDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
