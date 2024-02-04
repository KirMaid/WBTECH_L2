package ntptime

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	res, err := ntp.Query("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось получить время: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Текущее время: ", time.Now().UTC())
	fmt.Println("Точное время: ", res.Time)
}
