package handler

import (
	models "easytrady-backend/api/Models"
	repository "easytrady-backend/api/Repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// esse diretório irá lidar com as requisições

func PostUsuario(c echo.Context) error {
	usuario := models.Usuarios{}
	err := c.Bind(&usuario)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request payload")
	}

	// Attempt to insert the user into the database
	id, err := repository.InsertUser(usuario)
	if err != nil {
		// Handle the database insertion error
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error inserting user into the database")
	}

	// Return a success response with the inserted user ID
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": fmt.Sprintf("User inserted successfully with ID: %d", id),
	})
}

// func GetUsuario(c echo.Context) error {
// 	user := models.Usuarios{"dsdaa", "Ricardo", "ric@ric.com", "1234567"}
// 	user2 := models.Usuarios{"1", "Leo", "leo@leo.com", "123456"}

// 	users := []models.Usuarios{user, user2}

// 	return c.JSON(http.StatusOK, users)
// }

// func GetProduto(c echo.Context) error {
// 	produto := models.Produtos{"dfdsf", "Mouse", "Mouse Gamer", 10.5, 10}
// 	return c.JSON(http.StatusOK, produto)
// }

// func GetVenda(c echo.Context) error {
// 	user := models.Usuarios{"dsdaa", "Ricardo", "ric@ric.com", "1234567"}
// 	produto := models.Produtos{"dfdsf", "Mouse", "Mouse Gamer", 10.5, 10}
// 	produtosSlice := []models.Produtos{produto}
// 	venda := models.Venda{"dafdfds", time.Now(), 10.5, produtosSlice, user}
// 	return c.JSON(http.StatusOK, venda)
// }

// func GetProdutoVenda(c echo.Context) error {
// 	produto := models.Produtos{"dfdsf", "Mouse", "Mouse Gamer", 10.5, 10}
// 	venda := models.Venda{"dafdfds", time.Now(), 10.5, []models.Produtos{produto}, models.Usuarios{"dsdaa", "Ricardo", "ric@ric.com", "1234567"}}
// 	produtoVenda := models.Produto_Venda{produto.ID, venda.ID, 10.5, 1}
// 	return c.JSON(http.StatusOK, produtoVenda)
// }
