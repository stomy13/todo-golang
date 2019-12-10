package controller

import (
	"strconv"
	"encoding/json"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/MasatoTokuse/todo/model"
)

type Tasks struct {
	Tasks []model.Task
}

func Test_TasksGET_1(t *testing.T) {
	// ルーターのセットアップ
	router := setup_Router()

	// リクエストから結果を受け取る
	w:= performRequest(router, "GET", "/todo/api/v1/tasks")
	
	// JSONから構造体に変換
	var data Tasks
	json.Unmarshal([]byte(w.Body.String()), &data)
    
	// TODO:DB接続したくないため、sqlmockを使用してDBをモック化したい
	if lenTasks := len(data.Tasks);lenTasks == 3 {
		t.Log("成功しました")
	} else {
		t.Error("失敗しました") 
		t.Log("len of Tasks : " + strconv.Itoa(lenTasks) + " but expected 3")
	}
	// t.Log(data.Tasks[0].ID)
}

func setup_Router() *gin.Engine {
	router := gin.Default()
	router.GET("/todo/api/v1/tasks", TasksGET)
	return router
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}