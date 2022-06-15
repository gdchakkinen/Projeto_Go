package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go_modules/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsCursos := models.BuscaTodosOsCursos()
	temp.ExecuteTemplate(w, "Index", todosOsCursos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		duracao := r.FormValue("duracao")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		duracaoConvertidaParaInt, err := strconv.Atoi(duracao)
		if err != nil {
			log.Println("Erro na conversão do duracao:", err)
		}

		models.CriaNovoCurso(nome, descricao, precoConvertidoParaFloat, duracaoConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoCurso := r.URL.Query().Get("id")
	models.DeletaCurso(idDoCurso)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoCurso := r.URL.Query().Get("id")
	curso := models.EditaCurso(idDoCurso)
	temp.ExecuteTemplate(w, "Edit", curso)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		duracao := r.FormValue("duracao")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		duracaoConvertidaParaInt, err := strconv.Atoi(duracao)
		if err != nil {
			log.Println("Erro na convesão da duracao para int:", err)
		}

		models.AtualizaCurso(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, duracaoConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
