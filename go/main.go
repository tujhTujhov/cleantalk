package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/oschwald/geoip2-golang"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	mc *memcache.Client
)

func main() {
	mc = memcache.New("localhost:11211")
	http.HandleFunc("/", handleGetIpData)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func handleGetIpData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipParam := r.URL.Query().Get("ip")
	ip := net.ParseIP(ipParam)
	if ip == nil {
		io.WriteString(w, "not valid ip")
		return
	}

	ipData := getFromCache(ip.String())
	if ipData != nil {
		io.WriteString(w, string(ipData))
		return
	}

	db, err := geoip2.Open("geoip2_country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}

	response := Response{
		RuName:    record.Country.Names["ru"],
		IsoCode:   record.Country.IsoCode,
		GeoNameId: record.Country.GeoNameID,
	}

	fmt.Printf("got /getIpData request\n")

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	setToCache(ip.String(), jsonResponse)
	io.WriteString(w, string(jsonResponse))
}

func setToCache(key string, data []byte) {
	item := memcache.Item{
		Key:        key,
		Value:      data,
		Flags:      0,
		Expiration: 0,
		CasID:      0,
	}
	err := mc.Set(&item)

	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}
}

func getFromCache(ip string) []byte {
	ipData, err := mc.Get(ip)
	if err != nil {
		return nil
	}
	fmt.Printf("success get from cache\n")

	return ipData.Value
}
