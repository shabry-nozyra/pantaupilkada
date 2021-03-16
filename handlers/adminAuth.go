package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/pantaupilkada/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)



//getAll
func (ctx *Context) getAllAdminAuth(c *gin.Context) {
	p := models.Admins{}
	err := p.All(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, &p)
}


func (ctx *Context) getAdminAuth(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Admin{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (ctx *Context) createAdminAuth(c *gin.Context) {
	p := models.Admin{}
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

func (ctx *Context) deleteAdminAuth(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Admin{}

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

func (ctx *Context) updateAdminAuth(c *gin.Context) {
	cl := models.Admin{}
	err := c.ShouldBindJSON(&cl)
	if err != nil{
		ctx.Log.Warn(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	p := models.Admin{}
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

	p.Name = cl.Name
	p.Email = cl.Email
	p.Password = cl.Password
	p.Email = cl.Email
	p.Image = cl.Image
	p.RoleAdmin = cl.RoleAdmin
	p.IsActive = cl.IsActive
	p.TimeCreated = cl.TimeCreated

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

