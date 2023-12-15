package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/service"
)

func GeoIpMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		realIp := service.GetRealIp(c)
		geoip, err := service.GetGeoInfo(realIp)
		if err != nil {
			c.Set("geoip", &service.GeoIp{Ip: realIp}) // set default geoip if the source is not found
		}

		fmt.Println(geoip)
		c.Set("geoip", geoip)
		c.Next()
	}
}
