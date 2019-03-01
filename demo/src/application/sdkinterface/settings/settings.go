package settings

import (
    "errors"
    cf "github.com/scryinfo/iscap/demo/src/sdk/util/configuration"
    rlog "github.com/sirupsen/logrus"
)

const (
    SETTING_LOCATION = "./definition.yaml"
)

func LoadSettings() (*ScryInfo, error) {
    rv, err := cf.GetYAMLStructure(SETTING_LOCATION, &ScryInfo{})
    if err != nil {
        rlog.Error("failed to get yaml structure, error:", err)
        return nil, err
    }

    scryinfo, ok := rv.(*ScryInfo)
    if !ok {
        emsg := "failed to convert stream to yaml structure"
        rlog.Error(emsg)
        return nil, errors.New(emsg)
    }

    return scryinfo, nil
}