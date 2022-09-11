package console

import (
	"encoding/json"
	"fmt"
)

func Log(m interface{}) {
	logs, _ := json.MarshalIndent(m, "", "")
	fmt.Printf("============== %s\n", logs)
}
