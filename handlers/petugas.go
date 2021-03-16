package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/pantaupilkada/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)



//getAll
func (ctx *Context) getAllPetugas(c *gin.Context) {
	p := models.Petugass{}
	err := p.All(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, &p)
}


func (ctx *Context) getPetugas(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Petugas{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (ctx *Context) createPetugas(c *gin.Context) {
	p := models.Petugas{}
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

func (ctx *Context) deletePetugas(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Petugas{}

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

func (ctx *Context) updatePetugas(c *gin.Context) {
	cl := models.Petugas{}
	err := c.ShouldBindJSON(&cl)
	if err != nil{
		ctx.Log.Warn(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	p := models.Petugas{}
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

	p.NoTPS = cl.NoTPS
	p.Kecamatan = cl.Kecamatan
	p.Nagari = cl.Nagari
	p.Jorong = cl.Jorong
	p.Foto = cl.Foto
	p.NoHp = cl.NoHp
	p.Email = cl.Email
	p.Username = cl.Username
	p.NamaLengkap = cl.NamaLengkap
	p.Password = cl.Password
	p.RolePetugas = cl.RolePetugas
	p.IsActive = cl.IsActive
	p.TimeKirim = cl.TimeKirim

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

