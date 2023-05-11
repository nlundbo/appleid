package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/nlundbo/appleid"
)

func main() {
	v := appleid.New()
	err := v.AutoRefresh(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("test apple ID token")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		err = v.Verify(context.Background(), text)
		if err != nil {
			fmt.Println(err)
		}
	}
}
