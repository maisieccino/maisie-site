package api

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maisieccino/maisie-site/internal/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pgxgeom "github.com/twpayne/pgx-geom"
	"go.uber.org/zap"
)

var configFile string

func init() {
	RootCmd.Flags().String("staticPath", ".", "path to static files to serve")
	RootCmd.Flags().StringVarP(&configFile, "config", "c", "", "path to a config file")
}

var RootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, _ []string) {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		cfg := server.Config{
			Host:   "localhost",
			Port:   8080,
			Logger: logger,
			DB: &server.DB{
				Hostname: "localhost",
				Port:     5432,
				User:     "postgres",
				Password: "",
				Database: "postgres",
			},
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(configFile)
		if configFile != "" {
			viper.SetConfigFile(configFile)
		}
		viper.SetEnvPrefix("api")
		viper.AutomaticEnv()
		if err = viper.BindPFlags(cmd.Flags()); err != nil {
			panic("binding flags " + err.Error())
		}
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

		ctx := context.Background()

		if cfg.DB != nil && cfg.DB.Enabled {

			connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
				cfg.DB.User,
				cfg.DB.Password,
				cfg.DB.Hostname,
				cfg.DB.Port,
				cfg.DB.Database,
			)
			dbCfg, err := pgxpool.ParseConfig(connStr)
			if err != nil {
				panic("building connection config: " + err.Error())
			}

			dbCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
				return pgxgeom.Register(ctx, conn)
			}

			pool, err := pgxpool.NewWithConfig(ctx, dbCfg)
			if err != nil {
				panic("error connecting to database: " + err.Error())
			}
			defer pool.Close()

			cfg.DB.Conn = pool
			logger.Debug("DB connected")
		}

		s := server.NewServer(cfg)
		s.Serve()
	},
}
