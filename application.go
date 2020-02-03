package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/julienschmidt/httprouter"
)

type Application struct {
	converter    *Converter
	templates    *Templates
	dictionaries *Dictionary
	router       *httprouter.Router
}

func NewApplication() Application {
	r := httprouter.New()

	c := NewConverter()
	t := NewTemplates()
	d := NewDictionary(&c)

	app := Application{&c, &t, &d, r}

	app.router.GET("/", app.Index)
	app.router.GET("/convert/types/:type/", app.EnterValue)
	app.router.GET("/convert/result/:type/", app.Convert)

	return app
}

func (a *Application) start() {
	log.Fatal(http.ListenAndServe(":80", a.router))
}

type IndexPageData struct {
	Types []UnitType
}

type EnterValuePageData struct {
	UnitType string
	UnitMap  map[string][]Unit
}

type ConvertPageData struct {
	ConvertedUnits map[string][]ConvertedUnit
	SourceUit      ConvertedUnit
}

type ConvertedUnit struct {
	Value string
	Key   string
}

func (a *Application) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := IndexPageData{}
	for _, u := range a.dictionaries.unitTypes {
		data.Types = append(data.Types, u)
	}

	a.templates.index.Execute(w, data)
}

func (a *Application) EnterValue(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	data := EnterValuePageData{}
	data.UnitMap = map[string][]Unit{}
	data.UnitType = param.ByName("type")

	for _, u := range a.converter.units[a.dictionaries.getTypeBySlug(data.UnitType)] {
		key := a.dictionaries.unitTypes[u.unitType].Name
		_, ok := data.UnitMap[key]
		if ok {
			data.UnitMap[key] = append(data.UnitMap[key], u)
		} else {
			data.UnitMap[key] = []Unit{u}
		}
	}

	a.templates.enter.Execute(w, data)
}

func (a *Application) Convert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	unit := r.URL.Query()["unit"][0]
	v := r.URL.Query()["value"][0]

	value, _ := strconv.ParseFloat(v, 32)
	data := ConvertPageData{}
	data.SourceUit = ConvertedUnit{v, unit}

	unitType := a.dictionaries.getTypeBySlug(params.ByName("type"))

	data.ConvertedUnits = map[string][]ConvertedUnit{}
	for _, u := range a.converter.units[unitType] {
		if u.Key != unit {
			convertedValue := a.converter.Convert(value, unit, u.Key, unitType)
			cu := ConvertedUnit{humanize.Ftoa(convertedValue), u.Key}
			typeName := a.dictionaries.unitSystems[u.unitSystem].Name

			_, ok := data.ConvertedUnits[typeName]
			if ok {
				data.ConvertedUnits[typeName] = append(data.ConvertedUnits[typeName], cu)
			} else {
				data.ConvertedUnits[typeName] = []ConvertedUnit{cu}
			}
		}
	}

	a.templates.convert.Execute(w, data)
}
