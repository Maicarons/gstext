package gstext

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
