package console

import (
	"fmt"
)

func Log(m interface{}) {
	log := fmt.Sprintf("%+v", m)
	fmt.Println(log)
}
