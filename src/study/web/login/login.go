package login

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) //リクエストを取得するメソッド
	fmt.Println("parh: ", r.URL.Path) //Path
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		// fmt.Println("crutime: ", crutime)
		h := md5.New()
		// fmt.Println("h: ", h)
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		// fmt.Println("token: ", token)
		t, _ := template.ParseFiles("login/login.gohtml")
		t.Execute(w, token)
	} else {
		//ログインデータがリクエストされ、ログインのロジック判断が実行されます。
		r.ParseForm()
		// token check
		token := r.Form.Get("token")
		if token != "" {
			// tokenの合法性を検証
			fmt.Println(token)
		} else {
			// tokenが存在しなければエラーを出す
			fmt.Println(token)
		}
		validation(r) //Fromデータのvalidationを実行
		fmt.Println("mail:", r.Form["mail"], template.HTMLEscapeString(r.Form.Get("mail")))
		fmt.Println("password:", r.Form["password"], template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("mail")))
	}
}

func validation(r *http.Request) {
	// mailがからだった時のエラー処理
	// if len(r.Form["mail"][0]) == 0 {
	// 	fmt.Println("[ERR] mail must filled!")
	// }
	// mailがメールアドレス形式でなかった時のエラー処理
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("mail")); !m {
		fmt.Println("mail filed is invalid")
	} else {
		fmt.Println("valid")
	}
}
