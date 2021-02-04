package coinbasepro

import (
	"fmt"
)

type WithdrawalCrypto struct {
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	CryptoAddress string `json:"crypto_address"`
}

type WithdrawalCoinbase struct {
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id"`
}

type WithdrawlCryptoFeeEstimate struct {
	Currency      string `json:"currency"`
	CryptoAddress string `json:"crypto_address"`
	// Response fields
	Fee string `json:"fee,omitempty"`
}

func (c *Client) CreateWithdrawalCrypto(newWithdrawalCrypto *WithdrawalCrypto) (WithdrawalCrypto, error) {
	var savedWithdrawal WithdrawalCrypto
	url := fmt.Sprintf("/withdrawals/crypto")
	_, err := c.Request("POST", url, newWithdrawalCrypto, &savedWithdrawal)
	return savedWithdrawal, err
}

func (c *Client) CreateWithdrawalCoinbase(newWithdrawalCoinbase *WithdrawalCoinbase) (WithdrawalCoinbase, error) {
	var savedWithdrawal WithdrawalCoinbase
	url := fmt.Sprintf("/withdrawals/coinbase-account")
	_, err := c.Request("POST", url, newWithdrawalCoinbase, &savedWithdrawal)
	return savedWithdrawal, err
}

func (c *Client) GetFeeEstimate(newCryptoFeeEstimate *WithdrawlCryptoFeeEstimate) (WithdrawlCryptoFeeEstimate, error) {
	var savedFeeEstimate WithdrawlCryptoFeeEstimate
	url := fmt.Sprintf("/withdrawals/fee-estimate")
	_, err := c.Request("GET", url, newCryptoFeeEstimate, &savedFeeEstimate)
	return savedFeeEstimate, err
}
