package registry

import (
	"github.com/dataleodev/registry/pkg/errors"
	"regexp"
)

type Type int

const (
	Sensor Type = iota
	Actuator
	Controller
)

//IntToNodeType converts integer to Node Type
func IntToNodeType(value int) (nodeType Type, name string, err error) {
	if value == 0 {
		return Sensor, "sensor", nil
	} else if value == 1 {
		return Actuator, "actuator", nil
	} else if value == 2 {
		return Controller, "controller", err
	} else {
		err = errors.New("unrecognized value type: valid are 0,1 and 2")
		return Type(value), "", err
	}
}

var macAddrRegex = regexp.MustCompile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4})$")

func (node Node) ValidateMacAddr(macAddr string) bool {
	if len(macAddr) == 0 || macAddr == "" {
		return false
	}

	return macAddrRegex.MatchString(macAddr)
}

type Node struct {
	UUID    string `json:"uuid,omitempty"`
	Addr    string `json:"addr"`
	Key     string  `json:"key,omitempty"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Region  string `json:"region"`
	Latd    string `json:"latitude"`
	Long    string `json:"longitude"`
	Created string `json:"created,omitempty"`
	Master  string `json:"master,omitempty"`
}


//50759ce8-1148-4bd8-89e5-5dfbb5133f6c
//b6858384-b6d2-4cf7-b9ab-56a107e8efc8