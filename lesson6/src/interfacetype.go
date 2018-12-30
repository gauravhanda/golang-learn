package main

import "fmt"

//URLSafe encoder and decoder
type URLSafe interface {
	Encrypt(string) string
	Decrypt(string) string
}

//Object main object for toString
type Object interface {
	toString() string
}

// Cloneable interface
type Cloneable interface {
	clone() Cloneable
	isClone() bool
}

// URLSafeEncoderWithMultipleInterfaces implementations and implicit interfaces
// This makes the type more readable and understandable
type URLSafeEncoderWithMultipleInterfaces struct {
	Cloneable
	URLSafe
	Object
}

func main() {
	var urlEncoder = URLSafeEncoderWithMultipleInterfaces{}
	fmt.Println("Encrypted " + urlEncoder.Encrypt("gaurav") + " " + urlEncoder.toString())
	if urlEncoder.clone().isClone() {
		fmt.Println("Cloned object Found")
	} else {
		fmt.Println("Clone Object Not Found")
	}
}

//Encrypt base64 encryptor
func (t URLSafeEncoderWithMultipleInterfaces) Encrypt(string) string {
	return "Encrypted"
}

//Decrypt base64 encryptor
func (t URLSafeEncoderWithMultipleInterfaces) Decrypt(string) string {
	return "Encrypted"
}

//Decrypt base64 encryptor
func (t URLSafeEncoderWithMultipleInterfaces) toString() string {
	return "This is toString"
}

//Decrypt base64 encryptor
func (t URLSafeEncoderWithMultipleInterfaces) isClone() bool {
	return false
}

//Decrypt base64 encryptor
func (t URLSafeEncoderWithMultipleInterfaces) clone() Cloneable {
	return t
}
