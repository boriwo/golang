//
// Golang Workshop 2024
//

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
)

// capture ctrl-c

type WebService struct {
	Port     int
	Env      string
	LogLevel string
	logger   zerolog.Logger
}

type ServerOption func(*WebService)

func WithEnv(env string) ServerOption {
	return func(s *WebService) {
		s.Env = env
	}
}

func WithLogLevel(logLevel string) ServerOption {
	return func(s *WebService) {
		s.LogLevel = logLevel
	}
}

func WithPort(port int) ServerOption {
	return func(s *WebService) {
		s.Port = port
	}
}

func NewServer(options ...ServerOption) *WebService {
	server := &WebService{
		logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
	for _, option := range options {
		option(server)
	}
	return server
}

func (s *WebService) helloHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug().Msg("hello handler")
	vars := mux.Vars(r)
	status, err := strconv.Atoi(vars["status"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	message := map[string]string{"hello": "world"}
	buf, err := json.MarshalIndent(message, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, string(buf))
}

func (s *WebService) StartService() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{status}", s.helloHandler)
	http.Handle("/", r)
	fmt.Printf("service started on port %d\n", s.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}

func main() {

	var rootCmd = &cobra.Command{Use: "webserver"}

	var startCmd = &cobra.Command{
		Use:   "run",
		Short: "Starts the web service",
		Run: func(cmd *cobra.Command, args []string) {
			port := viper.GetInt("port")
			env := viper.GetString("env")
			logLevel := viper.GetString("logLevel")
			service := NewServer(
				WithPort(port),
				WithEnv(env),
				WithLogLevel(logLevel),
			)
			service.StartService()
		},
	}

	startCmd.Flags().StringP("port", "p", "8080", "web service port")
	startCmd.Flags().StringP("env", "e", "dev", "environment for logging")
	startCmd.Flags().StringP("logLevel", "l", "debug", "log level")

	viper.BindPFlag("port", startCmd.Flags().Lookup("port"))
	viper.BindPFlag("env", startCmd.Flags().Lookup("env"))
	viper.BindPFlag("logLevel", startCmd.Flags().Lookup("logLevel"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}
