package sdkinterface

import (
    "errors"
    "fmt"
    "github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
    "github.com/scryinfo/iscap/demo/src/sdk"
    "github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
    "github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
    "io/ioutil"
    "strings"
)

var (
    scryInfo *settings.ScryInfo        = nil
    failedToInitSDK = "failed to initialize sdk. "
    sep = "|"
)

func Initialize() error {
    // load definition
    scryInfo, err := settings.LoadSettings()
    if err != nil {
        fmt.Println(failedToInitSDK, err)
        return errors.New(failedToInitSDK)
    }

    // initialization
    contracts := getContracts(scryInfo.Chain.Contracts.ProtocolAddr,
        scryInfo.Chain.Contracts.TokenAddr,
        scryInfo.Chain.Contracts.ProtocolAbiPath,
        scryInfo.Chain.Contracts.TokenAbiPath,
        scryInfo.Chain.Contracts.ProtocolEvents,
        scryInfo.Chain.Contracts.TokenEvents)

    err = sdk.Init(scryInfo.Chain.Ethereum.EthNode,
        scryInfo.Services.Keystore,
        contracts,
        0,
        scryInfo.Services.Ipfs)
    if err != nil {
        fmt.Println(failedToInitSDK, err)
        return errors.New(failedToInitSDK)
    }

    return nil
}

func getContracts(protocolContractAddr string,
                  tokenContractAddr string,
                  protocolAbiPath string,
                  tokenAbiPath string,
                  protocolEvents string,
                  tokenEvents string) []chainevents.ContractInfo {
    pe := strings.Split(protocolEvents, sep)
    te := strings.Split(tokenEvents, sep)

	contracts := []chainevents.ContractInfo{
		{protocolContractAddr, getAbiText(protocolAbiPath), pe},
		{tokenContractAddr, getAbiText(tokenAbiPath), te},
	}

	return contracts
}

func getAbiText(fileName string) string {
	abi, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("failed to read abi text", err)
		return ""
	}

	return string(abi)
}

func onPublish(event events.Event) bool {
	//if err := bootstrap.SendMessage(w, "onPublish", "Publish event callback from go"); err != nil {
	//	astilog.Error(errors.Wrap(err, "sending onPublish event failed"))
	//}
	return true
}