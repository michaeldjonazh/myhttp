package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

type contextKey struct {
	key string
}

var ConnContextKey = &contextKey{"http-conn"}

func GetConn(r *http.Request) net.Conn {
	return r.Context().Value(ConnContextKey).(net.Conn)
}

// handler function to handle incoming requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Hello, World!")
	println("HANDLER")
	// conn := GetConn(r)
	println("HANDLER GLOBALKEY: " + globalKey)
	// body, _ := io.ReadAll(r.Body)

	// println(fmt.Sprintf("Body: %v", string(body)))
	println(fmt.Sprintf("R: %v", r))
}

func SaveConnInContext(ctx context.Context, c net.Conn) context.Context {

	return context.WithValue(ctx, ConnContextKey, c)
}

var globalKey string

func newListener() (net.Listener, error) {
	//ls, err := activation.Listeners()
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	println("1111111")
	conn, err1 := ls.Accept()

	//helloHandler()
	println("2222222")
	println(err1)
	reader := bufio.NewReader(conn)
	reader.Buffered()

	println("3333333")
	for {
		// conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		// println(line)
		if strings.Contains(line, "I want") {
			println("QWEQWEQWE = " + line)
			globalKey = line

			req, rerr := http.NewRequest("GET", "test", reader)
			println(rerr)
			req.Header = http.Header{
				"bode": {line},
			}
			//req.Header.Set("bodeh", "QWER")
			helloHandler(nil, req)

			break
		}
	}

	// println("OUT")

	println("AAAAA")

	// ひとつも受け取れなかった場合でもエラーにならないため
	// len で別途確認する必要がある
	// if len(ls) == 0 {
	// 	return nil, fmt.Errorf("[CRITICAL] can't find any listeners.")
	// }
	return ls, nil
}

var globalRaw string

func SystemdListenAndServe(srv *http.Server) error {
	// var ggaData string
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := newListener()

	println("BBBBB")

	if err != nil {
		// fmt.Println(Err5, err)
		return err
	}

	defer ln.Close()
	return srv.Serve(ln)
}

func main() {
	server := http.Server{
		Addr:        ":8080",
		ConnContext: SaveConnInContext,
		// NOTE 最初の読取までのタイムアウト秒数を設定する
		// 何か入力があってからハンドラ関数が起動されてセッションが開始されるため、
		// 開始前のタイムアウトチェックはここでしかできない
		ReadTimeout: 10 * time.Second,
		// NOTE write -> rover はconnに直接IOをするので、ここで設定してはいけない
	}

	// Register the handler function for the root URL path
	http.HandleFunc("/", helloHandler)

	err4 := SystemdListenAndServe(&server)

	if err4 != nil {
		println(err4.Error())
	}
	// Start the server and listen on port 8080
	// fmt.Println("Server is listening on port 8080...")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("Server failed to start:", err)
	// }
}
