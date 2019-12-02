package vault

import (
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	yaml "gopkg.in/yaml.v1"
)

func unmarshalCredentials(bytes []byte) (*Credentials, error) {
	var creds Credentials
	err := yaml.Unmarshal(bytes, &creds)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling lease: %v", err)
	}
	return &creds, nil
}

func (c *Credentials) Save(path string) error {
	//write out full secret
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshalling creds: %v", err)
	}

	err = ioutil.WriteFile(path, bytes, 0600)
	if err != nil {
		return fmt.Errorf("error writing creds to file: %v", err)
	}

	log.Printf("wrote lease to %s", path)
	return nil
}
