module go-slalom

replace github.com/slalom/go-slalom => ../go-slalom

go 1.12

require (
	github.com/gorilla/mux v1.7.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.4.1
	github.com/slalom/go-slalom v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.2
)
