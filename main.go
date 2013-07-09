package main

import (
	"fmt"
	"reflect"
)

// grep -e "^type " types.go | awk '{print "  compile(reflect.TypeOf(" $2 "{}))"}'
func main() {
	fmt.Println("package main")
	fmt.Println("")
	fmt.Println("import \"encoding/binary\"")
	fmt.Println("")
	compile(reflect.TypeOf(Message{}))
	compile(reflect.TypeOf(MessageSet{}))
	compile(reflect.TypeOf(TopicMetadata{}))
	compile(reflect.TypeOf(RequestCommon{}))
	compile(reflect.TypeOf(ResponseCommon{}))
	compile(reflect.TypeOf(MessageWithOffsetStruct{}))
	compile(reflect.TypeOf(MessageSetStructWithSize{}))
	compile(reflect.TypeOf(MessagesForPartition{}))
	compile(reflect.TypeOf(MessagesForTopic{}))
	compile(reflect.TypeOf(ProduceRequest{}))
	compile(reflect.TypeOf(ProducePartitionResponse{}))
	compile(reflect.TypeOf(ProduceTopicResponse{}))
	compile(reflect.TypeOf(ProduceResponse{}))
	compile(reflect.TypeOf(PartitionFetch{}))
	compile(reflect.TypeOf(TopicFetch{}))
	compile(reflect.TypeOf(FetchRequest{}))
	compile(reflect.TypeOf(PartitionFetchResponse{}))
	compile(reflect.TypeOf(TopicFetchResponse{}))
	compile(reflect.TypeOf(FetchResponse{}))
	compile(reflect.TypeOf(PartitionOffsetRequest{}))
	compile(reflect.TypeOf(TopicOffsetRequest{}))
	compile(reflect.TypeOf(OffsetRequest{}))
	compile(reflect.TypeOf(Offset{}))
	compile(reflect.TypeOf(PartitionOffset{}))
	compile(reflect.TypeOf(TopicOffsetResponse{}))
	compile(reflect.TypeOf(OffsetResponse{}))
	compile(reflect.TypeOf(MetadataRequest{}))
	compile(reflect.TypeOf(Broker{}))
	compile(reflect.TypeOf(PartitionMetadata{}))
	compile(reflect.TypeOf(TopicMetadataStruct{}))
	compile(reflect.TypeOf(MetadataResponse{}))
}
