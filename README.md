# FalconStream

`falconstream` is event forwarder of CrowdStrike Falcon. CrowdStrike Falcon has Event Stream API and the API provides events regarding audit, malware detection and so on. `falconstream` receives the events continuously and can store them to local file system or Amazon S3. (Also Amazon Kinesis Data Firehose is planned to implement)

## Architecture

![architecture](https://user-images.githubusercontent.com/605953/66090764-b635bc80-e5bf-11e9-9d2c-c7d35c247b59.png)

`falconstream` simply receives events from CrowdStrike Falcon Event Stream API by long time HTTPS connection.

## Getting Started

### Prerequisite

- Go >= 1.13
- API key (client_id + secret) of CrowdStrike Falcon

### Setup

```bsash
go get github.com/m-mizutani/falconstream
```

### Run and output to console

```bash
$ export FALCON_CLIENT_ID=xxxxxxxxxxxxx
$ export FALCON_SECRET=xxxxxxxxxxxxxxxxxxx
$ falconstream
falconstream.falconEvent{
  MetaData: &gofalcon.StreamEventMetaData{
    CustomerIDString:  "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    EventType:         "UserActivityAuditEvent",
    Offset:            12345,
    EventCreationTime: 1568947873000,
  },
  Event: map[string]interface {}{
    "AuditKeyValues": []interface {}{
      map[string]interface {}{
        "ValueString": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "Key":         "quarantined_file_id",
      },
      map[string]interface {}{
        "Key":         "action_taken",
        "ValueString": "quarantined",
      },
    },
    "UTCTimestamp":  1568947873.000000,
    "UserId":        "Crowdstrike",
    "UserIp":        "",
    "OperationName": "quarantined_file_update",
    "ServiceName":   "quarantined_files",
  },
}
```

## Basic usage

### Output to local file system

```
$ falconstream -e fs &
$ tail -f falcon.log
{"metadata":{"customerIDString":"xxxxxxxxxxx","eventType":"AuthActivityAuditEvent","offset":1100,"eventCreationTime":1567079329516},"event":{"OperationName":"twoFactorAuthenticate","ServiceName":"CrowdStrike Authentication","Success":true,"UTCTimestamp":1567079329516,"UserId":"xxxxxxxxx","UserIp":"10.0.0.1"}}
...(snip)...
```

### Output to Amazon S3

NOTE: You need to prepare AWS credential. See [following document](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) for more detail.

```bash
$ falconstream -e s3 --aws-region ap-northeast-1 --aws-s3-bucket YOUR-BUCKET-NAME
```

### Use AWS Secrets Manager to save Falcon credentials

NOTE: You need to setup a `secret` including `falcon_client_id` and `falcon_secret` in Secrets Manager at first. Then see ARN of the `secret`.

```bash
$ falconstream --aws-secret-arn arn:aws:secretsmanager:ap-northeast-1:1234567890:secret:my-secret
```

## License

- MIT License
- Author: Masayoshi Mizutani < mizutani@sfc.wide.ad.jp >
