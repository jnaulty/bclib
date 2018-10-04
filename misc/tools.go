package misc

import (
  "crypto/sha256"
  "golang.org/x/crypto/ripemd160"
  "encoding/hex"

  log "github.com/sirupsen/logrus"
)

// TODO: Maybe can optimize
func ReverseHex(b []byte) []byte {
  newb := make([]byte, len(b))
  copy(newb, b)
  for i := len(newb)/2 - 1; i >= 0; i-- {
    opp := len(newb) - 1 - i
    newb[i], newb[opp] = newb[opp], newb[i]
  }

  return newb
}

func HexToBinary(src []byte) []byte {
  b := make([]byte, hex.DecodedLen(len(src)))
  n, err := hex.Decode(b, src)
  if err != nil {
    log.Warn(err)
  }
  return b[:n]
}

// 65 bytes long ECDSA public key to address hash
// Fist byte is always 0x4 followed by two 32 bytes components
func EcdsaToPkeyHash(input []byte) []byte {
  if input[0] == 0x04 {
    output := make([]byte, 1, 24)
    output[0] = 0x00

    hash := sha256.New()          // Intermediate SHA256 hash computations
    hash.Write(input[:])
    sha256 := hash.Sum(nil)
    hash = ripemd160.New()
    hash.Write(sha256)

    output = append(output, hash.Sum(nil)...)    // hash160
    checksum := DoubleSha256(output)[0:4]
    output = append(output, checksum...)
    return output
  }
  return nil
}

/*
* 33 byte long compressed ECDSA public key
* Fist byte is always 0x4 followed by the 32 bytes component
*/
func ShortEcdsaToPkeyHash(input []byte) []byte {
  log.Fatal("Short ECDSA")
  if input[0] == 0x02 || input[0] == 0x03 {
  }
  return nil
}
