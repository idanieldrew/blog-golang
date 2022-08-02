package token

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
	"github.com/kataras/jwt"
	"log"
	"time"
)

type (
	claims struct {
		Email string `json:"email"`
	}

	header struct {
		Kid string `json:"kid"`
		Alg string `json:"alg"`
	}
)

func GenerateToken(input1 string) (string, *restError.RestError) {
	privateKey, loadKeyErr := jwt.LoadPrivateKeyRSA(".keys/rsa_private_key.pem")
	if loadKeyErr != nil {
		logger.Error("problem in load rsa key", loadKeyErr)
		restErr := restError.ServerError("server error")
		return "", restErr
	}

	claims := claims{Email: input1}
	header := header{
		Kid: "my_key_id_1",
		Alg: jwt.RS256.Name(),
	}

	bytes, signErr := jwt.SignWithHeader(jwt.RS256, privateKey, claims, header, jwt.MaxAge(1*time.Hour))
	if signErr != nil {
		logger.Error("problem when sign", signErr)
		restErr := restError.ServerError("server error")
		return "", restErr
	}

	return string(bytes), nil
}

var keys = map[string][]byte{
	"my_key_id_1": []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw6OJ4K9LUz6MugrF7uB+
/oZw8/f3J4CSPYZFXMTsWNVQSLlen6/pr7ZvyPsgLvBGikybxRu7ff6ufmHTWTm7
mlpxEv/bgFFUmfH/faY7SA1PJcWMaEMT6s7E96orefyTMNdLi4OKhUGYJ56L8cE1
yRIya+B2UMCg2ItK11TRQlHLwvKRGsFFirc23oHX8gMuduEkIb5dSD6rEaopR3ZM
O1tipfNrlCZs5kTaIubFRJ6K1xy2Rk2hVhqdaX6Ud2aWwrb7o21REkDbqY9YuOGV
/FnDiqDtIoS7MHl5CAguaL9YiOv3RRvCrUttfuHqbljlD7m6/69rMB1cVfbdr5IB
RQIDAQAB
-----END PUBLIC KEY-----
`),
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")

		x, verifyErr := jwt.VerifyWithHeaderValidator(jwt.RS256, keys, []byte(auth), validateHeader)
		if verifyErr != nil {
			logger.Error("token not valid", verifyErr)
			restError.UnauthorizedError("token not valid")
		}
		var c claims
		err := x.Claims(&c)
		if err != nil {
			log.Fatalln(err, "ojo")
		}
		context.Next()
	}
}

func validateHeader(alg string, headerDecoded []byte) (jwt.Alg, jwt.PublicKey, jwt.InjectFunc, error) {
	var h header
	err := jwt.Unmarshal(headerDecoded, &h)
	if err != nil {
		return nil, nil, nil, err
	}

	if h.Alg != alg {
		return nil, nil, nil, jwt.ErrTokenAlg
	}

	if h.Kid == "" {
		return nil, nil, nil, fmt.Errorf("kid is empty")
	}

	key, ok := keys[h.Kid]
	if !ok {
		return nil, nil, nil, fmt.Errorf("unknown kid")
	}

	publicKey, err := jwt.ParsePublicKeyRSA(key)
	if err != nil {
		return nil, nil, nil, jwt.ErrTokenAlg
	}

	return nil, publicKey, nil, nil
}
