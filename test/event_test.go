package opensea

import (
	"testing"
	"time"
)

func TestRetrieveEvents(t *testing.T) {
	assert := initializeTest(t)

	params := NewRetrievingEventsParams()
	if err := params.SetAssetContractAddress(contract); err != nil {
		assert.Nil(err)
	}

	currentTime := time.Now().Unix()
	params.OccurredAfter = currentTime - 86400 // 24 hours ago
	params.OccurredBefore = currentTime
	params.EventType = EventTypeCreated

	events, err := o.RetrievingEvents(params)
	assert.Nil(err)

	t.Logf("Number of events retrieved: %d", len(events))
}
