package signature

import (
        "crypto/hmac"
  	    "crypto/sha256"
        "crypto/md5"
	      "encoding/hex"
        "fmt"
        "net/url"
    		"crypto/rand"
    		"math/big"
)

const (
	ig_key	=	"fc4720e1bf9d79463f62608c86fbddd374cc71bbfb98216b52e3f75333bd130d"
	ig_key_version = "4"
	android_seed = "mxfivf3ju3dudip1ffcd"
)

func GenerateSignature(data string) string {
	ig_key_byte := []byte(ig_key)
  mac := hmac.New(sha256.New, ig_key_byte)
  mac.Write([]byte(data))
  hash := hex.EncodeToString(mac.Sum(nil))
  encoded := url.QueryEscape(data)
  sig := fmt.Sprintf("ig_sig_key_version=%s&signed_body=%s.%s", ig_key_version, hash, encoded)
  return sig
}

func GenerateDeviceID(seed string) string {
  seed_x := fmt.Sprintf("%s%s", seed, android_seed)
  seed_byte := []byte(seed_x)
  hash_init := md5.Sum(seed_byte)
	hash_string := fmt.Sprintf("%x", hash_init)
	hash_substr := fmt.Sprintf("android-%s", hash_string[:16])
  return hash_substr  
}

func GenerateUUID(stripes bool) string {
	pos1 := GenerateNumberUUIDX()
	pos2 := GenerateNumberUUIDX()
	pos3 := GenerateNumberUUIDX()
	pos4 := GenerateNumberUUIDY() + 16384
	pos5 := GenerateNumberUUIDZ() + 32768
	pos6 :=	GenerateNumberUUIDX()
	pos7 :=	GenerateNumberUUIDX()
	pos8 := GenerateNumberUUIDX()
	if stripes {
		return fmt.Sprintf("%04x%04x-%04x-%04x-%04x-%04x%04x%04x", pos1, pos2, pos3, pos4, pos5, pos6, pos7, pos8)
	} else {
		return fmt.Sprintf("%04x%04x%04x%04x%04x%04x%04x%04x", pos1, pos2, pos3, pos4, pos5, pos6, pos7, pos8)
	}
}

func GenerateNumberUUIDX() int64 {
	gen, err := rand.Int(rand.Reader, big.NewInt(65535))
	if err != nil{
		panic(err)
	}
	out := gen.Int64()
	return out
}

func GenerateNumberUUIDY() int64 {
	gen, err := rand.Int(rand.Reader, big.NewInt(4095))
	if err != nil{
		panic(err)
	}
	out := gen.Int64()
	return out
}

func GenerateNumberUUIDZ() int64 {
	gen, err := rand.Int(rand.Reader, big.NewInt(16383))
	if err != nil{
		panic(err)
	}
	out := gen.Int64()
	return out
}