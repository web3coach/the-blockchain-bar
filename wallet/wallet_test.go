// Copyright 2020 The the-blockchain-bar Authors
// This file is part of the the-blockchain-bar library.
//
// The the-blockchain-bar library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The the-blockchain-bar library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/fs"
)

// The password for testing keystore files:
// 	./node/test_andrej--3eb92807f1f91a8d4d85bc908c7f86dcddb1df57
// 	./node/test_babayaga--6fdc0d8d15ae6b4ebf45c52fd2aafbcbb19a65c8
const testKeystoreAccountsPwd = "security123"

// Prints a PK:
//
// (*ecdsa.PrivateKey)(0xc000099980)({
// PublicKey: (ecdsa.PublicKey) {
//  Curve: (*secp256k1.BitCurve)(0xc0000982d0)({
//   P: (*big.Int)(0xc0000b03c0)(115792089237316195423570985008687907853269984665640564039457584007908834671663),
//   N: (*big.Int)(0xc0000b0400)(115792089237316195423570985008687907852837564279074904382605163141518161494337),
//   B: (*big.Int)(0xc0000b0440)(7),
//   Gx: (*big.Int)(0xc0000b0480)(55066263022277343669578718895168534326250603453777594175500187360389116729240),
//   Gy: (*big.Int)(0xc0000b04c0)(32670510020758816978083085130507043184471273380659243275938904335757337482424),
//   BitSize: (int) 256
//  }),
//  X: (*big.Int)(0xc0000b1aa0)(1344160861301624411922901086431771879005615956563347131047269353924650464711),
//  Y: (*big.Int)(0xc0000b1ac0)(73524953917715096899857106141372214583670064515671280443711113049610951453654)
// },
// D: (*big.Int)(0xc0000b1a40)(41116516511979929812568468771132209652243963107895293136581156908462828164432)
//})
//
// And r, s, v signature params:
//
// (*big.Int)(0xc0000b1b20)(88181292280759186801869952076472415807575357966745986437065510600744149574656)
// (*big.Int)(0xc0000b1b40)(23476722530623450948411712153618947971604430187320320363672662539909827697049)
// (*big.Int)(0xc0000b1b60)(1)
func TestSignCryptoParams(t *testing.T) {
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(privKey)

	msg := []byte("the Web3Coach students are awesome")

	sig, err := Sign(msg, privKey)
	if err != nil {
		t.Fatal(err)
	}

	if len(sig) != crypto.SignatureLength {
		t.Fatal(fmt.Errorf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}

	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	v := new(big.Int).SetBytes([]byte{sig[64]})

	spew.Dump(r, s, v)
}

func TestSign(t *testing.T) {
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	pubKey := privKey.PublicKey
	pubKeyBytes := elliptic.Marshal(crypto.S256(), pubKey.X, pubKey.Y)
	pubKeyBytesHash := crypto.Keccak256(pubKeyBytes[1:])

	account := common.BytesToAddress(pubKeyBytesHash[12:])

	msg := []byte("the Web3Coach students are awesome")

	sig, err := Sign(msg, privKey)
	if err != nil {
		t.Fatal(err)
	}

	recoveredPubKey, err := Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}

	recoveredPubKeyBytes := elliptic.Marshal(crypto.S256(), recoveredPubKey.X, recoveredPubKey.Y)
	recoveredPubKeyBytesHash := crypto.Keccak256(recoveredPubKeyBytes[1:])
	recoveredAccount := common.BytesToAddress(recoveredPubKeyBytesHash[12:])

	if account.Hex() != recoveredAccount.Hex() {
		t.Fatalf("msg was signed by account %s but signature recovery produced an account %s", account.Hex(), recoveredAccount.Hex())
	}
}

func TestSignTxWithKeystoreAccount(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "wallet_test")
	if err != nil {
		t.Fatal(err)
	}
	defer fs.RemoveDir(tmpDir)

	andrej, err := NewKeystoreAccount(tmpDir, testKeystoreAccountsPwd)
	if err != nil {
		t.Error(err)
		return
	}

	babaYaga, err := NewKeystoreAccount(tmpDir, testKeystoreAccountsPwd)
	if err != nil {
		t.Error(err)
		return
	}

	tx := database.NewBaseTx(andrej, babaYaga, 100, 1, "")

	signedTx, err := SignTxWithKeystoreAccount(tx, andrej, testKeystoreAccountsPwd, GetKeystoreDirPath(tmpDir))
	if err != nil {
		t.Error(err)
		return
	}

	ok, err := signedTx.IsAuthentic()
	if err != nil {
		t.Error(err)
		return
	}

	if !ok {
		t.Error("the TX was signed by 'from' account and should have been authentic")
		return
	}

	// Test marshaling
	signedTxJson, err := json.Marshal(signedTx)
	if err != nil {
		t.Error(err)
		return
	}

	var signedTxUnmarshaled database.SignedTx
	err = json.Unmarshal(signedTxJson, &signedTxUnmarshaled)
	if err != nil {
		t.Error(err)
		return
	}

	require.Equal(t, signedTx, signedTxUnmarshaled)
}

func TestSignForgedTxWithKeystoreAccount(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "wallet_test")
	if err != nil {
		t.Fatal(err)
	}
	defer fs.RemoveDir(tmpDir)

	hacker, err := NewKeystoreAccount(tmpDir, testKeystoreAccountsPwd)
	if err != nil {
		t.Error(err)
		return
	}

	babaYaga, err := NewKeystoreAccount(tmpDir, testKeystoreAccountsPwd)
	if err != nil {
		t.Error(err)
		return
	}

	forgedTx := database.NewBaseTx(babaYaga, hacker, 100, 1, "")

	signedTx, err := SignTxWithKeystoreAccount(forgedTx, hacker, testKeystoreAccountsPwd, GetKeystoreDirPath(tmpDir))
	if err != nil {
		t.Error(err)
		return
	}

	ok, err := signedTx.IsAuthentic()
	if err != nil {
		t.Error(err)
		return
	}

	if ok {
		t.Fatal("the TX 'from' attribute was forged and should have not be authentic")
	}
}
