package gstext

//    gameText, err := UnmarshalGameText(bytes)
//    bytes, err = gameText.Marshal()

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func JSONUnmarshalGameText(data []byte) (GameText, error) {
	var r GameText
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GameText) JSONMarshal() ([]byte, error) {
	return json.Marshal(r)
}
func XMLUnmarshalGameText(data []byte) (GameText, error) {
	var r GameText
	err := xml.Unmarshal(data, &r)
	return r, err
}

func (r *GameText) XMLMarshal() ([]byte, error) {
	return xml.Marshal(r)
}

func YAMLUnmarshalGameText(data []byte) (GameText, error) {
	var r GameText
	err := yaml.Unmarshal(data, &r)
	return r, err
}

func (r *GameText) YAMLMarshal() ([]byte, error) {
	return yaml.Marshal(r)
}

func TOMLUnmarshalGameText(data []byte) (GameText, error) {
	var r GameText
	err := toml.Unmarshal(data, &r)
	return r, err
}

func (r *GameText) TOMLMarshal() ([]byte, error) {
	return toml.Marshal(r)
}
