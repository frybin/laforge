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

func configCompetitionUsersRouter(router *httprouter.Router) {
	router.GET("/competitionusers", GetAllCompetitionUsers)
	router.POST("/competitionusers", AddCompetitionUsers)

	router.GET("/competitionusers/:argId", GetCompetitionUsers)
	router.PUT("/competitionusers/:argId", UpdateCompetitionUsers)
	router.DELETE("/competitionusers/:argId", DeleteCompetitionUsers)
}

func configGinCompetitionUsersRouter(router gin.IRoutes) {
	router.GET("/competitionusers", ConverHttprouterToGin(GetAllCompetitionUsers))
	router.POST("/competitionusers", ConverHttprouterToGin(AddCompetitionUsers))
	router.GET("/competitionusers/:argId", ConverHttprouterToGin(GetCompetitionUsers))
	router.PUT("/competitionusers/:argId", ConverHttprouterToGin(UpdateCompetitionUsers))
	router.DELETE("/competitionusers/:argId", ConverHttprouterToGin(DeleteCompetitionUsers))
}

// GetAllCompetitionUsers is a function to get a slice of record(s) from competition_users table in the laforge-2 database
// @Summary Get list of CompetitionUsers
// @Tags CompetitionUsers
// @Description GetAllCompetitionUsers is a handler to get a slice of record(s) from competition_users table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.CompetitionUsers}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /competitionusers [get]
// http "http://localhost:8080/competitionusers?page=0&pagesize=20"
func GetAllCompetitionUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllCompetitionUsers(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetCompetitionUsers is a function to get a single record from the competition_users table in the laforge-2 database
// @Summary Get record from table CompetitionUsers by  argId
// @Tags CompetitionUsers
// @ID argId
// @Description GetCompetitionUsers is a function to get a single record from the competition_users table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.CompetitionUsers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /competitionusers/{argId} [get]
// http "http://localhost:8080/competitionusers/hello world"
func GetCompetitionUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetCompetitionUsers(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddCompetitionUsers add to add a single record to competition_users table in the laforge-2 database
// @Summary Add an record to competition_users table
// @Description add to add a single record to competition_users table in the laforge-2 database
// @Tags CompetitionUsers
// @Accept  json
// @Produce  json
// @Param CompetitionUsers body model.CompetitionUsers true "Add CompetitionUsers"
// @Success 200 {object} model.CompetitionUsers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /competitionusers [post]
// echo '{"id": "keAX\\YoOqjGA[wP^KmAHd]wJY","users_id": "]iQVSeYZKvrNRSvZHsnxy]VJ`","competitions_id": "cjOloGdn`wShNBa`gIQR]LgSf"}' | http POST "http://localhost:8080/competitionusers"
func AddCompetitionUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	competitionusers := &model.CompetitionUsers{}

	if err := readJSON(r, competitionusers); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := competitionusers.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	competitionusers.Prepare()

	if err := competitionusers.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	competitionusers, _, err = dao.AddCompetitionUsers(r.Context(), competitionusers)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, competitionusers)
}

// UpdateCompetitionUsers Update a single record from competition_users table in the laforge-2 database
// @Summary Update an record in table competition_users
// @Description Update a single record from competition_users table in the laforge-2 database
// @Tags CompetitionUsers
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  CompetitionUsers body model.CompetitionUsers true "Update CompetitionUsers record"
// @Success 200 {object} model.CompetitionUsers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /competitionusers/{argId} [patch]
// echo '{"id": "keAX\\YoOqjGA[wP^KmAHd]wJY","users_id": "]iQVSeYZKvrNRSvZHsnxy]VJ`","competitions_id": "cjOloGdn`wShNBa`gIQR]LgSf"}' | http PUT "http://localhost:8080/competitionusers/hello world"
func UpdateCompetitionUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	competitionusers := &model.CompetitionUsers{}
	if err := readJSON(r, competitionusers); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := competitionusers.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	competitionusers.Prepare()

	if err := competitionusers.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	competitionusers, _, err = dao.UpdateCompetitionUsers(r.Context(),
		argId,
		competitionusers)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, competitionusers)
}

// DeleteCompetitionUsers Delete a single record from competition_users table in the laforge-2 database
// @Summary Delete a record from competition_users
// @Description Delete a single record from competition_users table in the laforge-2 database
// @Tags CompetitionUsers
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.CompetitionUsers
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /competitionusers/{argId} [delete]
// http DELETE "http://localhost:8080/competitionusers/hello world"
func DeleteCompetitionUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCompetitionUsers(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
