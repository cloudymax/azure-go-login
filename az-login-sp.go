//usr/bin/env go run $0 $@ ; exit
//
//  dep ensure
//  az ad sp create-for-rbac --sdk-auth > my.auth
//  export AZURE_AUTH_LOCATION=my.auth
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

// Accepts a string path, returns json data
func readJSON(path string) (*map[string]interface{}, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.Wrap(err, "Can't open the file")
	}

	contents := make(map[string]interface{})
	err = json.Unmarshal(data, &contents)

	if err != nil {
		err = errors.Wrap(err, "Can't unmarshal file")
	}

	return &contents, err
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

	authInfo, err := readJSON(os.Getenv("AZURE_AUTH_LOCATION"))

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
