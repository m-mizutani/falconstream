package falconstream

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
)

// CredentialArguments manages arguments to retrieve credential (client_id and secret)
type CredentialArguments struct {
	AwsSecretsManagerARN string
}

type credential struct {
	ClientID string `json:"falcon_client_id"`
	Secret   string `json:"falcon_secret"`
}

func getCredential(args CredentialArguments) (*credential, error) {
	var cred credential
	if args.AwsSecretsManagerARN != "" {
		if err := getSecretValues(args.AwsSecretsManagerARN, &cred); err != nil {
			return nil, errors.Wrap(err, "Fail to get Falcon credential from SecretsManager")
		}

		return &cred, nil
	}

	clientID, secret := os.Getenv("FALCON_CLIENT_ID"), os.Getenv("FALCON_SECRET")
	if clientID != "" && secret != "" {
		cred.ClientID = clientID
		cred.Secret = secret
		return &cred, nil
	}

	return nil, fmt.Errorf("No available credentials (requied EnvVar or AWS SecretsManager ARN")
}

func getSecretValues(secretArn string, values interface{}) error {
	// sample: arn:aws:secretsmanager:ap-northeast-1:1234567890:secret:mytest
	arn := strings.Split(secretArn, ":")
	if len(arn) != 7 {
		return errors.New(fmt.Sprintf("Invalid SecretsManager ARN format: %s", secretArn))
	}
	region := arn[3]

	ssn := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	mgr := secretsmanager.New(ssn)

	result, err := mgr.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretArn),
	})

	if err != nil {
		return errors.Wrap(err, "Fail to retrieve secret values")
	}

	err = json.Unmarshal([]byte(*result.SecretString), values)
	if err != nil {
		return errors.Wrap(err, "Fail to parse secret values as JSON")
	}

	return nil
}
