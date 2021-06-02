package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/pantaupilkada/helpers/log"
	"gorm.io/gorm"
)

type Context struct {
	DB  *gorm.DB
	Log *log.AppLog
	Gin *gin.Engine
}


func (ctx *Context) Register(group string) {
	public := ctx.Gin.Group(group)
	{
		public.GET("/", ctx.Welcome)
	}
	private := ctx.Gin.Group(group)
	{
		private.POST("/login",  ctx.loginHandler)
		private.GET("/paslon", ctx.getAll)
		private.GET("/paslon/:id", ctx.getPaslon)
		private.POST("/paslon/add", ctx.createPaslon)
		private.PUT("/paslon/update", ctx.updatePaslon)
		private.DELETE("/paslon/delete/:id", ctx.delete)


		private.GET("/adminauth", ctx.getAllAdminAuth)
		private.GET("/adminauth/:id", ctx.getAdminAuth)

		private.PUT("/adminauth/update", ctx.updateAdminAuth)
		private.DELETE("/adminauth/delete/:id", ctx.deleteAdminAuth)

		private.GET("/adminrole", ctx.getAllAdminRole)
		private.GET("/adminrole/:id", ctx.getAdminRole)
		private.POST("/adminrole/add", ctx.createAdminRole)
		private.PUT("/adminrole/update", ctx.updateAdminRole)
		private.DELETE("/adminrole/delete/:id", ctx.deleteAdminRole)

		private.GET("/jorong", ctx.getAllJorong)
		private.GET("/jorong/:id", ctx.getJorong)
		private.POST("/jorong/add", ctx.createJorong)
		private.PUT("/jorong/update", ctx.updateJorong)
		private.DELETE("/jorong/delete/:id", ctx.deleteJorong)

		private.GET("/kecurangan", ctx.getAllKecurangan)
		private.GET("/kecurangan/:id", ctx.getKecurangan)
		private.POST("/kecurangan/add", ctx.createKecurangan)
		private.PUT("/kecurangan/update", ctx.updateKecurangan)
		private.DELETE("/kecurangan/delete/:id", ctx.deleteKecurangan)

		private.GET("/konten", ctx.getAllKonten)
		private.GET("/konten/:id", ctx.getKonten)
		private.POST("/konten/add", ctx.createKonten)
		private.PUT("/konten/update", ctx.updateKonten)
		private.DELETE("/konten/delete/:id", ctx.deleteKonten)

		private.GET("/lokasi", ctx.getAllLokasi)

		private.GET("/pesan", ctx.getAllPesan)
		private.GET("/pesan/:id", ctx.getPesan)
		private.POST("/pesan/add", ctx.createPesan)
		private.PUT("/pesan/update", ctx.updatePesan)
		private.DELETE("/pesan/delete/:id", ctx.deletePesan)

		private.GET("/petugas", ctx.getAllPetugas)
		private.GET("/petugastps", ctx.getAllPetugasTPS)
		private.GET("/petugaskec", ctx.getAllPetugasKec)
		private.GET("/petugaskab", ctx.getAllPetugasKab)
		private.GET("/petugas/:id", ctx.getPetugas)
		private.POST("/petugas/add", ctx.createPetugas)
		private.POST("/sendemail/:tujuan", ctx.sendEmail)
		private.PUT("/petugas/update", ctx.updatePetugas)
		private.DELETE("/petugas/delete/:id", ctx.deletePetugas)
		private.GET("/tps", ctx.getAllTPS)
		private.GET("/tps/:id", ctx.getTPS)
		private.POST("/tps/add", ctx.createTPS)
		private.PUT("/tps/update", ctx.updateTPS)
		private.DELETE("/tps/delete/:id", ctx.deleteTPS)
		private.POST("/api/login", ctx.Loginadmin)
		private.POST("/api/logout", ctx.Logoutadmin)
		private.GET("/api/admin", ctx.Admin)
		private.POST("/api/register", ctx.Registeradmin)
	}
}



