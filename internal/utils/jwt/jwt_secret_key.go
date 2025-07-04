package jwt_utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

const SECRET_KEY_LEN = 256
const SECRET_KEY_NAME = "JWT_SECRET"
const ENV_FILE_PATH = "config/.env"

func GenerateAndSaveJWTSecretKey() {
	secret, err := generateSecret()
	if err != nil {
		panic(err)
	}

	err = saveSecret(secret)
	if err != nil {
		panic(err)
	}
}

func generateSecret() (string, error) {
	bytes := make([]byte, SECRET_KEY_LEN)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func saveSecret(secret string) error {
	fileBytes, err := os.ReadFile(ENV_FILE_PATH)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	lines := []string{}
	if err == nil {
		lines = strings.Split(string(fileBytes), "\n")
	}

	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, SECRET_KEY_NAME+"=") {
			lines[i] = fmt.Sprintf("%s=%s", SECRET_KEY_NAME, secret)
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", SECRET_KEY_NAME, secret))
	}

	output := strings.Join(lines, "\n")

	if len(fileBytes) > 0 && strings.HasSuffix(string(fileBytes), "\n") {
		output += "\n"
	}

	return os.WriteFile(ENV_FILE_PATH, []byte(output), 0644)
}
