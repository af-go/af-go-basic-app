package cmd

import (
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	logging "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var MemoryLeakCmd = &cobra.Command{
	Use: "memory-leak",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			logging.Info(http.ListenAndServe("localhost:6060", nil))
		}()
		logging.Info("hello world")
		var wg sync.WaitGroup
		wg.Add(1)
		go leakyFunction(wg)
		wg.Wait()
	},
}

func leakyFunction(wg sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pandas")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
