package controllers

import (
	"encoding/json"
	"github.com/goldalworming/simple_kvdb/models"
	"github.com/goldalworming/simple_kvdb/modules/utils"
	// "errors"
	"strconv"
	// "strings"
	// "github.com/sanity-io/litter"

	"github.com/astaxie/beego"
)

// oprations for Kv
type KvController struct {
	beego.Controller
}
type resultKv struct {
	Success bool
	Rows    interface{}
	Message string
}

// @Title Post
// @Description create Kv
// @Param	body		body 	models.Kv	true		"body for Kv content"
// @Success 201 {int} models.Kv
// @Failure 403 body is empty
// @router / [post]
func (c *KvController) Post() {
	var v models.Kv
	var r resultKv

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
    // keystr := c.Ctx.Input.Param(":key")
    // unixnow := time.Now().Unix()
    // v.T = unixnow

	v.Id, _ = utils.NewId()
	v.IdStr = strconv.FormatInt(v.Id, 10)
	if _, err := models.AddKv(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		r.Success = true
		r.Rows = v
	} else {
		r.Success = false
		r.Message = err.Error()
	}
	c.Data["json"] = r
	c.ServeJSON()
}

// @Title Get
// @Description get Kv by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Kv
// @Failure 403 :id is empty
// @router /:id [get]
func (c *KvController) GetOne() {
	var r resultKv
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetKvById(id)
	if err != nil {
		r.Success = false
		r.Message = err.Error()
	} else {
		r.Rows = v
		r.Success = true
	}
	c.Data["json"] = r
	c.ServeJSON()
}

// @Title Get All
// @Description get Kv
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Kv
// @Failure 403
// @router / [get]
func (c *KvController) GetAll() {
	var limit int64 = 10
	var offset int64 = 0
	res := struct {
		Success bool
		Rows    []interface{}
		Total   int64
		Message string
	}{}

	fields := c.GetString("fields")
	limit, err := c.GetInt64("limit")
	offset, err = c.GetInt64("offset")
	sortby := c.GetString("sortby")
	order := c.GetString("order")
	query := c.GetString("query")


	if offset == 0 {
		total, err := models.CountGetAllKv(query)
		if err != nil {
			res.Message = err.Error()
		}
		res.Total = total
	} else {
		res.Total = 0
	}
	l, err := models.GetAllKv(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
		res.Success = false
		res.Rows = l
		res.Message = err.Error()
	} else {
		res.Success = true
		res.Rows = l
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// @Title Update
// @Description update the Kv
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Kv	true		"body for Kv content"
// @Success 200 {object} models.Kv
// @Failure 403 :id is not int
// @router /:id [put]
func (c *KvController) Put() {
	var r resultKv

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Kv{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err := models.UpdateKvById(&v); err == nil {
		r.Success = true
		r.Message = "OK"
	} else {
		r.Success = false
		r.Message = err.Error()
	}
	c.Data["json"] = r
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Kv
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *KvController) Delete() {
	var r resultKv
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Kv{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.DeleteKv(id); err == nil {
		r.Success = true
		r.Message = "OK"
	} else {
		r.Success = false
		r.Message = err.Error()
	}

	c.Data["json"] = r
	c.ServeJSON()
}
