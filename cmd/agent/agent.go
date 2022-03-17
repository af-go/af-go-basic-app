package agent

import (
	"context"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/af-go/basic-app/api"
	"github.com/af-go/basic-app/pkg/model"
	"github.com/af-go/basic-app/pkg/utils"
	logging "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

// AgentCmd agent command
var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "start agent",
	Run: func(cmd *cobra.Command, args []string) {
		logging.Info("start agent\n")

		var options Options

		err := utils.Load(configFile, &options)
		if err != nil {
			logging.WithError(err).Warnf("failed to load config file %s, use default", configFile)
		}

		ctx := context.Background()

		stopCh := make(chan os.Signal, 1)
		signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)

		apiController := api.NewRestfulController(options.Agent)
		if apiController == nil {
			logging.Error("cannot create API controller")
			return
		}
		apiController.Start(ctx)

		select {
		case s := <-stopCh:
			logging.Warningf("received signal %v, shutting down ...", s)
		}
		apiController.Stop(ctx)
	},
}

func init() {
	AgentCmd.Flags().StringVarP(&configFile, "config", "c", path.Join(os.Getenv("HOME"), ".af-go", "profiling-config.json"), "config file")
}

// Options Options
type Options struct {
	Agent model.RestServiceOptions `json:"agent" yaml:"agent"`
}
