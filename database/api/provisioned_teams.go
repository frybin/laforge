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

func configProvisionedTeamsRouter(router *httprouter.Router) {
	router.GET("/provisionedteams", GetAllProvisionedTeams)
	router.POST("/provisionedteams", AddProvisionedTeams)

	router.GET("/provisionedteams/:argId", GetProvisionedTeams)
	router.PUT("/provisionedteams/:argId", UpdateProvisionedTeams)
	router.DELETE("/provisionedteams/:argId", DeleteProvisionedTeams)
}

func configGinProvisionedTeamsRouter(router gin.IRoutes) {
	router.GET("/provisionedteams", ConverHttprouterToGin(GetAllProvisionedTeams))
	router.POST("/provisionedteams", ConverHttprouterToGin(AddProvisionedTeams))
	router.GET("/provisionedteams/:argId", ConverHttprouterToGin(GetProvisionedTeams))
	router.PUT("/provisionedteams/:argId", ConverHttprouterToGin(UpdateProvisionedTeams))
	router.DELETE("/provisionedteams/:argId", ConverHttprouterToGin(DeleteProvisionedTeams))
}

// GetAllProvisionedTeams is a function to get a slice of record(s) from provisioned_teams table in the laforge-2 database
// @Summary Get list of ProvisionedTeams
// @Tags ProvisionedTeams
// @Description GetAllProvisionedTeams is a handler to get a slice of record(s) from provisioned_teams table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ProvisionedTeams}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedteams [get]
// http "http://localhost:8080/provisionedteams?page=0&pagesize=20"
func GetAllProvisionedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllProvisionedTeams(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetProvisionedTeams is a function to get a single record from the provisioned_teams table in the laforge-2 database
// @Summary Get record from table ProvisionedTeams by  argId
// @Tags ProvisionedTeams
// @ID argId
// @Description GetProvisionedTeams is a function to get a single record from the provisioned_teams table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.ProvisionedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /provisionedteams/{argId} [get]
// http "http://localhost:8080/provisionedteams/hello world"
func GetProvisionedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetProvisionedTeams(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddProvisionedTeams add to add a single record to provisioned_teams table in the laforge-2 database
// @Summary Add an record to provisioned_teams table
// @Description add to add a single record to provisioned_teams table in the laforge-2 database
// @Tags ProvisionedTeams
// @Accept  json
// @Produce  json
// @Param ProvisionedTeams body model.ProvisionedTeams true "Add ProvisionedTeams"
// @Success 200 {object} model.ProvisionedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedteams [post]
// echo '{"current_checksum": "loSRhnZMHrPWpUyyqZcmiE\\dr","previous_checksum": "bLuMbntr[arOPkhDd\\sHHteFN","id": "cykkey_wnjJtUD]tTpff]yjfn","included_teams_id": "fiwZqtsquCR`ZVtBJAcy[nMVQ","build_configs_id": "^Ty[nJKLLdaSWdYSSuf`oiTpB","state": "PC[EPP`QmsNKLoQe\\qpisYNAJ","planned_checksum": "Pk\\IC[nSUQddDtUDs[UtT_ttX"}' | http POST "http://localhost:8080/provisionedteams"
func AddProvisionedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	provisionedteams := &model.ProvisionedTeams{}

	if err := readJSON(r, provisionedteams); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionedteams.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionedteams.Prepare()

	if err := provisionedteams.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	provisionedteams, _, err = dao.AddProvisionedTeams(r.Context(), provisionedteams)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionedteams)
}

// UpdateProvisionedTeams Update a single record from provisioned_teams table in the laforge-2 database
// @Summary Update an record in table provisioned_teams
// @Description Update a single record from provisioned_teams table in the laforge-2 database
// @Tags ProvisionedTeams
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  ProvisionedTeams body model.ProvisionedTeams true "Update ProvisionedTeams record"
// @Success 200 {object} model.ProvisionedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /provisionedteams/{argId} [patch]
// echo '{"current_checksum": "loSRhnZMHrPWpUyyqZcmiE\\dr","previous_checksum": "bLuMbntr[arOPkhDd\\sHHteFN","id": "cykkey_wnjJtUD]tTpff]yjfn","included_teams_id": "fiwZqtsquCR`ZVtBJAcy[nMVQ","build_configs_id": "^Ty[nJKLLdaSWdYSSuf`oiTpB","state": "PC[EPP`QmsNKLoQe\\qpisYNAJ","planned_checksum": "Pk\\IC[nSUQddDtUDs[UtT_ttX"}' | http PUT "http://localhost:8080/provisionedteams/hello world"
func UpdateProvisionedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	provisionedteams := &model.ProvisionedTeams{}
	if err := readJSON(r, provisionedteams); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := provisionedteams.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	provisionedteams.Prepare()

	if err := provisionedteams.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	provisionedteams, _, err = dao.UpdateProvisionedTeams(r.Context(),
		argId,
		provisionedteams)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, provisionedteams)
}

// DeleteProvisionedTeams Delete a single record from provisioned_teams table in the laforge-2 database
// @Summary Delete a record from provisioned_teams
// @Description Delete a single record from provisioned_teams table in the laforge-2 database
// @Tags ProvisionedTeams
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.ProvisionedTeams
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /provisionedteams/{argId} [delete]
// http DELETE "http://localhost:8080/provisionedteams/hello world"
func DeleteProvisionedTeams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProvisionedTeams(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
