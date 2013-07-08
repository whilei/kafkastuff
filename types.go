package main

// All protocol messages enumerated here:
// https://github.com/bpot/poseidon/blob/master/lib/poseidon/protocol.rb

// These are implemented as classes external to protocol.rb in Poseidon.
// Extra care will probably be required in implementing these in particular.
type Message struct{}
type MessageSet struct{}
type TopicMetadata struct{}

// REQUEST/RESPONSE COMMON STRUCTURE/////////////////////////
type RequestCommon struct {
	ApiKey        int16
	ApiVersion    int16
	CorrelationId int32
	ClientId      string
}

type ResponseCommon struct {
	CorrelationId int32
}

// MESSAGESET COMMON STRUCTURE/////////////////////////
type MessageStruct struct {
	crc32      int32
	size       int32
	MagicType  int8
	Attributes int8
	Key        []byte
	Value      []byte
}

type MessageWithOffsetStruct struct {
	Offset  int64
	Message MessageStruct
}

type MessageSetStruct struct {
	Messages []Message `size_bounded`
}

type MessageSetStructWithSize struct {
	size     int32
	Messages []Message `size_bounded`
}

// PRODUCE REQUEST //////////////////////////////////////////
type MessagesForPartition struct {
	Partition int32
	MessageSet
}

type MessagesForTopic struct {
	Topic                 string
	MessagesForPartitions []MessagesForPartition
}

type ProduceRequest struct {
	Common            RequestCommon
	RequiredAcks      int16
	Timeout           int32
	MessagesForTopics []MessagesForTopic
}

// PRODUCE RESPONSE //////////////////////////////////////////
type ProducePartitionResponse struct {
	Partition int32
	Error     int16
	Offset    int64
}

type ProduceTopicResponse struct {
	Topic      string
	Partitions []ProducePartitionResponse
}

type ProduceResponse struct {
	Common        ResponseCommon
	TopicResponse []ProduceTopicResponse
}

// FETCH REQUEST //////////////////////////////////////////
type PartitionFetch struct {
	Partition   int32
	FetchOffset int64
	MaxBytes    int32
}

type TopicFetch struct {
	Topic            string
	PartitionFetches []PartitionFetch
}

type FetchRequest struct {
	Common       RequestCommon
	ReplicaId    int32
	MaxWaitTime  int32
	MinBytes     int32
	TopicFetches []TopicFetch
}

// FETCH RESPONSE //////////////////////////////////////////

type PartitionFetchResponse struct {
	Partition           int32
	Error               int16
	HighwaterMarkOffset int64
	MessageSet
}

type TopicFetchResponse struct {
	Topic                   string
	PartitionFetchResponses []PartitionFetchResponse
}

type FetchResponse struct {
	Common              ResponseCommon
	TopicFetchResponses []TopicFetchResponse
}

// OFFSET REQUEST //////////////////////////////////////////
type PartitionOffsetRequest struct {
	Partition          int32
	Time               int64
	MaxNumberOfOffsets int32
}

type TopicOffsetRequest struct {
	Topic                   string
	PartitionOffsetRequests []PartitionOffsetRequest
}

type OffsetRequest struct {
	Common              RequestCommon
	ReplicaId           int32
	TopicOffsetReqeusts []TopicOffsetRequest
}

// OFFSET RESPONSE //////////////////////////////////////////
type Offset struct {
	Offset int64
}

type PartitionOffset struct {
	Partition int32
	Error     int16
	Offsets   []Offset
}

type TopicOffsetResponse struct {
	Topic            string
	PartitionOffsets []PartitionOffset
}

type OffsetResponse struct {
	Common               ResponseCommon
	TopicOffsetResponses []TopicOffsetResponse
}

// METADATA REQUEST //////////////////////////////////////////
type MetadataRequest struct {
	Common     RequestCommon
	TopicNames []string
}

// METADATA RESPONSE //////////////////////////////////////////
type Broker struct {
	Id   int32
	Host string
	Port int32
}

type PartitionMetadata struct {
	Error    int16
	Id       int32
	Leader   int32
	Replicas []int32
	Isr      []int32
}

type TopicMetadataStruct struct {
	Error      int16
	Name       string
	Partitions []PartitionMetadata
}

type MetadataResponse struct {
	Common  ResponseCommon
	Brokers []Broker
	Topics  []TopicMetadata
}
