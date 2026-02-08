package model

import (
	"database/sql"
	"log"

	"github.com/murilosolino/app-web/config/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func CadastraProduto(nome string, descricao string, preco float64, quantidade int) {
	produto := Produto{
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}
	err := insereProduto(produto)
	if err != nil {
		panic(err.Error())
	}

}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	produto := Produto{
		Id:         id,
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}
	editaProduto(produto)
}

func DeletaProduto(id int) {
	db := db.ConectaComBancoMysql()
	defer db.Close()

	prepare, err := db.Prepare("DELETE FROM produtos WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	prepare.Exec(id)

}

func BuscarProdutoPorId(id int) Produto {
	db := db.ConectaComBancoMysql()
	defer db.Close()
	var quantidade int
	var nome, descricao string
	var preco float64
	err := db.QueryRow("SELECT * FROM produtos WHERE id=?", id).Scan(&id, &nome, &descricao, &preco, &quantidade)
	if err == sql.ErrNoRows {
		log.Printf("no user with id %d\n", id)
	} else if err != nil {
		log.Fatalf("query error: %v\n", err)
	}

	return Produto{Id: id, Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}

}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoMysql()
	defer db.Close()

	listaDeProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	var id, quantidade int
	var nome, descricao string
	var preco float64
	p := Produto{}
	produtos := []Produto{}

	for listaDeProdutos.Next() {
		listaDeProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	return produtos
}

func editaProduto(produto Produto) {
	db := db.ConectaComBancoMysql()
	defer db.Close()

	prepare, err := db.Prepare("UPDATE produtos SET nome = ?, descricao = ?, preco = ?, quantidade = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)
}

func insereProduto(produto Produto) error {
	db := db.ConectaComBancoMysql()
	defer db.Close()

	prepare, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	prepare.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	return nil
}
