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

		public.POST("/login",  ctx.loginHandler)
		public.GET("/paslon", ctx.getAll)
		public.GET("/paslon/:id", ctx.getPaslon)
		public.POST("/paslon/add", ctx.createPaslon)
		public.PUT("/paslon/update", ctx.updatePaslon)
		public.DELETE("/paslon/delete/:id", ctx.delete)


		public.GET("/adminauth", ctx.getAllAdminAuth)
		public.GET("/adminauth/:id", ctx.getAdminAuth)
		public.POST("/api/register", ctx.Registeradmin)
		public.POST("/api/login", ctx.Loginadmin)
		public.POST("/api/logout", ctx.Logoutadmin)
		public.GET("/api/admin", ctx.Admin)

		public.PUT("/adminauth/update", ctx.updateAdminAuth)
		public.DELETE("/adminauth/delete/:id", ctx.deleteAdminAuth)

		public.GET("/adminrole", ctx.getAllAdminRole)
		public.GET("/adminrole/:id", ctx.getAdminRole)
		public.POST("/adminrole/add", ctx.createAdminRole)
		public.PUT("/adminrole/update", ctx.updateAdminRole)
		public.DELETE("/adminrole/delete/:id", ctx.deleteAdminRole)

		public.GET("/jorong", ctx.getAllJorong)
		public.GET("/jorong/:id", ctx.getJorong)
		public.POST("/jorong/add", ctx.createJorong)
		public.PUT("/jorong/update", ctx.updateJorong)
		public.DELETE("/jorong/delete/:id", ctx.deleteJorong)

		public.GET("/kecurangan", ctx.getAllKecurangan)
		public.GET("/kecurangan/:id", ctx.getKecurangan)
		public.POST("/kecurangan/add", ctx.createKecurangan)
		public.PUT("/kecurangan/update", ctx.updateKecurangan)
		public.DELETE("/kecurangan/delete/:id", ctx.deleteKecurangan)

		public.GET("/konten", ctx.getAllKonten)
		public.GET("/konten/:id", ctx.getKonten)
		public.POST("/konten/add", ctx.createKonten)
		public.PUT("/konten/update", ctx.updateKonten)
		public.DELETE("/konten/delete/:id", ctx.deleteKonten)

		public.GET("/lokasi", ctx.getAllLokasi)

		public.GET("/pesan", ctx.getAllPesan)
		public.GET("/pesan/:id", ctx.getPesan)
		public.POST("/pesan/add", ctx.createPesan)
		public.PUT("/pesan/update", ctx.updatePesan)
		public.DELETE("/pesan/delete/:id", ctx.deletePesan)

		public.GET("/petugas", ctx.getAllPetugas)
		public.GET("/petugastps", ctx.getAllPetugasTPS)
		public.GET("/petugaskec", ctx.getAllPetugasKec)
		public.GET("/petugaskab", ctx.getAllPetugasKab)
		public.GET("/petugas/:id", ctx.getPetugas)
		public.POST("/petugas/add", ctx.createPetugas)
		public.POST("/sendemail/:tujuan", ctx.sendEmail)
		public.PUT("/petugas/update", ctx.updatePetugas)
		public.DELETE("/petugas/delete/:id", ctx.deletePetugas)
		public.GET("/tps", ctx.getAllTPS)
		public.GET("/tps/:id", ctx.getTPS)
		public.POST("/tps/add", ctx.createTPS)
		public.PUT("/tps/update", ctx.updateTPS)
		public.DELETE("/tps/delete/:id", ctx.deleteTPS)
	}
}



