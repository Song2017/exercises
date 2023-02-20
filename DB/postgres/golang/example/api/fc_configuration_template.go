package api

import (
	"net/http"

	"example.com/rest/example/dao"
	"example.com/rest/example/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configFcConfigurationTemplateRouter(router *httprouter.Router) {
	router.GET("/fcconfigurationtemplate", GetAllFcConfigurationTemplate)
	router.POST("/fcconfigurationtemplate", AddFcConfigurationTemplate)
	router.GET("/fcconfigurationtemplate/:argFcConfigurationTemplateID", GetFcConfigurationTemplate)
	router.PUT("/fcconfigurationtemplate/:argFcConfigurationTemplateID", UpdateFcConfigurationTemplate)
	router.DELETE("/fcconfigurationtemplate/:argFcConfigurationTemplateID", DeleteFcConfigurationTemplate)
}

func configGinFcConfigurationTemplateRouter(router gin.IRoutes) {
	router.GET("/fcconfigurationtemplate", ConverHttprouterToGin(GetAllFcConfigurationTemplate))
	router.POST("/fcconfigurationtemplate", ConverHttprouterToGin(AddFcConfigurationTemplate))
	router.GET("/fcconfigurationtemplate/:argFcConfigurationTemplateID", ConverHttprouterToGin(GetFcConfigurationTemplate))
	router.PUT("/fcconfigurationtemplate/:argFcConfigurationTemplateID", ConverHttprouterToGin(UpdateFcConfigurationTemplate))
	router.DELETE("/fcconfigurationtemplate/:argFcConfigurationTemplateID", ConverHttprouterToGin(DeleteFcConfigurationTemplate))
}

