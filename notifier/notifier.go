package notifier

import (
	"github.com/getsentry/sentry-go"
	"time"
)

var configured bool

func Init(dsn string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            true,
		AttachStacktrace: true,
	})
	if err != nil {
		return err
	}

	configured = true

	return nil
}

// Flush buffered events before the program terminates.
// Set the timeout to the maximum duration the program can afford to wait.
func Flush() {
	sentry.Flush(2 * time.Second)
}

func NotifyError(err error) {
	if configured {
		sentry.CaptureException(err)
	}
}
