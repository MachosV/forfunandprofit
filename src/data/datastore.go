package data

import (
	"models"
	"net"
	"net/http"
	"time"
)

type Storage struct {
	Visits    uint64
	StartTime time.Time
	IPs       map[string]time.Time
}

func (storage *Storage) IncVisit() {
	storage.Visits = storage.Visits + 1
}

var DataStorage *Storage

func init() {
	go IPGC()
	DataStorage = new(Storage)
	DataStorage.Visits = 0
	DataStorage.StartTime = time.Now()
}

/*
IncrementVisit function
increments the visit variable
after checking for uniqueness
*/
func IncrementVisit() {
	//add checking for uniqueness
	DataStorage.IncVisit()
}

/*
FetchData returns the data to be parsed
*/
func FetchData(w http.ResponseWriter, r *http.Request) *models.IndexData {
	if CheckUnique(r) {
		IncrementVisit()
		expiration := time.Now().Add(12 * time.Hour)
		cookie := http.Cookie{
			Name:    "v",
			Value:   "testvalue",
			Expires: expiration}
		http.SetCookie(w, &cookie)
	}
	var TemplateData = models.IndexData{}
	TemplateData.Visits = DataStorage.Visits
	TemplateData.Uptime = time.Since(DataStorage.StartTime).Round(time.Second)
	return &TemplateData
}

/*
CheckUnique function checks if the request is unique
*/
func CheckUnique(r *http.Request) bool {
	_, err := r.Cookie("v")
	if err == nil {
		return false
	}
	if CheckSeenIP(r) {
		return false
	}
	return true
}

func CheckSeenIP(r *http.Request) bool {
	ip := r.RemoteAddr
	host, _, _ := net.SplitHostPort(ip)
	if _, ok := DataStorage.IPs[host]; !ok {
		return false
	}
	return true
}

func IPGC() {
	for {
		time.Sleep(60 * time.Second)
		for key, _ := range DataStorage.IPs {
			t1 := DataStorage.IPs[key]
			if time.Since(t1).Minutes() > 60 {
				delete(DataStorage.IPs, key)
			}
		}
	}
}