// GetAllFcConfigurationTemplate is a function to get a slice of record(s) from fc_configuration_template table in the main database
// @Summary Get list of FcConfigurationTemplate
// @Tags FcConfigurationTemplate
// @Description GetAllFcConfigurationTemplate is a handler to get a slice of record(s) from fc_configuration_template table in the main database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FcConfigurationTemplate}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fcconfigurationtemplate [get]
// http "http://localhost:8080/fcconfigurationtemplate?page=0&pagesize=20" X-Api-User:user123
func GetAllFcConfigurationTemplate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "fc_configuration_template", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFcConfigurationTemplate(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetFcConfigurationTemplate is a function to get a single record from the fc_configuration_template table in the main database
// @Summary Get record from table FcConfigurationTemplate by  argFcConfigurationTemplateID
// @Tags FcConfigurationTemplate
// @ID argFcConfigurationTemplateID
// @Description GetFcConfigurationTemplate is a function to get a single record from the fc_configuration_template table in the main database
// @Accept  json
// @Produce  json
// @Param  argFcConfigurationTemplateID path int true "fc_configuration_template_id"
// @Success 200 {object} model.FcConfigurationTemplate
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /fcconfigurationtemplate/{argFcConfigurationTemplateID} [get]
// http "http://localhost:8080/fcconfigurationtemplate/1" X-Api-User:user123
func GetFcConfigurationTemplate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argFcConfigurationTemplateID, err := parseInt32(ps, "argFcConfigurationTemplateID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fc_configuration_template", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFcConfigurationTemplate(ctx, argFcConfigurationTemplateID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFcConfigurationTemplate add to add a single record to fc_configuration_template table in the main database
// @Summary Add an record to fc_configuration_template table
// @Description add to add a single record to fc_configuration_template table in the main database
// @Tags FcConfigurationTemplate
// @Accept  json
// @Produce  json
// @Param FcConfigurationTemplate body model.FcConfigurationTemplate true "Add FcConfigurationTemplate"
// @Success 200 {object} model.FcConfigurationTemplate
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fcconfigurationtemplate [post]
// echo '{"fc_configuration_template_id": 48,"config_level": "oZDcUDCoUUcyylPKAtwXNFnRX","setting": "jsooYqRhFgWJybZURBpRMVeer","setting_secret": "meQTamLtNVSubfpsQrYjOhOEx","is_enabled": true,"create_time": "2246-05-21T11:55:29.603113421+08:00","modify_time": "2103-02-15T02:31:24.512196174+08:00","create_by": "tZOiRtFgNhbKKYOgruQwvYwfr","modify_by": "HNpPaugsUWvgNMWRHfqoIKYBk"}' | http POST "http://localhost:8080/fcconfigurationtemplate" X-Api-User:user123
func AddFcConfigurationTemplate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	fcconfigurationtemplate := &model.FcConfigurationTemplate{}

	if err := readJSON(r, fcconfigurationtemplate); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := fcconfigurationtemplate.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	fcconfigurationtemplate.Prepare()

	if err := fcconfigurationtemplate.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fc_configuration_template", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	fcconfigurationtemplate, _, err = dao.AddFcConfigurationTemplate(ctx, fcconfigurationtemplate)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, fcconfigurationtemplate)
}

// UpdateFcConfigurationTemplate Update a single record from fc_configuration_template table in the main database
// @Summary Update an record in table fc_configuration_template
// @Description Update a single record from fc_configuration_template table in the main database
// @Tags FcConfigurationTemplate
// @Accept  json
// @Produce  json
// @Param  argFcConfigurationTemplateID path int true "fc_configuration_template_id"
// @Param  FcConfigurationTemplate body model.FcConfigurationTemplate true "Update FcConfigurationTemplate record"
// @Success 200 {object} model.FcConfigurationTemplate
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fcconfigurationtemplate/{argFcConfigurationTemplateID} [put]
// echo '{"fc_configuration_template_id": 48,"config_level": "oZDcUDCoUUcyylPKAtwXNFnRX","setting": "jsooYqRhFgWJybZURBpRMVeer","setting_secret": "meQTamLtNVSubfpsQrYjOhOEx","is_enabled": true,"create_time": "2246-05-21T11:55:29.603113421+08:00","modify_time": "2103-02-15T02:31:24.512196174+08:00","create_by": "tZOiRtFgNhbKKYOgruQwvYwfr","modify_by": "HNpPaugsUWvgNMWRHfqoIKYBk"}' | http PUT "http://localhost:8080/fcconfigurationtemplate/1"  X-Api-User:user123
func UpdateFcConfigurationTemplate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argFcConfigurationTemplateID, err := parseInt32(ps, "argFcConfigurationTemplateID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	fcconfigurationtemplate := &model.FcConfigurationTemplate{}
	if err := readJSON(r, fcconfigurationtemplate); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := fcconfigurationtemplate.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	fcconfigurationtemplate.Prepare()

	if err := fcconfigurationtemplate.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fc_configuration_template", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	fcconfigurationtemplate, _, err = dao.UpdateFcConfigurationTemplate(ctx,
		argFcConfigurationTemplateID,
		fcconfigurationtemplate)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, fcconfigurationtemplate)
}

// DeleteFcConfigurationTemplate Delete a single record from fc_configuration_template table in the main database
// @Summary Delete a record from fc_configuration_template
// @Description Delete a single record from fc_configuration_template table in the main database
// @Tags FcConfigurationTemplate
// @Accept  json
// @Produce  json
// @Param  argFcConfigurationTemplateID path int true "fc_configuration_template_id"
// @Success 204 {object} model.FcConfigurationTemplate
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /fcconfigurationtemplate/{argFcConfigurationTemplateID} [delete]
// http DELETE "http://localhost:8080/fcconfigurationtemplate/1" X-Api-User:user123
func DeleteFcConfigurationTemplate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argFcConfigurationTemplateID, err := parseInt32(ps, "argFcConfigurationTemplateID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fc_configuration_template", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFcConfigurationTemplate(ctx, argFcConfigurationTemplateID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
