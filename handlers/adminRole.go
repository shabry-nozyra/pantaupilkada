package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/pantaupilkada/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)



//getAll
func (ctx *Context) getAllAdminRole(c *gin.Context) {
	p := models.AdminRoles{}
	err := p.All(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, &p)
}


func (ctx *Context) getAdminRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.AdminRole{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (ctx *Context) createAdminRole(c *gin.Context) {
	p := models.AdminRole{}
	err := c.ShouldBindJSON(&p)

	err = p.Create(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusBadRequest, errorm)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	c.JSON(http.StatusCreated, res)

}

func (ctx *Context) deleteAdminRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.AdminRole{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Warn("Bad Request")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = p.Delete(ctx.DB, id)
	if err != nil{
		errorm := "Gagal Menjalankan Query"
		ctx.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, errorm)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	c.JSON(http.StatusOK, res)
}

func (ctx *Context) updateAdminRole(c *gin.Context) {
	cl := models.AdminRole{}
	err := c.ShouldBindJSON(&cl)
	if err != nil{
		ctx.Log.Warn(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	p := models.AdminRole{}
	err = p.Get(ctx.DB, int(cl.ID))



	if err != nil{
		if err == gorm.ErrRecordNotFound{
			ctx.Log.Warn(err)
			c.JSON(http.StatusNotFound, nil)
			return
		}else{
			ctx.Log.Warn(err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	p.RoleAdmin = cl.RoleAdmin


	err = p.Update(ctx.DB)
	if err != nil{
		errorm := "Gagal Menjalankan Query"
		ctx.Log.Error(err)
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	c.JSON(http.StatusCreated, res)
}

