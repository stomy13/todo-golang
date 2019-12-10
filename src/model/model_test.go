package model

import (
	"testing"
)

// 疎通確認
func Test_DBConnect_1(t *testing.T) {
	db := DBConnect();
    _, err := db.Query("SELECT * FROM task ORDER BY id DESC")
    if err == nil {
		t.Log("疎通確認のテストがパスしました")
    } else {
		t.Error("疎通確認のテストが通りません") 
	}
}
