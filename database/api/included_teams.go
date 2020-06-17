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

func configIncludedTeamsRouter(router *httprouter.Router) {
	router.GET("/includedteams", GetAllIncludedTeams)
	router.POST("/includedteams", AddIncludedTeams)

	router.GET("/includedteams/:argId", GetIncludedTeams)
	router.PUT("/includedteams/:argId", UpdateIncludedTeams)
	router.DELETE("/includedteams/:argId", DeleteIncludedTeams)
}

func configGinIncludedTeamsRouter(router gin.IRoutes) {
	router.GET("/includedteams", ConverHttprouterToGin(GetAllIncludedTeams))
	router.POST("/includedteams", ConverHttprouterToGin(AddIncludedTeams))
	router.GET("/includedteams/:argId", ConverHttprouterToGin(GetIncludedTeams))
	router.PUT("/includedteams/:argId", ConverHttprouterToGin(UpdateIncludedTeams))
	router.DELETE("/includedteams/:argId", ConverHttprouterToGin(DeleteIncludedTeams))
}

// GetAllIncludedTeams is a function to get a slice of record(s) from included_teams table in the laforge-2 database
// @Summary Get list of IncludedTeams
// @Tags IncludedTeams
// @Description GetAllIncludedTeams is a handler to get a slice of record(s) from included_teams table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IncludedTeams}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedteams [get]
// http "http://localhost:8080/includedteams?page=0&pagesize=20"
func GetAllIncludedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllIncludedTeams(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetIncludedTeams is a function to get a single record from the included_teams table in the laforge-2 database
// @Summary Get record from table IncludedTeams by  argId
// @Tags IncludedTeams
// @ID argId
// @Description GetIncludedTeams is a function to get a single record from the included_teams table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.IncludedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /includedteams/{argId} [get]
// http "http://localhost:8080/includedteams/hello world"
func GetIncludedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetIncludedTeams(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddIncludedTeams add to add a single record to included_teams table in the laforge-2 database
// @Summary Add an record to included_teams table
// @Description add to add a single record to included_teams table in the laforge-2 database
// @Tags IncludedTeams
// @Accept  json
// @Produce  json
// @Param IncludedTeams body model.IncludedTeams true "Add IncludedTeams"
// @Success 200 {object} model.IncludedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedteams [post]
// echo '{"tags": "W_WMSpC[Jr[_GwScoUYkmIkKX","enabled": false,"planned_checksum": "HTG]cNkfjeHxNjo_OX\\uEGuvB","current_checksum": "ewemk[mPBabZyjhJMNiIqjq`r","previous_checksum": "[GnKo_l[NYM`xs^OcGPOwjQ^^","id": "\\aovSpsbWukcBXrK\\kaILrxdr","environments_id": "lGSL[QWnu]xilrnoi_FlSYNnQ","number": 97}' | http POST "http://localhost:8080/includedteams"
func AddIncludedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	includedteams := &model.IncludedTeams{}

	if err := readJSON(r, includedteams); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includedteams.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includedteams.Prepare()

	if err := includedteams.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	includedteams, _, err = dao.AddIncludedTeams(r.Context(), includedteams)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includedteams)
}

// UpdateIncludedTeams Update a single record from included_teams table in the laforge-2 database
// @Summary Update an record in table included_teams
// @Description Update a single record from included_teams table in the laforge-2 database
// @Tags IncludedTeams
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  IncludedTeams body model.IncludedTeams true "Update IncludedTeams record"
// @Success 200 {object} model.IncludedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /includedteams/{argId} [patch]
// echo '{"tags": "W_WMSpC[Jr[_GwScoUYkmIkKX","enabled": false,"planned_checksum": "HTG]cNkfjeHxNjo_OX\\uEGuvB","current_checksum": "ewemk[mPBabZyjhJMNiIqjq`r","previous_checksum": "[GnKo_l[NYM`xs^OcGPOwjQ^^","id": "\\aovSpsbWukcBXrK\\kaILrxdr","environments_id": "lGSL[QWnu]xilrnoi_FlSYNnQ","number": 97}' | http PUT "http://localhost:8080/includedteams/hello world"
func UpdateIncludedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	includedteams := &model.IncludedTeams{}
	if err := readJSON(r, includedteams); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := includedteams.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	includedteams.Prepare()

	if err := includedteams.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	includedteams, _, err = dao.UpdateIncludedTeams(r.Context(),
		argId,
		includedteams)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, includedteams)
}

// DeleteIncludedTeams Delete a single record from included_teams table in the laforge-2 database
// @Summary Delete a record from included_teams
// @Description Delete a single record from included_teams table in the laforge-2 database
// @Tags IncludedTeams
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.IncludedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /includedteams/{argId} [delete]
// http DELETE "http://localhost:8080/includedteams/hello world"
func DeleteIncludedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIncludedTeams(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
