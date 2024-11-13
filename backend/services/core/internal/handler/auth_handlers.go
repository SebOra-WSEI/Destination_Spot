package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const (
	AuthServiceHost     string = "http://host.docker.internal:8081"
	AuthorizationHeader        = "Authorization"
)

func SignIn(c *gin.Context) {
	statusCode, res := makePostCall(c, "sign-in", model.LoggedUserResponse{})
	c.JSON(statusCode, response.Create(res))
}

func SignUp(c *gin.Context) {
	statusCode, res := makePostCall(c, "sign-up", model.NoPasswordUser{})
	c.JSON(statusCode, response.Create(res))
}

func ResetPassword(c *gin.Context) {
	statusCode, res := makePutCall(c, "reset-password")
	c.JSON(statusCode, response.Create(res))
}

func AccessControl(c *gin.Context) {
	statusCode, res := makePutCall(c, "access-control")
	c.JSON(statusCode, response.Create(res))
}

func makePostCall(c *gin.Context, address string, t interface{}) (statusCode int, r interface{}) {
	res, err := http.Post(AuthServiceHost+"/"+address, "application/json", c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, gin.H{"error": response.ErrRequestNotExecuted}
	}

	body, err := handleBody(res, t)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err}
	}

	return res.StatusCode, body
}

func makePutCall(c *gin.Context, address string) (statusCode int, r interface{}) {
	id := c.Param("id")

	url := fmt.Sprintf("%s/%s/%s", AuthServiceHost, address, id)

	req, err := http.NewRequest(http.MethodPut, url, c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, gin.H{"error": response.ErrRequestNotExecuted}
	}

	req.Header.Set(AuthorizationHeader, c.GetHeader(AuthorizationHeader))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, gin.H{"error": response.ErrRequestNotExecuted}
	}

	body, err := handleBody(res, model.UserResponseWithAction{})
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err}
	}

	return res.StatusCode, body
}

func handleBody[T any](res *http.Response, t T) (interface{}, error) {
	b, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf(err.Error())
		return t, fmt.Errorf("Problem with reading body")
	}

	var responseBody struct {
		Response T      `json:"response"`
		Error    string `json:"error"`
	}

	unmarshalErr := json.Unmarshal(b, &responseBody)
	if unmarshalErr != nil {
		fmt.Printf(unmarshalErr.Error())
		return t, fmt.Errorf("Problem with parsing body")
	}

	if responseBody.Error != "" {
		return response.CreateError(errors.New(responseBody.Error)), nil
	}

	return responseBody.Response, nil
}
