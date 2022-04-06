package autenticacao

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// O TOKEN É GERADO APENAS UMA VEZ, LOGO, ESTA FUNCAO É INUTIL APÓS
//func init() {
//	chave := make([]byte, 64)
//
//	if _, erro := rand.Read(chave); erro != nil {
//		log.Fatal(erro)
//	}
//	// secret final gerado randomicamente
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//}

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
