package api

import (
	"github.com/hiMaisie/maisie-site/internal/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	RootCmd.Flags().String("staticPath", ".", "path to static files to serve")
}

var RootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		cfg := server.Config{
			Host:   "localhost",
			Port:   8080,
			Logger: logger,
		}

		viper.AddConfigPath(".")
		viper.SetEnvPrefix("api")
		viper.AutomaticEnv()
		viper.BindPFlags(cmd.Flags())
		if err = viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logger.Warn("no config file found, continuing...")
			} else {
				panic("error reading in config file: " + err.Error())
			}
		}
		if err = viper.Unmarshal(&cfg); err != nil {
			panic("error unmarshalling config: " + err.Error())
		}
		viper.Debug()

		s := server.NewServer(cfg)
		s.Serve()
	},
}
