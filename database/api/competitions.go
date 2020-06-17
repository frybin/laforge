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

func configCompetitionsRouter(router *httprouter.Router) {
	router.GET("/competitions", GetAllCompetitions)
	router.POST("/competitions", AddCompetitions)

	router.GET("/competitions/:argId", GetCompetitions)
	router.PUT("/competitions/:argId", UpdateCompetitions)
	router.DELETE("/competitions/:argId", DeleteCompetitions)
}

func configGinCompetitionsRouter(router gin.IRoutes) {
	router.GET("/competitions", ConverHttprouterToGin(GetAllCompetitions))
	router.POST("/competitions", ConverHttprouterToGin(AddCompetitions))
	router.GET("/competitions/:argId", ConverHttprouterToGin(GetCompetitions))
	router.PUT("/competitions/:argId", ConverHttprouterToGin(UpdateCompetitions))
	router.DELETE("/competitions/:argId", ConverHttprouterToGin(DeleteCompetitions))
}

// GetAllCompetitions is a function to get a slice of record(s) from competitions table in the laforge-2 database
// @Summary Get list of Competitions
// @Tags Competitions
// @Description GetAllCompetitions is a handler to get a slice of record(s) from competitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Competitions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /competitions [get]
// http "http://localhost:8080/competitions?page=0&pagesize=20"
func GetAllCompetitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllCompetitions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetCompetitions is a function to get a single record from the competitions table in the laforge-2 database
// @Summary Get record from table Competitions by  argId
// @Tags Competitions
// @ID argId
// @Description GetCompetitions is a function to get a single record from the competitions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.Competitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /competitions/{argId} [get]
// http "http://localhost:8080/competitions/hello world"
func GetCompetitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetCompetitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddCompetitions add to add a single record to competitions table in the laforge-2 database
// @Summary Add an record to competitions table
// @Description add to add a single record to competitions table in the laforge-2 database
// @Tags Competitions
// @Accept  json
// @Produce  json
// @Param Competitions body model.Competitions true "Add Competitions"
// @Success 200 {object} model.Competitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /competitions [post]
// echo '{"id": "ehHmVcRVU^Ox_ZLydHobaEZCR","name": "u^NYxiLywpu^GEc]pstPSoYeG"}' | http POST "http://localhost:8080/competitions"
func AddCompetitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	competitions := &model.Competitions{}

	if err := readJSON(r, competitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := competitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	competitions.Prepare()

	if err := competitions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	competitions, _, err = dao.AddCompetitions(r.Context(), competitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, competitions)
}

// UpdateCompetitions Update a single record from competitions table in the laforge-2 database
// @Summary Update an record in table competitions
// @Description Update a single record from competitions table in the laforge-2 database
// @Tags Competitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  Competitions body model.Competitions true "Update Competitions record"
// @Success 200 {object} model.Competitions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /competitions/{argId} [patch]
// echo '{"id": "ehHmVcRVU^Ox_ZLydHobaEZCR","name": "u^NYxiLywpu^GEc]pstPSoYeG"}' | http PUT "http://localhost:8080/competitions/hello world"
func UpdateCompetitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	competitions := &model.Competitions{}
	if err := readJSON(r, competitions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := competitions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	competitions.Prepare()

	if err := competitions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	competitions, _, err = dao.UpdateCompetitions(r.Context(),
		argId,
		competitions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, competitions)
}

// DeleteCompetitions Delete a single record from competitions table in the laforge-2 database
// @Summary Delete a record from competitions
// @Description Delete a single record from competitions table in the laforge-2 database
// @Tags Competitions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.Competitions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /competitions/{argId} [delete]
// http DELETE "http://localhost:8080/competitions/hello world"
func DeleteCompetitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCompetitions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
