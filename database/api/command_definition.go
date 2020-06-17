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

func configCommandDefinitionRouter(router *httprouter.Router) {
	router.GET("/commanddefinition", GetAllCommandDefinition)
	router.POST("/commanddefinition", AddCommandDefinition)

	router.GET("/commanddefinition/:argId", GetCommandDefinition)
	router.PUT("/commanddefinition/:argId", UpdateCommandDefinition)
	router.DELETE("/commanddefinition/:argId", DeleteCommandDefinition)
}

func configGinCommandDefinitionRouter(router gin.IRoutes) {
	router.GET("/commanddefinition", ConverHttprouterToGin(GetAllCommandDefinition))
	router.POST("/commanddefinition", ConverHttprouterToGin(AddCommandDefinition))
	router.GET("/commanddefinition/:argId", ConverHttprouterToGin(GetCommandDefinition))
	router.PUT("/commanddefinition/:argId", ConverHttprouterToGin(UpdateCommandDefinition))
	router.DELETE("/commanddefinition/:argId", ConverHttprouterToGin(DeleteCommandDefinition))
}

// GetAllCommandDefinition is a function to get a slice of record(s) from command_definition table in the laforge-2 database
// @Summary Get list of CommandDefinition
// @Tags CommandDefinition
// @Description GetAllCommandDefinition is a handler to get a slice of record(s) from command_definition table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.CommandDefinition}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /commanddefinition [get]
// http "http://localhost:8080/commanddefinition?page=0&pagesize=20"
func GetAllCommandDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllCommandDefinition(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetCommandDefinition is a function to get a single record from the command_definition table in the laforge-2 database
// @Summary Get record from table CommandDefinition by  argId
// @Tags CommandDefinition
// @ID argId
// @Description GetCommandDefinition is a function to get a single record from the command_definition table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.CommandDefinition
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /commanddefinition/{argId} [get]
// http "http://localhost:8080/commanddefinition/hello world"
func GetCommandDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetCommandDefinition(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddCommandDefinition add to add a single record to command_definition table in the laforge-2 database
// @Summary Add an record to command_definition table
// @Description add to add a single record to command_definition table in the laforge-2 database
// @Tags CommandDefinition
// @Accept  json
// @Produce  json
// @Param CommandDefinition body model.CommandDefinition true "Add CommandDefinition"
// @Success 200 {object} model.CommandDefinition
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /commanddefinition [post]
// echo '{"id": "mFTmm[sfngTUbNIy_sQCtZsVM","type": "ewU[ktuSad_[E]mKrraARilKs","name": "hUucibrRWslGsIC\\OFIugixGu"}' | http POST "http://localhost:8080/commanddefinition"
func AddCommandDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	commanddefinition := &model.CommandDefinition{}

	if err := readJSON(r, commanddefinition); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := commanddefinition.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	commanddefinition.Prepare()

	if err := commanddefinition.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	commanddefinition, _, err = dao.AddCommandDefinition(r.Context(), commanddefinition)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, commanddefinition)
}

// UpdateCommandDefinition Update a single record from command_definition table in the laforge-2 database
// @Summary Update an record in table command_definition
// @Description Update a single record from command_definition table in the laforge-2 database
// @Tags CommandDefinition
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  CommandDefinition body model.CommandDefinition true "Update CommandDefinition record"
// @Success 200 {object} model.CommandDefinition
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /commanddefinition/{argId} [patch]
// echo '{"id": "mFTmm[sfngTUbNIy_sQCtZsVM","type": "ewU[ktuSad_[E]mKrraARilKs","name": "hUucibrRWslGsIC\\OFIugixGu"}' | http PUT "http://localhost:8080/commanddefinition/hello world"
func UpdateCommandDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	commanddefinition := &model.CommandDefinition{}
	if err := readJSON(r, commanddefinition); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := commanddefinition.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	commanddefinition.Prepare()

	if err := commanddefinition.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	commanddefinition, _, err = dao.UpdateCommandDefinition(r.Context(),
		argId,
		commanddefinition)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, commanddefinition)
}

// DeleteCommandDefinition Delete a single record from command_definition table in the laforge-2 database
// @Summary Delete a record from command_definition
// @Description Delete a single record from command_definition table in the laforge-2 database
// @Tags CommandDefinition
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.CommandDefinition
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /commanddefinition/{argId} [delete]
// http DELETE "http://localhost:8080/commanddefinition/hello world"
func DeleteCommandDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCommandDefinition(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
