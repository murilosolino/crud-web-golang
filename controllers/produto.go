package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/murilosolino/app-web/model"
)

var page = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := model.BuscaTodosOsProdutos()
	page.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	page.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryId)
	if err != nil {
		panic(err.Error())
	}
	produto := model.BuscarProdutoPorId(id)
	page.ExecuteTemplate(w, "Edit", produto)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	queryId := r.URL.Query().Get("id")

	id, err := strconv.Atoi(queryId)
	if err != nil {
		panic(err.Error())
	}

	model.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Println("Método de solicitação é diferente de POST")
		return
	}

	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoFloat, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Fatalln("Erro ao convereter preco:", err)
	}
	quantidadeInt, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Fatalln("Erro ao converter quantidade:", err)
	}

	model.CadastraProduto(nome, descricao, precoFloat, quantidadeInt)

	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Println("Método de solicitação é diferente de POST")
		return
	}

	queryId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryId)
	if err != nil {
		panic(err.Error())
	}

	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoFloat, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Fatalln("Erro ao convereter preco:", err)
	}
	quantidadeInt, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Fatalln("Erro ao converter quantidade:", err)
	}

	model.AtualizaProduto(id, nome, descricao, precoFloat, quantidadeInt)

	http.Redirect(w, r, "/", 301)

}
