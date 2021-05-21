package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/pantaupilkada/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

const SecretKey = "secret"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//getAll
func (ctx *Context) getAllAdminAuth(c *gin.Context) {
	p := models.Admins{}
	err := p.All(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
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



func (ctx *Context) Registeradmin(c *gin.Context) {
	p := models.Admin{}
	err := c.ShouldBindJSON(&p)

	//err = p.Create(ctx.DB)
	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	//
	password, _ := HashPassword(p.Password)
	//
	a := models.Admin{
		ID: models.MaxId(ctx.DB)+1,
		Name: p.Name,
		Email: p.Email,
		Password: password,
		Image: "default.jpg",
		RoleAdmin: 1,
		IsActive: 1,
		TimeCreated: 0,
	}
	//c.JSON(http.StatusOK, a)
	//return
	err = a.Create(ctx.DB)
	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, a)
}

func (ctx *Context) Loginadmin(c *gin.Context){
	p := models.Admin{}
	err := c.ShouldBindJSON(&p)

	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	admin := models.Admin{}

	err = admin.GetByEmail(ctx.DB, p.Email)

	if admin.ID == 0{
		c.JSON(http.StatusNotFound, "user not found")
		return
	}

	//cek password
	match := CheckPasswordHash(p.Password, admin.Password)

	if (match != true){
		c.JSON(http.StatusBadRequest,"password didnt match")
		return
	}


	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(admin.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError,"could not login")
		return
	}

	c.SetCookie("jwt", token, 60*60*24, "/", "pantaufront.azurewebsites.net", false, false)
	c.JSON(http.StatusOK, "success")
}

func (ctx *Context) Admin(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		res := map[string]string{
			"status": "unauthenticated",
		}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var admin models.Admin

	id_admin, _ := strconv.Atoi(claims.Issuer)

	err = admin.Get(ctx.DB, id_admin)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}

	c.JSON(http.StatusOK, admin)

}

func (ctx *Context) Logoutadmin(c *gin.Context){
	c.SetCookie("jwt", "", -(60*60*24), "", "", false, true)
	c.JSON(http.StatusOK, "success")
}
