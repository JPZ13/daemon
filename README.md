## [Daemon](https://www.youtube.com/watch?v=9VIjJvWoKLY)

```go
import (
	"fmt"
  "time"

	"github.com/JPZ13/daemon"
)

// plug in any function that takes no arguments
// and returns an error
func action() error {
	return nil
}

func main() {
	dae := daemon.New(dae.Config{
		Action:   action,
		Interval: 2 * time.Minute,
	})

	errCh := dae.Start()
	go func() {
		for err := range errCh {
			fmt.Println("Handling error ", err)
		}
	}()

	time.Sleep(5 * time.Minute)

	dae.Stop()
}
```
