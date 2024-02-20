package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	// Записываем в переменную результат функции
	cl, err := selectClient(db, clientID)
	// Проверяем что вернулась ошибка nil, иначе тест будет завершен
	require.NoError(t, err)
	// Сравниваем значения из ответа со значением переменной clientID
	require.Equal(t, clientID, cl.ID)
	// Проверяем, что остальные поля не пустые
	assert.NotEmpty(t, cl)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1
	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	// Проверяем, что получается соответсвующая ошибка
	require.ErrorIs(t, err, sql.ErrNoRows)
	// Проверяем, что поля ответа пустые
	assert.Empty(t, cl)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	// Проверяем, что функция вернула не пустой идентификатор
	require.NotNil(t, id)
	// Проверяем, что функция вернула пустую ошибку
	require.NoError(t, err)

	client, err := selectClient(db, id)
	require.NoError(t, err)
	// Проверяем, что что значения всех полей полученного объекта совпадают со
	// значениями полей объекта в переменной cl
	cl = Client{
		ID:       id,
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	require.Equal(t, cl, client)

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	// Проверяем, что функция вернула не пустой идентификатор
	require.NotNil(t, id)
	// и пустую ошибку
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.ErrorIs(t, err, sql.ErrNoRows)

}
