package main

import (
	"fmt"
	"strings"
)

// Define the rotor mapping (substitution table)
var rotorMapping = map[rune]rune{
	'a': 'e', 'b': 'k', 'c': 'm', 'd': 'f', 'e': 'l', 'f': 'g',
	'g': 'd', 'h': 'q', 'i': 'v', 'j': 'z', 'k': 'n', 'l': 't',
	'm': 'o', 'n': 'w', 'o': 'y', 'p': 'h', 'q': 'x', 'r': 'u',
	's': 's', 't': 'p', 'u': 'a', 'v': 'i', 'w': 'b', 'x': 'r',
	'y': 'c', 'z': 'j',
}

func main() {
	plaintext := "hello"
	encryptedText := enigmaEncrypt(plaintext)
	decryptedText := enigmaDecrypt(encryptedText)

	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Encrypted Text: ", encryptedText)
	fmt.Println("Decrypted Text: ", decryptedText)
}

func enigmaEncrypt(plaintext string) string {
	plaintext = strings.ToLower(plaintext)
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			// Map the character using the rotor
			encryptedChar, exists := rotorMapping[char]
			if exists {
				encrypted.WriteRune(encryptedChar)
			} else {
				encrypted.WriteRune(char)
			}
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func enigmaDecrypt(encryptedText string) string {

	reverseKey := make(map[rune]rune)
	for k, v := range rotorMapping {
		reverseKey[v] = k
	}

	encryptedText = strings.ToLower(encryptedText)
	var encrypted strings.Builder

	for _, char := range encryptedText {
		if char >= 'a' && char <= 'z' {
			// Map the character using the rotor
			encryptedChar, exists := reverseKey[char]
			if exists {
				encrypted.WriteRune(encryptedChar)
			} else {
				encrypted.WriteRune(char)
			}
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}
