package categorycontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/mhdianrush/go-crud-web/entities"
	"github.com/mhdianrush/go-crud-web/models/categorymodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		// if failed
		ok := categorymodel.Create(category)
		if !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}
		// if success
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}
		temp.Execute(w, data)
	}

	if r.Method == http.MethodPost {
		var category entities.Category

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		ok := categorymodel.Update(id, category)

		if !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	err = categorymodel.Delete(id)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
