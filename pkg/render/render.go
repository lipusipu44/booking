package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/lipusipu44/booking/pkg/config"
	"github.com/lipusipu44/booking/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app=a
}

func AddDefaultTemplateData(t *models.TemplateData) *models.TemplateData{
	return t
}

func TemplateRender(w http.ResponseWriter, s string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache{
		tc=app.TemplateCache
	}else{
		tc,_=CreateTemplateCache()
	}
	t,ok:=tc[s]
	if !ok{
		log.Fatal()
	}
	buf:=new(bytes.Buffer)
	td=AddDefaultTemplateData(td) //as of now not doing anything here
	err:=t.Execute(buf,td)
	if err!=nil{
		log.Fatal()
	}
	_,err=buf.WriteTo(w)
	if err!=nil{
		log.Fatal()
	}
}

func CreateTemplateCache()(map[string]*template.Template,error){
	var mp = map[string]*template.Template{}
	pages,err:=filepath.Glob("./templates/*.page.tmpl")
	if err!=nil{
		log.Fatal()
		return mp,err
	}

	for _,page:=range pages{
		filename:=filepath.Base(page)
		ts,err:=template.New(filename).ParseFiles(page)
		if err!=nil{
			log.Fatal()
			return mp,err
		}

		layout,err:=filepath.Glob("./templates/*.layout.tmpl")
		if err!=nil{
			log.Fatal()
			return mp,err
		}

		if len(layout)>0{
			ts,err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err!=nil{
				log.Fatal()
				return mp,err
			}
		}

		mp[filename]=ts
	}
	return mp,nil
}

