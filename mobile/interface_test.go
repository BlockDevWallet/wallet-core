package mobile

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	wallet       = new(Wallet)
	testMnemonic = "lecture leg select like delay limit spread retire toward west grape bachelor"
)

func init() {
	wallet, _ = NewHDWalletFromMnemonic(testMnemonic, false)
}

func TestCoin_DeriveAddress(t *testing.T) {
	addr, err := wallet.DeriveAddress("ETH")
	assert.NoError(t, err)
	assert.Equal(t, "0x947ab281Df5ec46E801F78Ad1363FaaCbe4bfd12", addr)
	t.Log(addr)

	pk, err := wallet.DerivePrivateKey("WCG")
	t.Log(pk)
	addr, err = wallet.DeriveAddress("WCG")
	assert.NoError(t, err)
	assert.Equal(t, "WCG-3NHD-TWLP-KNDL-9PPKW", addr)
	t.Log(addr)

	addr, err = wallet.DeriveAddress("USDTK")
	assert.NoError(t, err)
	assert.Equal(t, "WCG-3NHD-TWLP-KNDL-9PPKW", addr)
	t.Log(addr)

	pk, err = wallet.DerivePrivateKey("NXT")
	t.Log(pk)
	addr, err = wallet.DeriveAddress("NXT")
	assert.NoError(t, err)
	assert.Equal(t, "NXT-MCMS-T636-MQ4G-4LSWJ", addr)
	t.Log(addr)

	addr, err = wallet.DeriveAddress("RMB")
	assert.NoError(t, err)
	assert.Equal(t, "NXT-MCMS-T636-MQ4G-4LSWJ", addr)
	t.Log(addr)

	addr, err = wallet.DeriveAddress("BTC")
	assert.NoError(t, err)
	assert.Equal(t, "13vvVPKZjsStYRZft3RyfgmCVVFsYm8nDT", addr)
	t.Log(addr)

	wallet.ShareAccountWithParentChain = false
	addr, err = wallet.DeriveAddress("USDT(Omni)")
	assert.NoError(t, err)
	assert.NotEqual(t, "13vvVPKZjsStYRZft3RyfgmCVVFsYm8nDT", addr)
	assert.Equal(t, "1AzTauTdhZ4VKC88MAb7iu9jU3yNzpx937", addr)
	t.Log(addr)

	wallet.ShareAccountWithParentChain = true
	addr, err = wallet.DeriveAddress("USDT(Omni)")
	assert.NoError(t, err)
	assert.NotEqual(t, "1AzTauTdhZ4VKC88MAb7iu9jU3yNzpx937", addr)
	assert.Equal(t, "13vvVPKZjsStYRZft3RyfgmCVVFsYm8nDT", addr)
	t.Log(addr)
	wallet.ShareAccountWithParentChain = false
}

func TestWallet_GetAvailableCoinList(t *testing.T) {
	bb := GetAvailableCoinList()
	t.Log(bb)
	cc := strings.Split(bb, " ")
	for i := range cc {
		addr, err := wallet.DeriveAddress(cc[i])
		assert.NoError(t, err)
		t.Log(cc[i], addr)
	}
}

func TestNewMnemonic(t *testing.T) {
	mn, err := NewMnemonic()
	assert.NoError(t, err)
	en, err := EntropyFromMnemonic(mn)
	assert.NoError(t, err)
	mn1, err := MnemonicFromEntropy(en)
	assert.NoError(t, err)
	assert.EqualValues(t, mn, mn1)
}

func TestGetVersion(t *testing.T) {
	t.Log(GetVersion())
	t.Log(GetBuildTime())
	t.Log(GetGitHash())
}
