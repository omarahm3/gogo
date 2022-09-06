package secret

import (
	"encoding/json"
	"errors"
	"io"
	"os"
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

	err := v.load()
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

	err := v.load()
	if err != nil {
		return err
	}

	v.keyValues[key] = value

	return v.save()
}

func (v *Vault) Delete(key string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}

	delete(v.keyValues, key)

	return v.save()
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)

	if err != nil && os.IsNotExist(err) {
		v.keyValues = make(map[string]string)
		return nil
	} else if err != nil {
		return err
	}

	defer f.Close()

	r, err := encrypt.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.readKeyValues(r)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&v.keyValues)
}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := encrypt.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.writeKeyValues(w)
}

func File(key, filepath string) *Vault {
	return &Vault{
		encodingKey: key,
		filepath:    filepath,
	}
}
