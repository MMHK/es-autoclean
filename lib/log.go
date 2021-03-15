package lib

import "github.com/op/go-logging"

var log = logging.MustGetLogger("es-autoclean")

func init() {
	format := logging.MustStringFormatter(
		`es-autoclean %{color} %{shortfunc} %{level:.4s} %{shortfile}
%{id:03x}%{color:reset} %{message}`,
	)
	logging.SetFormatter(format)
}
