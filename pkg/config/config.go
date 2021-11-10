package config

import (
	"fmt"
	"html/template"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

var once sync.Once
var App *AppConfig

// AppConfig holds the application config
type AppConfig struct {
	UseCache 		bool
	TemplateCache 	map[string]*template.Template
	PortNumber 		string
	InfoLog 		*log.Logger
	InProduction	bool
	Session			*scs.SessionManager
}

func GetConfig() *AppConfig {
	once.Do(func() {
		
		App = &AppConfig{
			UseCache: false,
			PortNumber: ":8080",
			InProduction: true,
		}
	})
	return App
}
func PrintAppConfig(){
	fmt.Println(App)
}