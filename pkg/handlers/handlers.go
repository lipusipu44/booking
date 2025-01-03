package handlers

import (
	"log"
	"net/http"

	"github.com/lipusipu44/booking/pkg/config"
	"github.com/lipusipu44/booking/pkg/models"
	"github.com/lipusipu44/booking/pkg/render"
)

var Repo *Repository
var tempData models.TemplateData
type Repository struct{
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App:a,
	}
}

func NewHandlers(repo *Repository){
	Repo=repo
}

func (m *Repository)Home(w http.ResponseWriter,r *http.Request){
	log.Println("Inside Home Handler")
	remoteIp:=r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIp)
	render.TemplateRender(w,"home.page.tmpl",&tempData)
}

func (m *Repository)About(w http.ResponseWriter, r *http.Request){
	stringVal:="Hello World"
	remoteIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	StringMap:=map[string]string{
		"test":stringVal,
		"remote_ip":remoteIP,
	}
	tempData=models.TemplateData{
		StringMap: StringMap,
	}
	

	log.Println("Inside About Handler")
	
	render.TemplateRender(w,"about.page.tmpl",&tempData)
}

func (m *Repository)Reservation(w http.ResponseWriter, r *http.Request){
	render.TemplateRender(w,"make-reservation.page.tmpl",&tempData)
}

func (m *Repository)Generals(w http.ResponseWriter, r *http.Request){
	render.TemplateRender(w,"generals.page.tmpl",&tempData)
}

func (m *Repository)Majors(w http.ResponseWriter, r *http.Request){
	render.TemplateRender(w,"majors.page.tmpl",&tempData)
}

func (m *Repository)Availability(w http.ResponseWriter, r *http.Request){
	render.TemplateRender(w,"search-availability.page.tmpl",&tempData)
}

func (m *Repository)Contact(w http.ResponseWriter,r *http.Request){
	render.TemplateRender(w,"contact.page.tmpl",&tempData)
}

