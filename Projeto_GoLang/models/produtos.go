package models

import "go_modules/db"

type Curso struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Duracao   int
}

func BuscaTodosOsCursos() []Curso {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsCursos, err := db.Query("select * from curso")
	if err != nil {
		panic(err.Error())
	}

	c := Curso{}
	cursos := []Curso{}

	for selectDeTodosOsCursos.Next() {
		var id, duracao int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsCursos.Scan(&id, &nome, &descricao, &preco, &duracao)
		if err != nil {
			panic(err.Error())
		}

		c.Id = id
		c.Nome = nome
		c.Descricao = descricao
		c.Preco = preco
		c.Duracao = duracao

		cursos = append(cursos, c)
	}
	defer db.Close()
	return cursos
}
func CriaNovoCurso(nome, descricao string, preco float64, duracao int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into curso(nome, descricao, preco, duracao) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, duracao)
	defer db.Close()

}

func DeletaCurso(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOCurso, err := db.Prepare("delete from curso where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOCurso.Exec(id)
	defer db.Close()

}

func EditaCurso(id string) Curso {
	db := db.ConectaComBancoDeDados()

	cursoDoBanco, err := db.Query("select * from curso where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	cursoParaAtualizar := Curso{}

	for cursoDoBanco.Next() {
		var id, duracao int
		var nome, descricao string
		var preco float64

		err = cursoDoBanco.Scan(&id, &nome, &descricao, &preco, &duracao)
		if err != nil {
			panic(err.Error())
		}
		cursoParaAtualizar.Id = id
		cursoParaAtualizar.Nome = nome
		cursoParaAtualizar.Descricao = descricao
		cursoParaAtualizar.Preco = preco
		cursoParaAtualizar.Duracao = duracao
	}
	defer db.Close()
	return cursoParaAtualizar
}

func AtualizaCurso(id int, nome, descricao string, preco float64, duracao int) {
	db := db.ConectaComBancoDeDados()

	AtualizaCurso, err := db.Prepare("update curso set nome=$1, descricao=$2, preco=$3, duracao=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaCurso.Exec(nome, descricao, preco, duracao, id)
	defer db.Close()
}
