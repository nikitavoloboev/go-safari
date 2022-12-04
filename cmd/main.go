package safari

import (
	"fmt"
	"log"

	"github.com/nikitavoloboev/go-safari"
)

func main() {
	url, err := safari.GetCurrentSafariURL()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
