package security

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/distributed-fs/pkg/common"
	"github.com/google/uuid"
)

const (
	tokenLifetime = 15 * time.Minute
	test          = true
)

type RequestClaim struct {
	ApplicationId string                `json:"AppId"`
	Filename      string                `json:"filename"`
	PermittedOps  common.PermissionType `json:"allowedOperations"`
	jwt.StandardClaims
}

// getSigningKey obtains signing key. Doesn't get true key as of now for testing puproses
func getSigningKey() []byte {
	if test {
		return []byte("signingkey")
	} else {
		panic("shouldn't get here")
	}
}

// createToken creates a json web token that expires within specified timeout
func CreateToken(applicationId uuid.UUID, filename string, operations common.PermissionType) (string, error) {
	signedKey := getSigningKey()

	claims := &RequestClaim{
		applicationId.String(),
		filename,
		operations,
		jwt.StandardClaims{
			ExpiresAt: int64(tokenLifetime.Seconds()),
			Issuer:    "DFS Master",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(signedKey)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func VerifyToken(tokenString string, appId uuid.UUID, filename string, operation common.FileOperation) error {
	signingKey := getSigningKey()

	token, err := jwt.ParseWithClaims(tokenString, &RequestClaim{}, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if claims, ok := token.Claims.(*RequestClaim); ok && token.Valid {
		permissionType := common.OperationToPermissionType(operation)
		if claims.ApplicationId != appId.String() {
			return errors.New("invalid application id")
		} else if claims.Filename != filename {
			return errors.New("invalid filename")
		} else if claims.PermittedOps&permissionType == 0 {
			return errors.New("invalid operation")
		} else {
			return nil
		}
	} else {
		return err
	}
}
