package marshaller

import (
	"encoding/json"
	"testing"
)

func TestContainer(t *testing.T) {

	var toMap map[string]interface{}
	var result map[string]interface{}

	// handle bussiness here

	conv := ConventionalMarshaller{Value: toMap}
	b, err := conv.MarshalJSON()
	if err != nil {
		// handle error here
	}
	json.Unmarshal(b, &result)
}
