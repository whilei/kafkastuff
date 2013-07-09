package main

import "testing"
import "fmt"

func TestStuff(t *testing.T) {
	x := &ProduceRequest{
		Common: RequestCommon{
			ApiKey:        0,
			ApiVersion:    1,
			CorrelationId: 2,
			ClientId:      "client_id",
		},
		RequiredAcks: 1,
		Timeout:      1000,
		MessagesForTopics: []MessagesForTopic{
			{
				Topic: "test",
				MessagesForPartitions: []MessagesForPartition{
					{
						Partition: 0,
						MessageSet: MessageSet{Messages: []Message{
							{
								crc32:      0,
								size:       0,
								MagicType:  0,
								Attributes: 0,
								Key:        []byte("lol"),
								Value:      []byte("wtf"),
							},
						}},
					},
				},
			},
		},
	}
	fmt.Println(x.ToWire())
}
