package standard

import (
	"encoding/json"
	"github.com/Maicarons/gstext"
)

func JSONUnmarshalGameText(data []byte) (gstext.GameText, error) {
	var r gstext.GameText
	err := json.Unmarshal(data, &r)
	return r, err
}
