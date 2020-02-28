package helpers

/*
 * This file is part of the Ocrend Framewok package.
 *
 * (c) Ocrend Software <info@ocrend.com>
 * @author Brayan Narváez <prinick@ocrend.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
*/

import (
	"fmt"
	"reflect"
	"crypto/rand"
	randomize "math/rand"
	"strings"
	"time"
	"crypto/md5"
    "encoding/hex"
)

/**
	* Genera una contraseña aleatoria
*/
func RandomPassword(length int) string {
	randomize.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNÑOPQRSTUVWXYZ" +
		"abcdefghijklmnñopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[randomize.Intn(len(chars))])
	}

	return b.String()
}

/**
	* Genera un token aleatorio
*/
func TokenGenerator() string {
	b := make([]byte, 50)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/**
	* Devuelve un HASH MD5
*/
func Encrypt(str string) string {
	hasher := md5.New()
    hasher.Write([]byte(str))
    return hex.EncodeToString(hasher.Sum(nil))
}

/**
	* Verifica si una variable está vacía y devuelve true cuando lo está
*/
func Empty(val interface{}) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}