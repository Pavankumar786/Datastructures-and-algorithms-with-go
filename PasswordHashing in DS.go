//1.Password Hashing
package main

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    password := "Pavan@1234"
    hash, _ := HashPassword(password) 

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
}
$go run main.go
Password: :Pavan@1234
Hash:     $2a$14$EvenbM.4d.uUd1uMf98ZfOMsKGIHqmL0dZmpj4kq.8sUYVPQHya.y
Match:    tru