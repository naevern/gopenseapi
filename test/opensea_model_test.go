package opensea

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func readFileAndUnmarshal(t *testing.T, filePath string, target interface{}) {
	inputFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %s", filePath)
	}

	err = json.Unmarshal(inputFile, target)
	if err != nil {
		assert.FailNow(t, "failed to unmarshal JSON: "+err.Error())
	}

	assert.NotNil(t, target, "unmarshalled object should not be nil")
}

func TestAssetResponse(t *testing.T) {
	osAsset := &AssetResponse{}
	readFileAndUnmarshal(t, "test-files/opensea-assets-collectibles.json", osAsset)
}

func TestStatResponse(t *testing.T) {
	osStat := &StatResponse{}
	readFileAndUnmarshal(t, "test-files/opensea-stats-doodles.json", osStat)
}

func TestSingleCollectionResponse(t *testing.T) {
	osColl := &CollectionSingleResponse{}
	readFileAndUnmarshal(t, "test-files/opensea-collection-doodles.json", osColl)
}
