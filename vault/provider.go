package vault

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/99designs/keyring"
)

type EnvVarsSet = map[string]string

type KeyringProvider struct {
    Keyring keyring.Keyring
    Profile string
}

func (p *KeyringProvider) Retrieve() (val EnvVarsSet, err error) {
    log.Printf("Looking up keyring for %s", p.Profile)
    item, err := p.Keyring.Get(p.Profile)
    if err != nil {
        log.Println("Error from keyring", err)
        return val, err
    }
    if err = json.Unmarshal(item.Data, &val); err != nil {
        return val, fmt.Errorf("Invalid data in keyring: %v", err)
    }
    return val, err
}

func (p *KeyringProvider) Store(val EnvVarsSet) error {
    bytes, err := json.Marshal(val)
    if err != nil {
        return err
    }

    return p.Keyring.Set(keyring.Item{
        Key:   p.Profile,
        Label: fmt.Sprintf("env-vault (%s)", p.Profile),
        Data:  bytes,

        // specific Keychain settings
        KeychainNotTrustApplication: true,
    })
}

func (p *KeyringProvider) Delete() error {
    return p.Keyring.Remove(p.Profile)
}
