package autenticacao

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// CriarToke : Cria o token JWT
func CriarToke(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	// secret - para gerar o hash
	return token.SignedString([]byte("só para teste"))
}
