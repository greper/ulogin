// 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
package utility

import gonanoid "github.com/matoous/go-nanoid/v2"

func NanoId(size int) string {
	return gonanoid.MustGenerate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", size)
}

func GenAppKey() string {
	return NanoId(18)
}

func GenSecret() string {
	return NanoId(32)
}

func NanoIdLong() string {
	return NanoId(36)
}

func GenSimple() string {
	return NanoId(12)
}
