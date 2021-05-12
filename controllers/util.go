package controllers

import (
	"log"
	"net"

	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
)

func getOffset(pageIndex int, pageSize int) int {
	return (pageIndex - 1) * pageSize
}

func getIPLoc(addr string) string {
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer func(db *geoip2.DBReader) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
	record, _ := db.City(net.ParseIP(addr))
	return record.Country.Names["zh-CN"]
}
