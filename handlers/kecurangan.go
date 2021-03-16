package handlers

import (
"github.com/gin-gonic/gin"
"github.com/shabry-nozyra/pantaupilkada/models"
"gorm.io/gorm"
"net/http"
"strconv"
)



//getAll
func (ctx *Context) getAllKecurangan(c *gin.Context) {
	p := models.Kecurangans{}
	err := p.All(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, &p)
}


func (ctx *Context) getKecurangan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Kecurangan{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (ctx *Context) createKecurangan(c *gin.Context) {
	p := models.Kecurangan{}
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

func (ctx *Context) deleteKecurangan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Kecurangan{}

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

func (ctx *Context) updateKecurangan(c *gin.Context) {
	cl := models.Kecurangan{}
	err := c.ShouldBindJSON(&cl)
	if err != nil{
		ctx.Log.Warn(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	p := models.Kecurangan{}
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

	p.Detail = cl.Detail
	p.WaktuKejadian = cl.WaktuKejadian
	p.Jam = cl.Jam
	p.Tempat = cl.Tempat
	p.Bukti1 = cl.Bukti1
	p.Bukti2 = cl.Bukti2
	p.Bukti3 = cl.Bukti3

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


