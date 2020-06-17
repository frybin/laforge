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

func configIdentityDefinitionsRouter(router *httprouter.Router) {
	router.GET("/identitydefinitions", GetAllIdentityDefinitions)
	router.POST("/identitydefinitions", AddIdentityDefinitions)

	router.GET("/identitydefinitions/:argId", GetIdentityDefinitions)
	router.PUT("/identitydefinitions/:argId", UpdateIdentityDefinitions)
	router.DELETE("/identitydefinitions/:argId", DeleteIdentityDefinitions)
}

func configGinIdentityDefinitionsRouter(router gin.IRoutes) {
	router.GET("/identitydefinitions", ConverHttprouterToGin(GetAllIdentityDefinitions))
	router.POST("/identitydefinitions", ConverHttprouterToGin(AddIdentityDefinitions))
	router.GET("/identitydefinitions/:argId", ConverHttprouterToGin(GetIdentityDefinitions))
	router.PUT("/identitydefinitions/:argId", ConverHttprouterToGin(UpdateIdentityDefinitions))
	router.DELETE("/identitydefinitions/:argId", ConverHttprouterToGin(DeleteIdentityDefinitions))
}

// GetAllIdentityDefinitions is a function to get a slice of record(s) from identity_definitions table in the laforge-2 database
// @Summary Get list of IdentityDefinitions
// @Tags IdentityDefinitions
// @Description GetAllIdentityDefinitions is a handler to get a slice of record(s) from identity_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IdentityDefinitions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identitydefinitions [get]
// http "http://localhost:8080/identitydefinitions?page=0&pagesize=20"
func GetAllIdentityDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllIdentityDefinitions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetIdentityDefinitions is a function to get a single record from the identity_definitions table in the laforge-2 database
// @Summary Get record from table IdentityDefinitions by  argId
// @Tags IdentityDefinitions
// @ID argId
// @Description GetIdentityDefinitions is a function to get a single record from the identity_definitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.IdentityDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /identitydefinitions/{argId} [get]
// http "http://localhost:8080/identitydefinitions/hello world"
func GetIdentityDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetIdentityDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddIdentityDefinitions add to add a single record to identity_definitions table in the laforge-2 database
// @Summary Add an record to identity_definitions table
// @Description add to add a single record to identity_definitions table in the laforge-2 database
// @Tags IdentityDefinitions
// @Accept  json
// @Produce  json
// @Param IdentityDefinitions body model.IdentityDefinitions true "Add IdentityDefinitions"
// @Success 200 {object} model.IdentityDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identitydefinitions [post]
// echo '{"email": "FOcRFhsfUB`QmVkBThWgSgMbE","id": "PkWjMq^vlE_JTxOuIxfKxZIsP"}' | http POST "http://localhost:8080/identitydefinitions"
func AddIdentityDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	identitydefinitions := &model.IdentityDefinitions{}

	if err := readJSON(r, identitydefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := identitydefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	identitydefinitions.Prepare()

	if err := identitydefinitions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	identitydefinitions, _, err = dao.AddIdentityDefinitions(r.Context(), identitydefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, identitydefinitions)
}

// UpdateIdentityDefinitions Update a single record from identity_definitions table in the laforge-2 database
// @Summary Update an record in table identity_definitions
// @Description Update a single record from identity_definitions table in the laforge-2 database
// @Tags IdentityDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  IdentityDefinitions body model.IdentityDefinitions true "Update IdentityDefinitions record"
// @Success 200 {object} model.IdentityDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identitydefinitions/{argId} [patch]
// echo '{"email": "FOcRFhsfUB`QmVkBThWgSgMbE","id": "PkWjMq^vlE_JTxOuIxfKxZIsP"}' | http PUT "http://localhost:8080/identitydefinitions/hello world"
func UpdateIdentityDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	identitydefinitions := &model.IdentityDefinitions{}
	if err := readJSON(r, identitydefinitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := identitydefinitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	identitydefinitions.Prepare()

	if err := identitydefinitions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	identitydefinitions, _, err = dao.UpdateIdentityDefinitions(r.Context(),
		argId,
		identitydefinitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, identitydefinitions)
}

// DeleteIdentityDefinitions Delete a single record from identity_definitions table in the laforge-2 database
// @Summary Delete a record from identity_definitions
// @Description Delete a single record from identity_definitions table in the laforge-2 database
// @Tags IdentityDefinitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.IdentityDefinitions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /identitydefinitions/{argId} [delete]
// http DELETE "http://localhost:8080/identitydefinitions/hello world"
func DeleteIdentityDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIdentityDefinitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
