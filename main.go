package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/m-mizutani/falconstream/pkg/falconstream"
)

var logLevelMap = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

func main() {
	var logLevel string
	var args falconstream.Arguments

	app := cli.NewApp()
	app.Name = "falconstream"
	app.Usage = "Event forwarder for CrowdStrike Falcon"
	app.Version = falconstream.Version
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Masayoshi Mizutani",
			Email: "mizutani@sfc.wide.ad.jp",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "log-level, l", Value: "info",
			Usage:       "Log level [trace,debug,info,warn,error]",
			Destination: &logLevel,
		},

		cli.StringFlag{
			Name: "emitter, e", Value: "console",
			Usage:       "Choose emitter [console,fs,s3]",
			Destination: &args.Emitter.Type,
		},

		cli.StringFlag{
			Name:        "aws-secret-arn",
			Usage:       "AWS SecretsManager ARN for Falcon credentials",
			Destination: &args.Credential.AwsSecretsManagerARN,
		},
		cli.StringFlag{
			Name:        "aws-region",
			Usage:       "AWS Region",
			Destination: &args.Emitter.AwsRegion,
		},
		cli.StringFlag{
			Name:        "aws-s3-bucket",
			Usage:       "AWS S3 Bucket for S3 emitter",
			Destination: &args.Emitter.AwsS3Bucket,
		},
		cli.StringFlag{
			Name:        "aws-s3-prefix",
			Usage:       "AWS S3 prefix for S3 emitter",
			Destination: &args.Emitter.AwsS3Prefix,
		},
	}

	app.Action = func(c *cli.Context) error {
		level, ok := logLevelMap[logLevel]
		if !ok {
			return fmt.Errorf("Invalid log level: %s", logLevel)
		}
		falconstream.Logger.SetLevel(level)

		falconstream.Logger.WithFields(logrus.Fields{
			"args":     args,
			"logLevel": logLevel,
		}).Debug("Given options")

		stream := falconstream.NewStream()

		if err := stream.Start(args); err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		falconstream.Logger.WithError(err).Fatal("Fatal Error")
	}
}
