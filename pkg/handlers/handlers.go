package handlers

import (
	"log"
	"net/http"

	"github.com/anilpatro044/router/pkg/config"
	"github.com/anilpatro044/router/pkg/models"
	"github.com/anilpatro044/router/pkg/render"
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