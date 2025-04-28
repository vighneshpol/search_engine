//  For data models (structs)
package model

type Record struct {
	MsgId          string
	PartitionId    int32
	Timestamp      int64
	Hostname       string
	Priority       int32
	Facility       int32
	FacilityString string
	Severity       int32
	SeverityString string
	AppName        string
	ProcId         string
	Message        string
	MessageRaw     string
	StructuredData string
	Tag            string
	Sender         string
	Groupings      string
	Event          string
	EventId        string
	NanoTimeStamp  int64
	Namespace      string
}
