package services

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/emidev98/raspberry-security-camara/util"
)

type tokenService struct {
	storeDir      string
	tokenFileName string
	chars         string
	length        int
}

func NewTokenService() *tokenService {
	return &tokenService{
		storeDir:      util.STORE_DIR,
		tokenFileName: "token.txt",
		chars:         "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}[]\\|:\";'<>?,./",
		length:        16,
	}
}

func (s tokenService) HandleValidateToken(w http.ResponseWriter, r *http.Request) {
	// Read the token from the request header
	// and compare it with the token stored in the file
	token := r.Header.Get("Token")

	if !s.IsValidToken(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (s tokenService) IsValidToken(token string) bool {
	file, err := os.Open(s.storeDir + s.tokenFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fileToken string
	_, err = fmt.Fscanf(file, "%s", &fileToken)
	if err != nil {
		log.Fatal(err)
	}

	return token == fileToken
}

func (s *tokenService) CreateToeknIfDoesNotExist() {
	if _, err := os.Stat(s.storeDir); os.IsNotExist(err) {
		err := os.Mkdir(s.storeDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(s.storeDir + s.tokenFileName); os.IsNotExist(err) {
		s.createToken()
	}
}

func (s *tokenService) createToken() {
	file, err := os.Create(s.storeDir + s.tokenFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	token := s.randomToken()
	_, err = file.WriteString(token)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *tokenService) randomToken() string {
	rand.Int63()

	bytes := make([]byte, s.length)
	for i := 0; i < s.length; i++ {
		bytes[i] = s.chars[rand.Intn(len(s.chars))]
	}

	return string(bytes)
}
