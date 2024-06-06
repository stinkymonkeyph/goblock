package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    string
	mnemonic   string
}

func NewWallet() *Wallet {
	w := new(Wallet)

	// Generate a random 256-bit seed for the mnemonic
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Generate the mnemonic phrase from the entropy
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	w.mnemonic = mnemonic

	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Hash the seed to get a private key
	hash := sha256.New()
	hash.Write(seed)
	hashedSeed := hash.Sum(nil)

	privateKey, publicKey := generateKeyFromSeed(hashedSeed)
	w.privateKey = privateKey
	w.publicKey = publicKey

	// Generate address
	w.address = generateAddress(publicKey)
	return w
}

func ImportWallet(mnemonic string) (*Wallet, error) {
	w := new(Wallet)

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, fmt.Errorf("invalid mnemonic")
	}

	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Hash the seed to get a private key
	hash := sha256.New()
	hash.Write(seed)
	hashedSeed := hash.Sum(nil)

	privateKey, publicKey := generateKeyFromSeed(hashedSeed)
	w.privateKey = privateKey
	w.publicKey = publicKey

	// Generate address
	w.address = generateAddress(publicKey)
	w.mnemonic = mnemonic
	return w, nil
}

func generateKeyFromSeed(seed []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	curve := elliptic.P256()
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = curve
	privateKey.D = new(big.Int).SetBytes(seed)
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(seed)

	return privateKey, &privateKey.PublicKey
}

func generateAddress(publicKey *ecdsa.PublicKey) string {
	h2 := sha256.New()
	h2.Write(publicKey.X.Bytes())
	h2.Write(publicKey.Y.Bytes())

	digest2 := h2.Sum(nil)

	h3 := ripemd160.New()
	h3.Write(digest2)

	digest3 := h3.Sum(nil)

	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	chksum := digest6[:4]

	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chksum[:])

	return fmt.Sprintf("goblock%s", base58.Encode(dc8))
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) Address() string {
	return w.address
}

func (w *Wallet) Mnemonic() string {
	return w.mnemonic
}
