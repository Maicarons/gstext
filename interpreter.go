package gstext

//    gameText, err := UnmarshalGameText(bytes)
//    bytes, err = gameText.Marshal()

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

type GameText struct {
	Scriptname    string `json:"scriptname" xml:"scriptname" yaml:"scriptname" toml:"scriptname"`
	Description   string `json:"description" xml:"description" yaml:"description" toml:"description"`
	Author        string `json:"author" xml:"author" yaml:"author" toml:"author"`
	Formatversion string `json:"formatversion" xml:"formatversion" yaml:"formatversion" toml:"formatversion"`
	Text          []Text `json:"text" xml:"text" yaml:"text" toml:"text"`
}

type Text struct {
	Type    string `json:"type" xml:"type" yaml:"type" toml:"type"`
	Meta    string `json:"meta" xml:"meta" yaml:"meta" toml:"meta"`
	Remark  string `json:"remark" xml:"remark" yaml:"remark" toml:"remark"`
	Content string `json:"content" xml:"content" yaml:"content" toml:"content"`
}

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
