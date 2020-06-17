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

func configUsersRouter(router *httprouter.Router) {
	router.GET("/users", GetAllUsers)
	router.POST("/users", AddUsers)

	router.GET("/users/:argId", GetUsers)
	router.PUT("/users/:argId", UpdateUsers)
	router.DELETE("/users/:argId", DeleteUsers)
}

func configGinUsersRouter(router gin.IRoutes) {
	router.GET("/users", ConverHttprouterToGin(GetAllUsers))
	router.POST("/users", ConverHttprouterToGin(AddUsers))
	router.GET("/users/:argId", ConverHttprouterToGin(GetUsers))
	router.PUT("/users/:argId", ConverHttprouterToGin(UpdateUsers))
	router.DELETE("/users/:argId", ConverHttprouterToGin(DeleteUsers))
}

// GetAllUsers is a function to get a slice of record(s) from users table in the laforge-2 database
// @Summary Get list of Users
// @Tags Users
// @Description GetAllUsers is a handler to get a slice of record(s) from users table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Users}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users [get]
// http "http://localhost:8080/users?page=0&pagesize=20"
func GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	records, totalRows, err := dao.GetAllUsers(r.Context(), page, pagesize, order)
	if err != nil {
		returnError(w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(w, result)
}

// GetUsers is a function to get a single record from the users table in the laforge-2 database
// @Summary Get record from table Users by  argId
// @Tags Users
// @ID argId
// @Description GetUsers is a function to get a single record from the users table in the laforge-2 database
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 200 {object} model.Users
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /users/{argId} [get]
// http "http://localhost:8080/users/hello world"
func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	record, err := dao.GetUsers(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, record)
}

// AddUsers add to add a single record to users table in the laforge-2 database
// @Summary Add an record to users table
// @Description add to add a single record to users table in the laforge-2 database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Users body model.Users true "Add Users"
// @Success 200 {object} model.Users
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users [post]
// echo '{"email": "`S`ksUTaaSwEgR^ZwQJdUj]_^","roles_id": "eJM]bMs_CqdSrbpiQ^chAjsOE","id": "E\\^Cig^lGdbHlYUUUEibVxX\\C"}' | http POST "http://localhost:8080/users"
func AddUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users := &model.Users{}

	if err := readJSON(r, users); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := users.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	users.Prepare()

	if err := users.Validate(model.Create); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	var err error
	users, _, err = dao.AddUsers(r.Context(), users)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, users)
}

// UpdateUsers Update a single record from users table in the laforge-2 database
// @Summary Update an record in table users
// @Description Update a single record from users table in the laforge-2 database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Param  Users body model.Users true "Update Users record"
// @Success 200 {object} model.Users
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users/{argId} [patch]
// echo '{"email": "`S`ksUTaaSwEgR^ZwQJdUj]_^","roles_id": "eJM]bMs_CqdSrbpiQ^chAjsOE","id": "E\\^Cig^lGdbHlYUUUEibVxX\\C"}' | http PUT "http://localhost:8080/users/hello world"
func UpdateUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	users := &model.Users{}
	if err := readJSON(r, users); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	if err := users.BeforeSave(); err != nil {
		returnError(w, r, dao.ErrBadParams)
	}

	users.Prepare()

	if err := users.Validate(model.Update); err != nil {
		returnError(w, r, dao.ErrBadParams)
		return
	}

	users, _, err = dao.UpdateUsers(r.Context(),
		argId,
		users)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeJSON(w, users)
}

// DeleteUsers Delete a single record from users table in the laforge-2 database
// @Summary Delete a record from users
// @Description Delete a single record from users table in the laforge-2 database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  argId path string true "id"
// @Success 204 {object} model.Users
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /users/{argId} [delete]
// http DELETE "http://localhost:8080/users/hello world"
func DeleteUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	argId, err := parseString(ps, "argId")
	if err != nil {
		returnError(w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUsers(r.Context(), argId)
	if err != nil {
		returnError(w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
