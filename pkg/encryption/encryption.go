package encryption

import (
	"os"

	"github.com/palantir/go-encrypted-config-value/encryptedconfigvalue"
)

func CreateKeyFile(createKeyFiles string) error {

	keyPair, err := encryptedconfigvalue.RSA.GenerateKeyPair()
	if err != nil {
		return err
	}

	serializedPublicKey := keyPair.EncryptionKey.ToSerializable()
	serializedPrivateKey := keyPair.DecryptionKey.ToSerializable()

	fPub, err := os.Create(createKeyFiles + ".pub")
	if err != nil {
		return err
	}
	defer fPub.Close()
	fPub.Write([]byte(serializedPublicKey))

	fPri, err := os.Create(createKeyFiles)
	if err != nil {
		return err
	}
	defer fPri.Close()
	fPri.Write([]byte(serializedPrivateKey))

	return nil
}

func EncryptString(encryptString string, encryptionKeyFile string) (string, error) {

	publicKeyFileBytes, err := os.ReadFile(encryptionKeyFile)
	if err != nil {
		return "", err
	}

	privateKey, err := encryptedconfigvalue.NewKeyWithType(string(publicKeyFileBytes))
	if err != nil {
		return "", err
	}

	encryptedVal, err := encryptedconfigvalue.RSA.Encrypter().Encrypt(encryptString, privateKey)
	if err != nil {
		return "", err
	}
	return string(encryptedVal.ToSerializable()), nil
}

func CheckEncryption(config []byte) bool {
	return encryptedconfigvalue.ContainsEncryptedConfigValueStringVars(config)
}

func DecryptConfig(cfg interface{}, encryptionKeyFile string) error {

	privateKeyFileBytes, err := os.ReadFile(encryptionKeyFile)
	if err != nil {
		return err
	}

	privateKey, err := encryptedconfigvalue.NewKeyWithType(string(privateKeyFileBytes))
	if err != nil {
		return err
	}
	encryptedconfigvalue.DecryptEncryptedStringVariables(cfg, privateKey)
	return nil
}
