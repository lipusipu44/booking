package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/anilpatro044/router/pkg/config"
	"github.com/anilpatro044/router/pkg/handlers"
	"github.com/anilpatro044/router/pkg/render"
)

const portNumber string=":8080"
var session *scs.SessionManager
var app config.AppConfig
func main(){
  
  tc,err:=render.CreateTemplateCache()
  if err!=nil{
    log.Fatal("Failed to load config from app cache")
  }
  app.TemplateCache=tc
  app.UseCache=false
  app.InProduction=false

  // set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

  app.Session=session
  
  repo:=handlers.NewRepo(&app)
  handlers.NewHandlers(repo)
  render.NewTemplates(&app)
    
  srv:=&http.Server{
    Addr: portNumber,
    Handler: routes(&app),
  }
  err=srv.ListenAndServe()
  if err != nil {
		log.Fatal(err)
	}
  fmt.Println("Main code ends here")
}