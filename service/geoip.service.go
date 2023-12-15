package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeoIp struct {
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	TimeZone    string  `json:"time_zone"`
	Lat         float32 `json:"latitude"`
	Lon         float32 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

func GetRealIp(c *gin.Context) string {
	ip := c.ClientIP()
	return ip
}

func GetGeoInfo(addr string) (*GeoIp, error) {
	if addr == "" {
		return nil, errors.New("ip is empty")
	}

	geo := &GeoIp{}

	// get ip info from freegeoip.live
	// The 3rd party service may die or be unavailable at any time.
	// For the improvement of performance, we could use the cache to store the geo info
	url := fmt.Sprintf("https://freegeoip.live/json/%s", addr)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return geo, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return geo, err
	}

	err = json.Unmarshal(body, &geo)
	if err != nil {
		fmt.Println(err)
		return geo, err
	}

	return geo, nil
}
