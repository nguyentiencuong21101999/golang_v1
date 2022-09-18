package console

import (
	"encoding/json"
	"fmt"
)

func Log(m interface{}) {
	logs, err := json.MarshalIndent(m, "", "")
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Printf("============== %s\n", logs)
}
