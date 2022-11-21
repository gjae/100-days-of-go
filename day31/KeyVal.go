package main

import (
	"fmt"
	"context"
)

type aKey string
type aValue string

func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("Found value: ", v)
		return
	} else {
		fmt.Println("Key not found: ", k)
	}
}

func setValue(ctx context.Context, key aKey, val aValue) context.Context {
	return context.WithValue(ctx,key, val)
}

func main() {
	myKey := aKey("mySecretValue")
	ctx := context.WithValue(context.Background(), myKey, "secretValue")
	ctxTodo := context.TODO()
	ctxTodo = context.WithValue(ctxTodo, myKey, "Another secret value")
	ctxTodo = setValue(ctx, aKey("t"), aValue("a"))

	searchKey(ctx, myKey)
	searchKey(ctx, aKey("secret"))

	searchKey(ctxTodo, myKey)
	searchKey(ctxTodo, aKey("secert"))
	searchKey(ctxTodo, aKey("t"))

}