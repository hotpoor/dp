package scryclient

import (
    "../core/chainevents"
    "../core/chainoperations"
    "../util/accounts"
    "./chaininterfacewrapper"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    rlog "github.com/sirupsen/logrus"
    "math/big"
)

type ScryClient struct {
	Account *accounts.Account
}

func NewScryClient(publicKey string) (*ScryClient, error) {
    return &ScryClient{
		Account: &accounts.Account{publicKey},
	}, nil
}

func CreateScryClient(password string) (*ScryClient, error) {
    account, err := accounts.GetAMInstance().CreateAccount(password)
    if err != nil {
        rlog.Error("failed to create Account, error:", err)
        return nil, err
    }

    return &ScryClient{
        Account: account,
    }, nil
}

func (client ScryClient) SubscribeEvent(eventName string, callback chainevents.EventCallback)  {
	chainevents.SubscribeExternal(common.HexToAddress(client.Account.Address), eventName, callback)
}

func (client ScryClient) Authenticate(password string) (bool, error) {
    return accounts.GetAMInstance().AuthAccount(client.Account.Address, password)
}

func (client ScryClient) TransferEthFrom(from common.Address, password string, value *big.Int, ec *ethclient.Client) (error) {
    tx, err := chainoperations.TransferEth(from, password, common.HexToAddress(client.Account.Address), value, ec)
    if err == nil {
        rlog.Debug("transferEthFrom: ", tx.Hash(), tx.Data())
    }

    return err
}

func (client ScryClient) TransferTokenFrom(from common.Address, password string, value *big.Int) (error) {
    txParam := &chainoperations.TransactParams{From: from, Password: password, Value:value}
    return chaininterfacewrapper.TransferTokens(txParam,
                                                    common.HexToAddress(client.Account.Address),
                                                    value)
}

func (client ScryClient) GetEth(owner common.Address, ec *ethclient.Client) (*big.Int, error) {
    return chainoperations.GetEthBalance(owner, ec)
}

func (client ScryClient) GetScryToken(owner common.Address) (*big.Int, error)  {
    from := common.HexToAddress(client.Account.Address)
    txParam := &chainoperations.TransactParams{From: from, Pending: true}

    return chaininterfacewrapper.GetTokenBalance(txParam, owner)
}