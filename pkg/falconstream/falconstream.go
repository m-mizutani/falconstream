package falconstream

import (
	"github.com/m-mizutani/gofalcon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// Version number
	Version = "v0.3.0"
)

// Logger is exposed to control logging behavior from outside
var Logger = logrus.New()

// SetGoFalconLoggerLevel changes log level of gofalcon
func SetGoFalconLoggerLevel(level logrus.Level) {
	gofalcon.Logger.SetLevel(level)
}

// Arguments includes all parameters that can be controlled from outside for Stream.
type Arguments struct {
	Endpoint   string
	Credential CredentialArguments
	Emitter    EmitterArguments
}

// Stream is main interface of falconstream
type Stream struct {
}

// NewStream is constructor of Stream
func NewStream() *Stream {
	return &Stream{}
}

// Start invokes all procedure of falconstream
func (x *Stream) Start(args Arguments) error {
	cred, err := getCredential(args.Credential)
	if err != nil {
		return err
	}

	client := gofalcon.NewClient()
	client.Endpoint = args.Endpoint

	if err := client.EnableOAuth2(cred.ClientID, cred.Secret); err != nil {
		return err
	}

	emitter := newEmitter(args.Emitter)
	if err := emitter.setup(); err != nil {
		return err
	}

	ch := client.Sensor.EventStream(nil)
	for q := range ch {
		if q.Error != nil {
			return errors.Wrap(q.Error, "Fail in EventStream")
		}

		ev := new(falconEvent)
		ev.MetaData = q.Meta
		ev.Event = q.Event

		if err := emitter.emit(ev); err != nil {
			return errors.Wrap(err, "Fail to emit falconEvent")
		}
	}

	Logger.Warn("EventStream has shut down")
	if err := emitter.teardown(); err != nil {
		return err
	}

	return nil
}

// Stop cancels all procedure of falconstream (maybe)
func (x *Stream) Stop() error {
	return nil
}
