package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

/* エンドポイントを作成 */
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!! Golang API.") // Fprint ...書き込み先を指定して、文字列を書き込む
}

/* REST API */
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: 1, Name: "Taro", Age: 20},
	{Id: 2, Name: "Jiro", Age: 30},
	{Id: 3, Name: "Saburo", Age: 40},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	/* CORS 全許可 */
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	/* 各APIメソッド */
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(users) // JSONを返す

	case "POST":
		var user User
		json.NewDecoder(r.Body).Decode(&user) // JSONを受け取る
		fmt.Println(user)                     // 見てみる
		user.Id = len(users) + 1              // IDを設定
		users = append(users, user)           // 追加
		json.NewEncoder(w).Encode(users)      // 作成したデータを返す

	case "PUT":
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		for i, v := range users {
			if v.Id == user.Id {
				users[i] = user // 更新
				break           // 更新したら抜ける
			}
		}
		fmt.Println(user)                // 見てみる
		json.NewEncoder(w).Encode(users) // 更新したデータを返す

	case "DELETE":
		id, err := strconv.Atoi(path.Base(r.URL.Path)) // IDを取得
		fmt.Println(id)                                // 見てみる
		if err != nil {
			fmt.Println(err)
		}
		var user User
		json.NewDecoder(r.Body).Decode(&user) // JSONを受け取る
		fmt.Println(user)                     // 見てみる
		for i, v := range users {
			if v.Id == id {
				users = append(users[:i], users[i+1:]...) // 削除
			}
		}
		json.NewEncoder(w).Encode(users) // 削除したデータを返す
	}
}

func main() {
	/* ハンドラーを追加 */
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/", userHandler)

	/* サーバーを起動 */
	http.ListenAndServe(":8080", nil)
	println("run server: http://localhost:8080")
}
