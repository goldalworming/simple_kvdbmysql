package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func EncryptPasswd(passwd string) (passwdhash string,err error) {
    password := []byte(passwd)
    // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        return "",err
    }
    return string(hashedPassword),nil
}
func ComparePasswd(passwd string, dbpasswd string) (err error) {
    password := []byte(passwd)
    dbpassword := []byte(dbpasswd)
    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword(dbpassword, password)
    if(err!=nil){
        return err
    }
    return nil
}