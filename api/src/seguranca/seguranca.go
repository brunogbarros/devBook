package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash : Retorna um hash com a senha informada pelo usuário
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificaSenha - recebe a senha e o hash, compara os dois se são iguais e retorna
func VerificaSenha(senha string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senha), []byte(hash))
}
