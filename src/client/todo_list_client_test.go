package tddClient

import (
	"testing"
	"github.com/stretchr/testify/require"
	httptransport "github.com/go-openapi/runtime/client"
	apiClient "github.com/windperson/goSwaggerDemo/src/client/client"
)

func TestGetTodoList(t *testing.T){
	//arrange
	client := apiClient.Default
	client.Transport = httptransport.New("127.0.0.1:8080","/",[]string{"http"})
	//act
	result, err := client.Todos.Find(nil)
	//assert
	if err != nil {
		require.FailNow(t,err.Error())
	}
	require.NotNil(t,*result)
}
