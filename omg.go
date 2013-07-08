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
	return
}
func (r *MessageWithOffsetStruct) ToWire() (b []byte) {
	return
}
func (r *MessageSetStruct) ToWire() (b []byte) {
	return
}
func (r *MessageSetStructWithSize) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.size))
	return
}
func (r *MessagesForPartition) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	return
}
func (r *MessagesForTopic) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	return
}
func (r *ProduceRequest) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.RequiredAcks))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Timeout))
	return
}
func (r *ProducePartitionResponse) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	return
}
func (r *ProduceTopicResponse) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	return
}
func (r *ProduceResponse) ToWire() (b []byte) {
	return
}
func (r *PartitionFetch) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MaxBytes))
	return
}
func (r *TopicFetch) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	return
}
func (r *FetchRequest) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.ReplicaId))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MaxWaitTime))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MinBytes))
	return
}
func (r *PartitionFetchResponse) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	return
}
func (r *TopicFetchResponse) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	return
}
func (r *FetchResponse) ToWire() (b []byte) {
	return
}
func (r *PartitionOffsetRequest) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.MaxNumberOfOffsets))
	return
}
func (r *TopicOffsetRequest) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	return
}
func (r *OffsetRequest) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.ReplicaId))
	return
}
func (r *Offset) ToWire() (b []byte) {
	return
}
func (r *PartitionOffset) ToWire() (b []byte) {
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.Partition))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	return
}
func (r *TopicOffsetResponse) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Topic)))
	b = append(b, []byte(r.Topic)...)
	return
}
func (r *OffsetResponse) ToWire() (b []byte) {
	return
}
func (r *MetadataRequest) ToWire() (b []byte) {
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
	return
}
func (r *TopicMetadataStruct) ToWire() (b []byte) {
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.Error))
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.Name)))
	b = append(b, []byte(r.Name)...)
	return
}
func (r *MetadataResponse) ToWire() (b []byte) {
	return
}
