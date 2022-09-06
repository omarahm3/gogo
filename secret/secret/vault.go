package secret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/omarahm3/gogo/secret/encrypt"
)

type Vault struct {
	encodingKey string
	keyValues   map[string]string
	filepath    string
	mutex       sync.Mutex
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.loadKeyValues()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("vault: no value for that key")
	}

	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.loadKeyValues()
	if err != nil {
		return err
	}

	v.keyValues[key] = value

	return v.saveKeyValues()
}

func (v *Vault) loadKeyValues() error {
	f, err := os.Open(v.filepath)

	if err != nil && os.IsNotExist(err) {
		v.keyValues = make(map[string]string)
		return nil
	} else if err != nil {
		return err
	}

	defer f.Close()
	var sb strings.Builder

	_, err = io.Copy(&sb, f)
	if err != nil {
		return err
	}

	str, err := encrypt.Decrypt(v.encodingKey, sb.String())
	r := strings.NewReader(str)

	dec := json.NewDecoder(r)
	err = dec.Decode(&v.keyValues)

	return err
}

func (v *Vault) saveKeyValues() error {
	var sb strings.Builder

	enc := json.NewEncoder(&sb)
	err := enc.Encode(v.keyValues)
	if err != nil {
		return err
	}

	str, err := encrypt.Encrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}

	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, str)
	return err
}

func Memory(key string) Vault {
	return Vault{
		encodingKey: key,
		keyValues:   make(map[string]string),
	}
}

func File(key, filepath string) *Vault {
	return &Vault{
		encodingKey: key,
		filepath:    filepath,
	}
}
