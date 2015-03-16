package kinesis

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// AddTagsToStreamRequest generates a request for the AddTagsToStream operation.
func (c *Kinesis) AddTagsToStreamRequest(input *AddTagsToStreamInput) (req *aws.Request, output *AddTagsToStreamOutput) {
	if opAddTagsToStream == nil {
		opAddTagsToStream = &aws.Operation{
			Name:       "AddTagsToStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAddTagsToStream, input, output)
	output = &AddTagsToStreamOutput{}
	req.Data = output
	return
}

func (c *Kinesis) AddTagsToStream(input *AddTagsToStreamInput) (output *AddTagsToStreamOutput, err error) {
	req, out := c.AddTagsToStreamRequest(input)
	output = out
	err = req.Send()
	return
}

var opAddTagsToStream *aws.Operation

// CreateStreamRequest generates a request for the CreateStream operation.
func (c *Kinesis) CreateStreamRequest(input *CreateStreamInput) (req *aws.Request, output *CreateStreamOutput) {
	if opCreateStream == nil {
		opCreateStream = &aws.Operation{
			Name:       "CreateStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateStream, input, output)
	output = &CreateStreamOutput{}
	req.Data = output
	return
}

func (c *Kinesis) CreateStream(input *CreateStreamInput) (output *CreateStreamOutput, err error) {
	req, out := c.CreateStreamRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateStream *aws.Operation

// DeleteStreamRequest generates a request for the DeleteStream operation.
func (c *Kinesis) DeleteStreamRequest(input *DeleteStreamInput) (req *aws.Request, output *DeleteStreamOutput) {
	if opDeleteStream == nil {
		opDeleteStream = &aws.Operation{
			Name:       "DeleteStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteStream, input, output)
	output = &DeleteStreamOutput{}
	req.Data = output
	return
}

func (c *Kinesis) DeleteStream(input *DeleteStreamInput) (output *DeleteStreamOutput, err error) {
	req, out := c.DeleteStreamRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteStream *aws.Operation

// DescribeStreamRequest generates a request for the DescribeStream operation.
func (c *Kinesis) DescribeStreamRequest(input *DescribeStreamInput) (req *aws.Request, output *DescribeStreamOutput) {
	if opDescribeStream == nil {
		opDescribeStream = &aws.Operation{
			Name:       "DescribeStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeStream, input, output)
	output = &DescribeStreamOutput{}
	req.Data = output
	return
}

func (c *Kinesis) DescribeStream(input *DescribeStreamInput) (output *DescribeStreamOutput, err error) {
	req, out := c.DescribeStreamRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeStream *aws.Operation

// GetRecordsRequest generates a request for the GetRecords operation.
func (c *Kinesis) GetRecordsRequest(input *GetRecordsInput) (req *aws.Request, output *GetRecordsOutput) {
	if opGetRecords == nil {
		opGetRecords = &aws.Operation{
			Name:       "GetRecords",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetRecords, input, output)
	output = &GetRecordsOutput{}
	req.Data = output
	return
}

func (c *Kinesis) GetRecords(input *GetRecordsInput) (output *GetRecordsOutput, err error) {
	req, out := c.GetRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetRecords *aws.Operation

// GetShardIteratorRequest generates a request for the GetShardIterator operation.
func (c *Kinesis) GetShardIteratorRequest(input *GetShardIteratorInput) (req *aws.Request, output *GetShardIteratorOutput) {
	if opGetShardIterator == nil {
		opGetShardIterator = &aws.Operation{
			Name:       "GetShardIterator",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetShardIterator, input, output)
	output = &GetShardIteratorOutput{}
	req.Data = output
	return
}

func (c *Kinesis) GetShardIterator(input *GetShardIteratorInput) (output *GetShardIteratorOutput, err error) {
	req, out := c.GetShardIteratorRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetShardIterator *aws.Operation

// ListStreamsRequest generates a request for the ListStreams operation.
func (c *Kinesis) ListStreamsRequest(input *ListStreamsInput) (req *aws.Request, output *ListStreamsOutput) {
	if opListStreams == nil {
		opListStreams = &aws.Operation{
			Name:       "ListStreams",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListStreams, input, output)
	output = &ListStreamsOutput{}
	req.Data = output
	return
}

func (c *Kinesis) ListStreams(input *ListStreamsInput) (output *ListStreamsOutput, err error) {
	req, out := c.ListStreamsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListStreams *aws.Operation

// ListTagsForStreamRequest generates a request for the ListTagsForStream operation.
func (c *Kinesis) ListTagsForStreamRequest(input *ListTagsForStreamInput) (req *aws.Request, output *ListTagsForStreamOutput) {
	if opListTagsForStream == nil {
		opListTagsForStream = &aws.Operation{
			Name:       "ListTagsForStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListTagsForStream, input, output)
	output = &ListTagsForStreamOutput{}
	req.Data = output
	return
}

func (c *Kinesis) ListTagsForStream(input *ListTagsForStreamInput) (output *ListTagsForStreamOutput, err error) {
	req, out := c.ListTagsForStreamRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTagsForStream *aws.Operation

// MergeShardsRequest generates a request for the MergeShards operation.
func (c *Kinesis) MergeShardsRequest(input *MergeShardsInput) (req *aws.Request, output *MergeShardsOutput) {
	if opMergeShards == nil {
		opMergeShards = &aws.Operation{
			Name:       "MergeShards",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opMergeShards, input, output)
	output = &MergeShardsOutput{}
	req.Data = output
	return
}

func (c *Kinesis) MergeShards(input *MergeShardsInput) (output *MergeShardsOutput, err error) {
	req, out := c.MergeShardsRequest(input)
	output = out
	err = req.Send()
	return
}

var opMergeShards *aws.Operation

// PutRecordRequest generates a request for the PutRecord operation.
func (c *Kinesis) PutRecordRequest(input *PutRecordInput) (req *aws.Request, output *PutRecordOutput) {
	if opPutRecord == nil {
		opPutRecord = &aws.Operation{
			Name:       "PutRecord",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutRecord, input, output)
	output = &PutRecordOutput{}
	req.Data = output
	return
}

func (c *Kinesis) PutRecord(input *PutRecordInput) (output *PutRecordOutput, err error) {
	req, out := c.PutRecordRequest(input)
	output = out
	err = req.Send()
	return
}

var opPutRecord *aws.Operation

// PutRecordsRequest generates a request for the PutRecords operation.
func (c *Kinesis) PutRecordsRequest(input *PutRecordsInput) (req *aws.Request, output *PutRecordsOutput) {
	if opPutRecords == nil {
		opPutRecords = &aws.Operation{
			Name:       "PutRecords",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutRecords, input, output)
	output = &PutRecordsOutput{}
	req.Data = output
	return
}

func (c *Kinesis) PutRecords(input *PutRecordsInput) (output *PutRecordsOutput, err error) {
	req, out := c.PutRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opPutRecords *aws.Operation

// RemoveTagsFromStreamRequest generates a request for the RemoveTagsFromStream operation.
func (c *Kinesis) RemoveTagsFromStreamRequest(input *RemoveTagsFromStreamInput) (req *aws.Request, output *RemoveTagsFromStreamOutput) {
	if opRemoveTagsFromStream == nil {
		opRemoveTagsFromStream = &aws.Operation{
			Name:       "RemoveTagsFromStream",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRemoveTagsFromStream, input, output)
	output = &RemoveTagsFromStreamOutput{}
	req.Data = output
	return
}

func (c *Kinesis) RemoveTagsFromStream(input *RemoveTagsFromStreamInput) (output *RemoveTagsFromStreamOutput, err error) {
	req, out := c.RemoveTagsFromStreamRequest(input)
	output = out
	err = req.Send()
	return
}

var opRemoveTagsFromStream *aws.Operation

// SplitShardRequest generates a request for the SplitShard operation.
func (c *Kinesis) SplitShardRequest(input *SplitShardInput) (req *aws.Request, output *SplitShardOutput) {
	if opSplitShard == nil {
		opSplitShard = &aws.Operation{
			Name:       "SplitShard",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSplitShard, input, output)
	output = &SplitShardOutput{}
	req.Data = output
	return
}

func (c *Kinesis) SplitShard(input *SplitShardInput) (output *SplitShardOutput, err error) {
	req, out := c.SplitShardRequest(input)
	output = out
	err = req.Send()
	return
}

var opSplitShard *aws.Operation

type AddTagsToStreamInput struct {
	StreamName *string             `type:"string" required:"true"json:",omitempty"`
	Tags       *map[string]*string `type:"map" required:"true"json:",omitempty"`

	metadataAddTagsToStreamInput `json:"-", xml:"-"`
}

type metadataAddTagsToStreamInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AddTagsToStreamOutput struct {
	metadataAddTagsToStreamOutput `json:"-", xml:"-"`
}

type metadataAddTagsToStreamOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CreateStreamInput struct {
	ShardCount *int64  `type:"integer" required:"true"json:",omitempty"`
	StreamName *string `type:"string" required:"true"json:",omitempty"`

	metadataCreateStreamInput `json:"-", xml:"-"`
}

type metadataCreateStreamInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CreateStreamOutput struct {
	metadataCreateStreamOutput `json:"-", xml:"-"`
}

type metadataCreateStreamOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteStreamInput struct {
	StreamName *string `type:"string" required:"true"json:",omitempty"`

	metadataDeleteStreamInput `json:"-", xml:"-"`
}

type metadataDeleteStreamInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteStreamOutput struct {
	metadataDeleteStreamOutput `json:"-", xml:"-"`
}

type metadataDeleteStreamOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeStreamInput struct {
	ExclusiveStartShardID *string `locationName:"ExclusiveStartShardId" type:"string" json:"ExclusiveStartShardId,omitempty"`
	Limit                 *int64  `type:"integer" json:",omitempty"`
	StreamName            *string `type:"string" required:"true"json:",omitempty"`

	metadataDescribeStreamInput `json:"-", xml:"-"`
}

type metadataDescribeStreamInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeStreamOutput struct {
	StreamDescription *StreamDescription `type:"structure" required:"true"json:",omitempty"`

	metadataDescribeStreamOutput `json:"-", xml:"-"`
}

type metadataDescribeStreamOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetRecordsInput struct {
	Limit         *int64  `type:"integer" json:",omitempty"`
	ShardIterator *string `type:"string" required:"true"json:",omitempty"`

	metadataGetRecordsInput `json:"-", xml:"-"`
}

type metadataGetRecordsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetRecordsOutput struct {
	NextShardIterator *string   `type:"string" json:",omitempty"`
	Records           []*Record `type:"list" required:"true"json:",omitempty"`

	metadataGetRecordsOutput `json:"-", xml:"-"`
}

type metadataGetRecordsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetShardIteratorInput struct {
	ShardID                *string `locationName:"ShardId" type:"string" required:"true"json:"ShardId,omitempty"`
	ShardIteratorType      *string `type:"string" required:"true"json:",omitempty"`
	StartingSequenceNumber *string `type:"string" json:",omitempty"`
	StreamName             *string `type:"string" required:"true"json:",omitempty"`

	metadataGetShardIteratorInput `json:"-", xml:"-"`
}

type metadataGetShardIteratorInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetShardIteratorOutput struct {
	ShardIterator *string `type:"string" json:",omitempty"`

	metadataGetShardIteratorOutput `json:"-", xml:"-"`
}

type metadataGetShardIteratorOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type HashKeyRange struct {
	EndingHashKey   *string `type:"string" required:"true"json:",omitempty"`
	StartingHashKey *string `type:"string" required:"true"json:",omitempty"`

	metadataHashKeyRange `json:"-", xml:"-"`
}

type metadataHashKeyRange struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListStreamsInput struct {
	ExclusiveStartStreamName *string `type:"string" json:",omitempty"`
	Limit                    *int64  `type:"integer" json:",omitempty"`

	metadataListStreamsInput `json:"-", xml:"-"`
}

type metadataListStreamsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListStreamsOutput struct {
	HasMoreStreams *bool     `type:"boolean" required:"true"json:",omitempty"`
	StreamNames    []*string `type:"list" required:"true"json:",omitempty"`

	metadataListStreamsOutput `json:"-", xml:"-"`
}

type metadataListStreamsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListTagsForStreamInput struct {
	ExclusiveStartTagKey *string `type:"string" json:",omitempty"`
	Limit                *int64  `type:"integer" json:",omitempty"`
	StreamName           *string `type:"string" required:"true"json:",omitempty"`

	metadataListTagsForStreamInput `json:"-", xml:"-"`
}

type metadataListTagsForStreamInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListTagsForStreamOutput struct {
	HasMoreTags *bool  `type:"boolean" required:"true"json:",omitempty"`
	Tags        []*Tag `type:"list" required:"true"json:",omitempty"`

	metadataListTagsForStreamOutput `json:"-", xml:"-"`
}

type metadataListTagsForStreamOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type MergeShardsInput struct {
	AdjacentShardToMerge *string `type:"string" required:"true"json:",omitempty"`
	ShardToMerge         *string `type:"string" required:"true"json:",omitempty"`
	StreamName           *string `type:"string" required:"true"json:",omitempty"`

	metadataMergeShardsInput `json:"-", xml:"-"`
}

type metadataMergeShardsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type MergeShardsOutput struct {
	metadataMergeShardsOutput `json:"-", xml:"-"`
}

type metadataMergeShardsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutRecordInput struct {
	Data                      []byte  `type:"blob" required:"true"json:",omitempty"`
	ExplicitHashKey           *string `type:"string" json:",omitempty"`
	PartitionKey              *string `type:"string" required:"true"json:",omitempty"`
	SequenceNumberForOrdering *string `type:"string" json:",omitempty"`
	StreamName                *string `type:"string" required:"true"json:",omitempty"`

	metadataPutRecordInput `json:"-", xml:"-"`
}

type metadataPutRecordInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutRecordOutput struct {
	SequenceNumber *string `type:"string" required:"true"json:",omitempty"`
	ShardID        *string `locationName:"ShardId" type:"string" required:"true"json:"ShardId,omitempty"`

	metadataPutRecordOutput `json:"-", xml:"-"`
}

type metadataPutRecordOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutRecordsInput struct {
	Records    []*PutRecordsRequestEntry `type:"list" required:"true"json:",omitempty"`
	StreamName *string                   `type:"string" required:"true"json:",omitempty"`

	metadataPutRecordsInput `json:"-", xml:"-"`
}

type metadataPutRecordsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutRecordsOutput struct {
	FailedRecordCount *int64                   `type:"integer" json:",omitempty"`
	Records           []*PutRecordsResultEntry `type:"list" required:"true"json:",omitempty"`

	metadataPutRecordsOutput `json:"-", xml:"-"`
}

type metadataPutRecordsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutRecordsRequestEntry struct {
	Data            []byte  `type:"blob" required:"true"json:",omitempty"`
	ExplicitHashKey *string `type:"string" json:",omitempty"`
	PartitionKey    *string `type:"string" required:"true"json:",omitempty"`

	metadataPutRecordsRequestEntry `json:"-", xml:"-"`
}

type metadataPutRecordsRequestEntry struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutRecordsResultEntry struct {
	ErrorCode      *string `type:"string" json:",omitempty"`
	ErrorMessage   *string `type:"string" json:",omitempty"`
	SequenceNumber *string `type:"string" json:",omitempty"`
	ShardID        *string `locationName:"ShardId" type:"string" json:"ShardId,omitempty"`

	metadataPutRecordsResultEntry `json:"-", xml:"-"`
}

type metadataPutRecordsResultEntry struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Record struct {
	Data           []byte  `type:"blob" required:"true"json:",omitempty"`
	PartitionKey   *string `type:"string" required:"true"json:",omitempty"`
	SequenceNumber *string `type:"string" required:"true"json:",omitempty"`

	metadataRecord `json:"-", xml:"-"`
}

type metadataRecord struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RemoveTagsFromStreamInput struct {
	StreamName *string   `type:"string" required:"true"json:",omitempty"`
	TagKeys    []*string `type:"list" required:"true"json:",omitempty"`

	metadataRemoveTagsFromStreamInput `json:"-", xml:"-"`
}

type metadataRemoveTagsFromStreamInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RemoveTagsFromStreamOutput struct {
	metadataRemoveTagsFromStreamOutput `json:"-", xml:"-"`
}

type metadataRemoveTagsFromStreamOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SequenceNumberRange struct {
	EndingSequenceNumber   *string `type:"string" json:",omitempty"`
	StartingSequenceNumber *string `type:"string" required:"true"json:",omitempty"`

	metadataSequenceNumberRange `json:"-", xml:"-"`
}

type metadataSequenceNumberRange struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Shard struct {
	AdjacentParentShardID *string              `locationName:"AdjacentParentShardId" type:"string" json:"AdjacentParentShardId,omitempty"`
	HashKeyRange          *HashKeyRange        `type:"structure" required:"true"json:",omitempty"`
	ParentShardID         *string              `locationName:"ParentShardId" type:"string" json:"ParentShardId,omitempty"`
	SequenceNumberRange   *SequenceNumberRange `type:"structure" required:"true"json:",omitempty"`
	ShardID               *string              `locationName:"ShardId" type:"string" required:"true"json:"ShardId,omitempty"`

	metadataShard `json:"-", xml:"-"`
}

type metadataShard struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SplitShardInput struct {
	NewStartingHashKey *string `type:"string" required:"true"json:",omitempty"`
	ShardToSplit       *string `type:"string" required:"true"json:",omitempty"`
	StreamName         *string `type:"string" required:"true"json:",omitempty"`

	metadataSplitShardInput `json:"-", xml:"-"`
}

type metadataSplitShardInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SplitShardOutput struct {
	metadataSplitShardOutput `json:"-", xml:"-"`
}

type metadataSplitShardOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type StreamDescription struct {
	HasMoreShards *bool    `type:"boolean" required:"true"json:",omitempty"`
	Shards        []*Shard `type:"list" required:"true"json:",omitempty"`
	StreamARN     *string  `type:"string" required:"true"json:",omitempty"`
	StreamName    *string  `type:"string" required:"true"json:",omitempty"`
	StreamStatus  *string  `type:"string" required:"true"json:",omitempty"`

	metadataStreamDescription `json:"-", xml:"-"`
}

type metadataStreamDescription struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Tag struct {
	Key   *string `type:"string" required:"true"json:",omitempty"`
	Value *string `type:"string" json:",omitempty"`

	metadataTag `json:"-", xml:"-"`
}

type metadataTag struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}