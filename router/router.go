package router

import (
	"net/http"
	"web/logger"
	v1 "web/router/v1"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	//加载logger
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// // 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "./templates/statics/")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/templates/*.html")

	// 在这里调用中间键即可
	r.Use(CORSMiddleware())

	//路由组data****************************************************
	apiV1:=r.Group("/data")
	{
		iteaq:=apiV1.Group("/iteaq")
		{
			
			iteaq.GET("/airConditioner1",v1.GetAirConditioner1)
			iteaq.GET("/airConditioner2",v1.GetAirConditioner2)
			iteaq.GET("/airConditioner3",v1.GetAirConditioner3)
			iteaq.GET("/airConditioner4",v1.GetAirConditioner4)
			iteaq.GET("/airConditioner5",v1.GetAirConditioner5)
			iteaq.GET("/airConditioner6",v1.GetAirConditioner6)
			iteaq.GET("/jiDianQiZhuZhuang",v1.GetIteaqJiDianQiZhuZhuang)
			iteaq.GET("/lenTongDaoYanGan",v1.GetLenTongDaoYanGan)
			iteaq.GET("/menJinJiFang",v1.GetIteaqMenJinJiFang)
			iteaq.GET("/menJinWeiMoKuai",v1.GetIteaqMenJinWeiMoKuai)
			iteaq.GET("/otherDianChi",v1.GetGetIteaqOtherDianChi)
			iteaq.GET("/peiDianJingMi",v1.GetIteaqPeiDianJingMi)
			iteaq.GET("/peiDianPd1s",v1.GetIteaqPeiDianPd1s)
			iteaq.GET("/peiDianUps1s",v1.GetIteaqPeiDianUps1s)
			iteaq.GET("lenTongDaoWenDu",v1.GetLenTongDaoWenDu)
			iteaq.GET("/ups",v1.GetIteaqUps)
		}	

		gateway:=apiV1.Group("/gateway")
		{
			
			gateway.GET("/cpuUsage",v1.GetGatewayCpuUsage)
			gateway.GET("/memoryUsage",v1.GetGatewayMemoryUsage)
			gateway.GET("/diskUsage",v1.GetGatewayDiskUsage)
			gateway.GET("/session",v1.GetGatewaySession)
		}

		
	}

	//绑定页面***************************************************
	//绑定首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//绑定其他页面
	r.GET("/gateway_cpu_usages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gateway_cpu_usages.html", nil)
	})
	r.GET("/gateway_disk_usages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gateway_disk_usages.html", nil)
	})
	r.GET("/gateway_memory_usages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gateway_memory_usages.html", nil)
	})
	r.GET("/gateway_sessions", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gateway_sessions.html", nil)
	})
	r.GET("/iteaq_air_conditioner1s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_air_conditioner1s.html", nil)
	})
	r.GET("/iteaq_air_conditioner2s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_air_conditioner2s.html", nil)
	})
	r.GET("/iteaq_air_conditioner3s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_air_conditioner3s.html", nil)
	})
	r.GET("/iteaq_air_conditioner4s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_air_conditioner4s.html", nil)
	})
	r.GET("/iteaq_air_conditioner5s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_air_conditioner5s.html", nil)
	})
	r.GET("/iteaq_air_conditioner6s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_air_conditioner6s.html", nil)
	})
	r.GET("/iteaq_ji_dian_qi_zhu_zhuangs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_ji_dian_qi_zhu_zhuangs.html", nil)
	})
	r.GET("/iteaq_leng_tong_dao_yan_gans", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_leng_tong_dao_yan_gans.html", nil)
	})
	r.GET("/iteaq_men_jin_wei_mo_kuais", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_men_jin_wei_mo_kuais.html", nil)
	})
	r.GET("/iteaq_other_dian_chis", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_other_dian_chis.html", nil)
	})
	r.GET("/iteaq_pei_dian_jing_mis", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_pei_dian_jing_mis.html", nil)
	})
	r.GET("/iteaq_pei_dian_pd1s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_pei_dian_pd1s.html", nil)
	})
	r.GET("/iteaq_pei_dian_ups1s", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_pei_dian_ups1s.html", nil)
	})
	r.GET("/iteaq_temps", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_temps.html", nil)
	})
	r.GET("/iteaq_ups", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iteaq_ups.html", nil)
	})

	//绑定测试
	r.GET("/Test1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Test1.html", nil)
	})
	r.GET("/Test2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Test2.html", nil)
	})
	r.GET("/Test3", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Test3.html", nil)
	})
	r.GET("/Test4", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Test4.html", nil)
	})
	


	return r
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://10.161.98.4/")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Control-Allow-Credentials")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
