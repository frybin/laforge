package api

import (
	"encoding/json"
	_ "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/frybin/laforge/database/dao"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = time.Second
)

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ConfigRouter configure http.Handler router
func ConfigRouter() http.Handler {
	router := httprouter.New()
	configBuildConfigsRouter(router)
	configBuildsRouter(router)
	configCommandDefinitionRouter(router)
	configCompetitionUsersRouter(router)
	configCompetitionsRouter(router)
	configDNSRecordDefinitionsRouter(router)
	configEnvironmentsRouter(router)
	configHostDefinitionsRouter(router)
	configIdentityDefinitionsRouter(router)
	configIncludedHostsRouter(router)
	configIncludedNetworksRouter(router)
	configIncludedStepsRouter(router)
	configIncludedTeamsRouter(router)
	configNetworkDefintionsRouter(router)
	configProvisionedHostsRouter(router)
	configProvisionedNetworksRouter(router)
	configProvisionedStepsRouter(router)
	configProvisionedTeamsRouter(router)
	configRemoteFileDefinitionsRouter(router)
	configRolesRouter(router)
	configScriptDefinitionsRouter(router)
	configStepDefinitionsRouter(router)
	configUsersRouter(router)

	return router
}

// ConfigGinRouter configure gin router
func ConfigGinRouter(router gin.IRoutes) {
	configGinBuildConfigsRouter(router)
	configGinBuildsRouter(router)
	configGinCommandDefinitionRouter(router)
	configGinCompetitionUsersRouter(router)
	configGinCompetitionsRouter(router)
	configGinDNSRecordDefinitionsRouter(router)
	configGinEnvironmentsRouter(router)
	configGinHostDefinitionsRouter(router)
	configGinIdentityDefinitionsRouter(router)
	configGinIncludedHostsRouter(router)
	configGinIncludedNetworksRouter(router)
	configGinIncludedStepsRouter(router)
	configGinIncludedTeamsRouter(router)
	configGinNetworkDefintionsRouter(router)
	configGinProvisionedHostsRouter(router)
	configGinProvisionedNetworksRouter(router)
	configGinProvisionedStepsRouter(router)
	configGinProvisionedTeamsRouter(router)
	configGinRemoteFileDefinitionsRouter(router)
	configGinRolesRouter(router)
	configGinScriptDefinitionsRouter(router)
	configGinStepDefinitionsRouter(router)
	configGinUsersRouter(router)

	return
}

// ConverHttprouterToGin wrap httprouter.Handle to gin.HandlerFunc
func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func writeRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func returnError(w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJSON(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func parseUint8(ps httprouter.Params, key string) (uint8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}
func parseUint16(ps httprouter.Params, key string) (uint16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}
func parseUint32(ps httprouter.Params, key string) (uint32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}
func parseUint64(ps httprouter.Params, key string) (uint64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}
func parseInt(ps httprouter.Params, key string) (int, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}
func parseInt8(ps httprouter.Params, key string) (int8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}
func parseInt16(ps httprouter.Params, key string) (int16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}
func parseInt32(ps httprouter.Params, key string) (int32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}
func parseInt64(ps httprouter.Params, key string) (int64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err
}
func parseString(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
func parseUUID(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
