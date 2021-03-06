package rtm

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToEvent(t *testing.T) {
	tests := []struct {
		in  string
		out Event
	}{
		{
			`{"type":"butt"}`,
			BasicEvent{"butt", map[string]interface{}{"type": "butt"}},
		},
		{`{"type":"hello"}`, Hello{}},
		{
			`{
				"type": "message",
				"channel": "C2147483705",
				"user": "U2147483697",
				"text": "Hello world",
				"ts": "1355517523.000005"
			}`,
			BasicMessage{
				channel:   "C2147483705",
				user:      "U2147483697",
				text:      "Hello world",
				timestamp: "1355517523.000005",
			},
		},
		{
			`{
				"type":"channel_created",
				"channel": {
					"id": "C024BE91L",
					"name": "fun",
					"created": 1360782804,
					"creator": "U024BE7LH"
				}
			}`,
			ChannelCreated{
				ID:      "C024BE91L",
				Name:    "fun",
				Created: 1360782804,
				Creator: "U024BE7LH",
			},
		},
		{
			`{
				"type": "channel_rename",
				"channel": {
					"id":"C02ELGNBH",
					"name":"new_name",
					"created":1360782804
				}
			}`,
			ChannelRename{
				ID:      "C02ELGNBH",
				Name:    "new_name",
				Created: 1360782804,
			},
		},
	}

	for _, test := range tests {
		in := make(map[string]interface{})
		json.Unmarshal([]byte(test.in), &in)

		out := toEvent(in)
		assert.Equal(t, test.out, out, "not the right event?")
	}
}
