package mobile

import (
	"errors"

	"github.com/lomocoin/wallet-core/bip44"
)

// Clone make a copy of existing Wallet instance, with original attributes override by the given options
func (c Wallet) Clone(options ...WalletOption) (wallet *Wallet, err error) {
	cloned := c
	for _, opt := range options {
		err = opt(&cloned)
		if err != nil {
			return nil, err
		}
	}
	//TODO verify wallet
	return &cloned, nil
}

type WalletOption func(*Wallet) error

func WithShareAccountWithParentChain(shareAccountWithParentChain bool) WalletOption {
	return func(wallet *Wallet) error {
		wallet.ShareAccountWithParentChain = shareAccountWithParentChain
		return nil
	}
}

func WithPathFormat(pathFormat string) WalletOption {
	return func(wallet *Wallet) error {
		wallet.path = pathFormat
		return nil
	}
}

func WithPassword(password string) WalletOption {
	return func(wallet *Wallet) error {
		wallet.password = password
		return nil
	}
}

// NewWalletBuilder normal builder pattern, not so good
func NewWalletBuilder() *WalletBuilder {
	return &WalletBuilder{}
}

type WalletBuilder struct {
	mnemonic                    string
	shareAccountWithParentChain bool
	seed                        []byte
	testNet                     bool
	password                    string
	pathFormat                  string
}

func (wb *WalletBuilder) SetMnemonic(mnemonic string) *WalletBuilder {
	wb.mnemonic = mnemonic
	return wb
}

func (wb *WalletBuilder) SetTestNet(testNet bool) *WalletBuilder {
	wb.testNet = testNet
	return wb
}

func (wb *WalletBuilder) SetPassword(password string) *WalletBuilder {
	wb.password = password
	return wb
}

func (wb *WalletBuilder) SetShareAccountWithParentChain(shareAccountWithParentChain bool) *WalletBuilder {
	wb.shareAccountWithParentChain = shareAccountWithParentChain
	return wb
}

func (wb *WalletBuilder) SetUseShortestPath(useShortestPath bool) *WalletBuilder {
	var pathFormat string
	if useShortestPath {
		pathFormat = bip44.PathFormat
	} else {
		pathFormat = bip44.FullPathFormat
	}
	wb.pathFormat = pathFormat
	return wb
}

func (wb *WalletBuilder) Wallet() (wallet *Wallet, err error) {
	if wb.mnemonic == "" {
		return nil, errors.New("mnemonic should not be empty")
	}
	wallet, err = NewHDWalletFromMnemonic(wb.mnemonic, wb.testNet)
	if err != nil {
		return nil, err
	}
	wallet.path = wb.pathFormat
	wallet.ShareAccountWithParentChain = wb.shareAccountWithParentChain
	wallet.password = wb.password
	//TODO verify wallet
	return
}

// BuildWallet create a Wallet instance with fixed args (here is mnemonic and testNet) and other options
func BuildWalletFromMnemonic(mnemonic string, testNet bool, options ...WalletOption) (wallet *Wallet, err error) {
	wallet, err = NewHDWalletFromMnemonic(mnemonic, testNet)
	if err != nil {
		return
	}
	for _, opt := range options {
		err = opt(wallet)
		if err != nil {
			return
		}
	}
	//TODO verify wallet
	return
}

// TODO not implemented
// BuildWallet create a Wallet instance with fixed args (here is privateKey and testNet) and other options
func BuildWalletFromPrivateKey(privateKey string, testNet bool, options ...WalletOption) (wallet *Wallet, err error) {
	panic("implement me")
}
