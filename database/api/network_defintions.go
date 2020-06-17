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

func configNetworkDefintionsRouter(router *httprouter.Router) {
	router.GET("/networkdefintions", GetAllNetworkDefintions)
	router.POST("/networkdefintions", AddNetworkDefintions)

	router.GET("/networkdefintions/:argId", GetNetworkDefintions)
	router.PUT("/networkdefintions/:argId", UpdateNetworkDefintions)
	router.DELETE("/networkdefintions/:argId", DeleteNetworkDefintions)
}

func configGinNetworkDefintionsRouter(router gin.IRoutes) {
	router.GET("/networkdefintions", ConverHttprouterToGin(GetAllNetworkDefintions))
	router.POST("/networkdefintions", ConverHttprouterToGin(AddNetworkDefintions))
	router.GET("/networkdefintions/:argId", ConverHttprouterToGin(GetNetworkDefintions))
	router.PUT("/networkdefintions/:argId", ConverHttprouterToGin(UpdateNetworkDefintions))
	router.DELETE("/networkdefintions/:argId", ConverHttprouterToGin(DeleteNetworkDefintions))
}

// GetAllNetworkDefintions is a function to get a slice of record(s) from network_defintions table in the laforge-2 database
// @Summary Get list of NetworkDefintions
// @Tags NetworkDefintions
// @Description GetAllNetworkDefintions is a handler to get a slice of record(s) from network_defintions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.NetworkDefintions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /networkdefintions [get]
// http "http://localhost:8080/networkdefintions?page=0&pagesize=20"
func GetAllNetworkDefintions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllNetworkDefintions(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetNetworkDefintions is a function to get a single record from the network_defintions table in the laforge-2 database
// @Summary Get record from table NetworkDefintions by  argId
// @Tags NetworkDefintions
// @ID argId
// @Description GetNetworkDefintions is a function to get a single record from the network_defintions table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.NetworkDefintions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /networkdefintions/{argId} [get]
// http "http://localhost:8080/networkdefintions/hello world"
func GetNetworkDefintions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetNetworkDefintions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddNetworkDefintions add to add a single record to network_defintions table in the laforge-2 database
// @Summary Add an record to network_defintions table
// @Description add to add a single record to network_defintions table in the laforge-2 database
// @Tags NetworkDefintions
// @Accept  json
// @Produce  json
// @Param NetworkDefintions body model.NetworkDefintions true "Add NetworkDefintions"
// @Success 200 {object} model.NetworkDefintions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /networkdefintions [post]
// echo '{"id": "bwROlhoSIem[FU^T_[rFoLRKN","domain": "o^gpJxcZ[OIJv[lUinSXaysxc"}' | http POST "http://localhost:8080/networkdefintions"
func AddNetworkDefintions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	networkdefintions := &model.NetworkDefintions{}

	if err := readJSON(r, networkdefintions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := networkdefintions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	networkdefintions.Prepare()

	if err := networkdefintions.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	networkdefintions, _, err = dao.AddNetworkDefintions(r.Context(), networkdefintions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, networkdefintions)
}

// UpdateNetworkDefintions Update a single record from network_defintions table in the laforge-2 database
// @Summary Update an record in table network_defintions
// @Description Update a single record from network_defintions table in the laforge-2 database
// @Tags NetworkDefintions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  NetworkDefintions body model.NetworkDefintions true "Update NetworkDefintions record"
// @Success 200 {object} model.NetworkDefintions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /networkdefintions/{argId} [patch]
// echo '{"id": "bwROlhoSIem[FU^T_[rFoLRKN","domain": "o^gpJxcZ[OIJv[lUinSXaysxc"}' | http PUT "http://localhost:8080/networkdefintions/hello world"
func UpdateNetworkDefintions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	networkdefintions := &model.NetworkDefintions{}
	if err := readJSON(r, networkdefintions); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := networkdefintions.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	networkdefintions.Prepare()

	if err := networkdefintions.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	networkdefintions, _, err = dao.UpdateNetworkDefintions(r.Context(),
		argId,
		networkdefintions)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, networkdefintions)
}

// DeleteNetworkDefintions Delete a single record from network_defintions table in the laforge-2 database
// @Summary Delete a record from network_defintions
// @Description Delete a single record from network_defintions table in the laforge-2 database
// @Tags NetworkDefintions
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.NetworkDefintions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /networkdefintions/{argId} [delete]
// http DELETE "http://localhost:8080/networkdefintions/hello world"
func DeleteNetworkDefintions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteNetworkDefintions(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
