package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
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

// CriarToken : Cria o token JWT
func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	// secret é passada como assinatura
	return token.SignedString(config.SecretKey)
}

// ValidarToken : verifica se o token passado pelo usuario é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveVerificacao)
	if erro != nil {
		return erro
	}
	// token correto e no formato correto
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	// Bearer 123 - token possui 2 partes, pegamos a segunda (o token)
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornaChaveVerificacao(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura incorreto: %v ", token.Header["alg"])
	}
	return config.SecretKey, nil
}

//ExtrairUsuarioID : retorna o ID do token
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		// TODO: understand why this claimns change
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioI"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("token invalido")
}
