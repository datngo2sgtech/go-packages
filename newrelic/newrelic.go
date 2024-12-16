package newrelic

import (
	nragent "github.com/newrelic/go-agent/v3/newrelic"

	"github.com/saigontechnology/Golang-Project-Base/go-packages/must"
)

func CreateNewRelicApp() *nragent.Application {
	app, err := nragent.NewApplication(
		nragent.ConfigFromEnvironment(),
	)
	if nil != err {
		must.NotFail(err)
	}

	return app
}
