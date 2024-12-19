package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type contextKey string

const testHeader = "testheader" // "Testheader"

func main() {
	// Error return value of `os.Open` is not checked (errcheck)
	file, _ := os.Open("hogetest.txt")

	// pass []any as any to func printAll func(args ...any) (asasalint)
	// NOTE: そもそもこんな私かたはNGだけど。。。
	args := []any{1, "hello", 3.14}
	printTestForAsasalint(args)

	// 可変長引数として
	// asasalint: 成功
	printTestForAsasalint(1, "hello", 3.14)

	v := http.Header{}
	// const "testHeader" used as a key at http.Header, but "testheader" is not canonical, want "Testheader" (canonicalheader)
	v.Get(testHeader)
	v.Get("Test-HEader")          // v.Get("Test-Header")
	v.Set("Test-header", "value") // v.Get("Test-Header")

	// main.go:31:23: response body must be closed (bodyclose)
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
	}
	// defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// メイン関数でコンテキストを初期化
	rootCtx := context.Background()
	call1(rootCtx)

	// ミススペル検出
	// fmt.Println("stady")

	unusedSlice := make([]int, 10)
	_, _ = fmt.Println(unusedSlice, file)
}

// asasalintのテスト
func printTestForAsasalint(args ...any) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// contextcheckのチェック
func call1(ctx context.Context) {
	ctx = getNewCtx(ctx) // 新しいコンテキストを親から受け継ぐ → OK
	call2(ctx)

	// Non-inherited new context, use function like `context.WithXXX` instead (contextcheck)
	call2(context.Background()) // 親の情報を持たない新しいコンテキストを使用 → NG

	// Function `call3` should pass the context parameter (contextcheck)
	call3() // コンテキストを引数として渡していない → NG
}

func call2(ctx context.Context) {
}

func call3() {
	ctx := context.TODO()
	call2(ctx)
}

// nolint: contextcheck
func call4() {
	getNewCtx(context.Background()) // 警告をスキップ
}

func getNewCtx(ctx context.Context) context.Context {
	newCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	_ = cancel
	return newCtx
}
