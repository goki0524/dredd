package main

import (
	"fmt"

	"github.com/snikch/goodman/hooks"
	trans "github.com/snikch/goodman/transaction"
)

func main() {
	h := hooks.NewHooks()
	server := hooks.NewServer(hooks.NewHooksRunner(h))
	h.Before("/message > GET", func(t *trans.Transaction) {
		fmt.Println("-------------------")
		fmt.Println("before modification")
		fmt.Println("-------------------")
	})
	server.Serve()
	defer server.Listener.Close()
}
