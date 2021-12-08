package auth_from_file

import (
	"fmt"
	"os"
	"Azure_golang_modules/auth_from_file/auth_from_file"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

// AzureSession is an object representing session our connection
// it takes an <Authorizer> and a <SubscriptionID>
type AzureSession struct {
	SubscriptionID string
	Authorizer     autorest.Authorizer
}

//creates a new <session> from file
func newSessionFromFile(quiet bool) (*AzureSession, error) {
	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)

	if err != nil {
		return nil, errors.Wrap(err, "Can't initialize authorizer")
	} else {
		if quiet != true {
			spew.Dump(authorizer)
		}
	}

	authInfo, err := auth_from_file.readJSON(os.Getenv("AZURE_AUTH_LOCATION"))

	if err != nil {
		return nil, errors.Wrap(err, "Can't get authinfo")
	} else {
		if quiet != true {
			spew.Dump(authorizer)
		}
	}

	sess := AzureSession{
		SubscriptionID: (*authInfo)["subscriptionId"].(string),
		Authorizer:     authorizer,
	}

	if quiet != true {
		spew.Dump(authorizer)
	}

	return &sess, nil
}

func main() {

	// create a new session using an auth file
	sess, err := newSessionFromFile(false)

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%v\n", sess.SubscriptionID)
	}

}
