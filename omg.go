package main

import "encoding/binary"

func (r *Message) ToWire() (b []byte) {
	return
}
func (r *MessageSet) ToWire() (b []byte) {
	return
}
func (r *TopicMetadata) ToWire() (b []byte) {
	return
}
func (r *RequestCommon) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.ApiKey))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.ApiVersion))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.CorrelationId))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.ClientId)))
	b = append(b, []byte(r.ClientId)...)
	return
}
func (r *ResponseCommon) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.CorrelationId))
	return
}
func (r *MessageStruct) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.crc32))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.size))
	b = append(b, byte(r.MagicType))
	b = append(b, byte(r.Attributes))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Key)))
	b = append(b, r.Key...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Value)))
	b = append(b, r.Value...)
	return
}
func (r *MessageWithOffsetStruct) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(b[len(b)-8:], uint64(r.Offset))
	b = append(b, []byte(r.Message.ToWire())...)
	return
}
func (r *MessageSetStruct) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Messages)))
	for _, el := range r.Messages {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *MessageSetStructWithSize) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.size))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Messages)))
	for _, el := range r.Messages {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *MessagesForPartition) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, []byte(r.MessageSet.ToWire())...)
	return
}
func (r *MessagesForTopic) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.MessagesForPartitions)))
	for _, el := range r.MessagesForPartitions {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *ProduceRequest) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.RequiredAcks))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Timeout))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.MessagesForTopics)))
	for _, el := range r.MessagesForTopics {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *ProducePartitionResponse) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(b[len(b)-8:], uint64(r.Offset))
	return
}
func (r *ProduceTopicResponse) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Partitions)))
	for _, el := range r.Partitions {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *ProduceResponse) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.TopicResponse)))
	for _, el := range r.TopicResponse {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *PartitionFetch) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(b[len(b)-8:], uint64(r.FetchOffset))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MaxBytes))
	return
}
func (r *TopicFetch) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.PartitionFetches)))
	for _, el := range r.PartitionFetches {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *FetchRequest) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.ReplicaId))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MaxWaitTime))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MinBytes))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.TopicFetches)))
	for _, el := range r.TopicFetches {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *PartitionFetchResponse) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(b[len(b)-8:], uint64(r.HighwaterMarkOffset))
	b = append(b, []byte(r.MessageSet.ToWire())...)
	return
}
func (r *TopicFetchResponse) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.PartitionFetchResponses)))
	for _, el := range r.PartitionFetchResponses {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *FetchResponse) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.TopicFetchResponses)))
	for _, el := range r.TopicFetchResponses {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *PartitionOffsetRequest) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(b[len(b)-8:], uint64(r.Time))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MaxNumberOfOffsets))
	return
}
func (r *TopicOffsetRequest) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.PartitionOffsetRequests)))
	for _, el := range r.PartitionOffsetRequests {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *OffsetRequest) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.ReplicaId))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.TopicOffsetReqeusts)))
	for _, el := range r.TopicOffsetReqeusts {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *Offset) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(b[len(b)-8:], uint64(r.Offset))
	return
}
func (r *PartitionOffset) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Offsets)))
	for _, el := range r.Offsets {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *TopicOffsetResponse) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.PartitionOffsets)))
	for _, el := range r.PartitionOffsets {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *OffsetResponse) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.TopicOffsetResponses)))
	for _, el := range r.TopicOffsetResponses {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *MetadataRequest) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.TopicNames)))
	for _, el := range r.TopicNames {
		b = append(b, 0, 0)
		binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(el)))
		b = append(b, []byte(el)...)
	}
	return
}
func (r *Broker) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Id))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Host)))
	b = append(b, []byte(r.Host)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Port))
	return
}
func (r *PartitionMetadata) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Id))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Leader))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Replicas)))
	for _, el := range r.Replicas {
		b = append(b, 0, 0, 0, 0)
		binary.BigEndian.PutUint32(b[len(b)-4:], uint32(el))
	}
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Isr)))
	for _, el := range r.Isr {
		b = append(b, 0, 0, 0, 0)
		binary.BigEndian.PutUint32(b[len(b)-4:], uint32(el))
	}
	return
}
func (r *TopicMetadataStruct) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Name)))
	b = append(b, []byte(r.Name)...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Partitions)))
	for _, el := range r.Partitions {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
func (r *MetadataResponse) ToWire() (b []byte) {
	b = append(b, []byte(r.Common.ToWire())...)
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Brokers)))
	for _, el := range r.Brokers {
		b = append(b, []byte(el.ToWire())...)
	}
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.Topics)))
	for _, el := range r.Topics {
		b = append(b, []byte(el.ToWire())...)
	}
	return
}
