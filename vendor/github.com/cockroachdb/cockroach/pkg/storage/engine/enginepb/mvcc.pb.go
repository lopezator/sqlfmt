// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/engine/enginepb/mvcc.proto

package enginepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_util_hlc1 "github.com/cockroachdb/cockroach/pkg/util/hlc"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MVCCMetadata holds MVCC metadata for a key. Used by storage/engine/mvcc.go.
type MVCCMetadata struct {
	Txn *TxnMeta `protobuf:"bytes,1,opt,name=txn" json:"txn,omitempty"`
	// The timestamp of the most recent versioned value if this is a
	// value that may have multiple versions. For values which may have
	// only one version, the data is stored inline (via raw_bytes), and
	// timestamp is set to zero.
	Timestamp cockroach_util_hlc1.LegacyTimestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp"`
	// Is the most recent value a deletion tombstone?
	Deleted bool `protobuf:"varint,3,opt,name=deleted" json:"deleted"`
	// The size in bytes of the most recent encoded key.
	KeyBytes int64 `protobuf:"varint,4,opt,name=key_bytes,json=keyBytes" json:"key_bytes"`
	// The size in bytes of the most recent versioned value.
	ValBytes int64 `protobuf:"varint,5,opt,name=val_bytes,json=valBytes" json:"val_bytes"`
	// Inline value, used for non-versioned values with zero
	// timestamp. This provides an efficient short circuit of the normal
	// MVCC metadata sentinel and subsequent version rows. If timestamp
	// == (0, 0), then there is only a single MVCC metadata row with
	// value inlined, and with empty timestamp, key_bytes, and
	// val_bytes.
	RawBytes []byte `protobuf:"bytes,6,opt,name=raw_bytes,json=rawBytes" json:"raw_bytes,omitempty"`
	// This provides a measure of protection against replays caused by
	// Raft duplicating merge commands.
	MergeTimestamp *cockroach_util_hlc1.LegacyTimestamp `protobuf:"bytes,7,opt,name=merge_timestamp,json=mergeTimestamp" json:"merge_timestamp,omitempty"`
}

func (m *MVCCMetadata) Reset()                    { *m = MVCCMetadata{} }
func (m *MVCCMetadata) String() string            { return proto.CompactTextString(m) }
func (*MVCCMetadata) ProtoMessage()               {}
func (*MVCCMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMvcc, []int{0} }

// MVCCStats tracks byte and instance counts for various groups of keys,
// values, or key-value pairs; see the field comments for details.
//
// It also tracks two cumulative ages, namely that of intents and non-live
// (i.e. GC-able) bytes. This computation is intrinsically linked to
// last_update_nanos and is easy to get wrong. Updates happen only once every
// full second, as measured by last_update_nanos/1e9. That is, forward updates
// don't change last_update_nanos until an update at a timestamp which,
// truncated to the second, is ahead of last_update_nanos/1e9. Then, that
// difference in seconds times the base quantity (excluding the currently
// running update) is added to the age.
//
// To give an example, if an intent is around from `t=2.5s` to `t=4.1s` (the
// current time), then it contributes an intent age of two seconds (one second
// picked up when crossing `t=3s`, another one at `t=4s`). Similarly, if a
// GC'able kv pair is around for this amount of time, it contributes two seconds
// times its size in bytes.
//
// It gets more complicated when data is
// accounted for with a timestamp behind last_update_nanos. In this case, if
// more than a second has passed (computed via truncation above), the ages have
// to be adjusted to account for this late addition. This isn't hard: add the
// new data's base quantity times the (truncated) number of seconds behind.
// Important to keep in mind with those computations is that (x/1e9 - y/1e9)
// does not equal (x-y)/1e9 in most cases.
//
// Note that this struct must be kept at a fixed size by using fixed-size
// encodings for all fields and by making all fields non-nullable. This is
// so that it can predict its own impact on the size of the system-local
// kv-pairs.
type MVCCStats struct {
	// contains_estimates indicates that the MVCCStats object contains values
	// which have been estimated. This means that the stats should not be used
	// where complete accuracy is required, and instead should be recomputed
	// when necessary.
	ContainsEstimates bool `protobuf:"varint,14,opt,name=contains_estimates,json=containsEstimates" json:"contains_estimates"`
	// last_update_nanos is a timestamp at which the ages were last
	// updated. See the comment on MVCCStats.
	LastUpdateNanos int64 `protobuf:"fixed64,1,opt,name=last_update_nanos,json=lastUpdateNanos" json:"last_update_nanos"`
	// intent_age is the cumulative age of the tracked intents.
	// See the comment on MVCCStats.
	IntentAge int64 `protobuf:"fixed64,2,opt,name=intent_age,json=intentAge" json:"intent_age"`
	// gc_bytes_age is the cumulative age of the non-live data (i.e.
	// data included in key_bytes and val_bytes, but not live_bytes).
	// See the comment on MVCCStats.
	GCBytesAge int64 `protobuf:"fixed64,3,opt,name=gc_bytes_age,json=gcBytesAge" json:"gc_bytes_age"`
	// live_bytes is the number of bytes stored in keys and values which can in
	// principle be read by means of a Scan or Get in the far future, including
	// intents but not deletion tombstones (or their intents). Note that the
	// size of the meta kv pair (which could be explicit or implicit) is
	// included in this. Only the meta kv pair counts for the actual length of
	// the encoded key (regular pairs only count the timestamp suffix).
	LiveBytes int64 `protobuf:"fixed64,4,opt,name=live_bytes,json=liveBytes" json:"live_bytes"`
	// live_count is the number of meta keys tracked under live_bytes.
	LiveCount int64 `protobuf:"fixed64,5,opt,name=live_count,json=liveCount" json:"live_count"`
	// key_bytes is the number of bytes stored in all non-system
	// keys, including live, meta, old, and deleted keys.
	// Only meta keys really account for the "full" key; value
	// keys only for the timestamp suffix.
	KeyBytes int64 `protobuf:"fixed64,6,opt,name=key_bytes,json=keyBytes" json:"key_bytes"`
	// key_count is the number of meta keys tracked under key_bytes.
	KeyCount int64 `protobuf:"fixed64,7,opt,name=key_count,json=keyCount" json:"key_count"`
	// value_bytes is the number of bytes in all non-system version
	// values, including meta values.
	ValBytes int64 `protobuf:"fixed64,8,opt,name=val_bytes,json=valBytes" json:"val_bytes"`
	// val_count is the number of meta values tracked under val_bytes.
	ValCount int64 `protobuf:"fixed64,9,opt,name=val_count,json=valCount" json:"val_count"`
	// intent_bytes is the number of bytes in intent key-value
	// pairs (without their meta keys).
	IntentBytes int64 `protobuf:"fixed64,10,opt,name=intent_bytes,json=intentBytes" json:"intent_bytes"`
	// intent_count is the number of keys tracked under intent_bytes.
	// It is equal to the number of meta keys in the system with
	// a non-empty Transaction proto.
	IntentCount int64 `protobuf:"fixed64,11,opt,name=intent_count,json=intentCount" json:"intent_count"`
	// sys_bytes is the number of bytes stored in system-local kv-pairs.
	// This tracks the same quantity as (key_bytes + val_bytes), but
	// for system-local metadata keys (which aren't counted in either
	// key_bytes or val_bytes). Each of the keys falling into this group
	// is documented in keys/constants.go under the localPrefix constant
	// and is prefixed by either LocalRangeIDPrefix or LocalRangePrefix.
	SysBytes int64 `protobuf:"fixed64,12,opt,name=sys_bytes,json=sysBytes" json:"sys_bytes"`
	// sys_count is the number of meta keys tracked under sys_bytes.
	SysCount int64 `protobuf:"fixed64,13,opt,name=sys_count,json=sysCount" json:"sys_count"`
}

func (m *MVCCStats) Reset()                    { *m = MVCCStats{} }
func (m *MVCCStats) String() string            { return proto.CompactTextString(m) }
func (*MVCCStats) ProtoMessage()               {}
func (*MVCCStats) Descriptor() ([]byte, []int) { return fileDescriptorMvcc, []int{1} }

func init() {
	proto.RegisterType((*MVCCMetadata)(nil), "cockroach.storage.engine.enginepb.MVCCMetadata")
	proto.RegisterType((*MVCCStats)(nil), "cockroach.storage.engine.enginepb.MVCCStats")
}
func (this *MVCCStats) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MVCCStats)
	if !ok {
		that2, ok := that.(MVCCStats)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ContainsEstimates != that1.ContainsEstimates {
		return false
	}
	if this.LastUpdateNanos != that1.LastUpdateNanos {
		return false
	}
	if this.IntentAge != that1.IntentAge {
		return false
	}
	if this.GCBytesAge != that1.GCBytesAge {
		return false
	}
	if this.LiveBytes != that1.LiveBytes {
		return false
	}
	if this.LiveCount != that1.LiveCount {
		return false
	}
	if this.KeyBytes != that1.KeyBytes {
		return false
	}
	if this.KeyCount != that1.KeyCount {
		return false
	}
	if this.ValBytes != that1.ValBytes {
		return false
	}
	if this.ValCount != that1.ValCount {
		return false
	}
	if this.IntentBytes != that1.IntentBytes {
		return false
	}
	if this.IntentCount != that1.IntentCount {
		return false
	}
	if this.SysBytes != that1.SysBytes {
		return false
	}
	if this.SysCount != that1.SysCount {
		return false
	}
	return true
}
func (m *MVCCMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MVCCMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Txn != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMvcc(dAtA, i, uint64(m.Txn.Size()))
		n1, err := m.Txn.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintMvcc(dAtA, i, uint64(m.Timestamp.Size()))
	n2, err := m.Timestamp.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x18
	i++
	if m.Deleted {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x20
	i++
	i = encodeVarintMvcc(dAtA, i, uint64(m.KeyBytes))
	dAtA[i] = 0x28
	i++
	i = encodeVarintMvcc(dAtA, i, uint64(m.ValBytes))
	if m.RawBytes != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMvcc(dAtA, i, uint64(len(m.RawBytes)))
		i += copy(dAtA[i:], m.RawBytes)
	}
	if m.MergeTimestamp != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMvcc(dAtA, i, uint64(m.MergeTimestamp.Size()))
		n3, err := m.MergeTimestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *MVCCStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MVCCStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x9
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.LastUpdateNanos))
	i += 8
	dAtA[i] = 0x11
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.IntentAge))
	i += 8
	dAtA[i] = 0x19
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.GCBytesAge))
	i += 8
	dAtA[i] = 0x21
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.LiveBytes))
	i += 8
	dAtA[i] = 0x29
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.LiveCount))
	i += 8
	dAtA[i] = 0x31
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.KeyBytes))
	i += 8
	dAtA[i] = 0x39
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.KeyCount))
	i += 8
	dAtA[i] = 0x41
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.ValBytes))
	i += 8
	dAtA[i] = 0x49
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.ValCount))
	i += 8
	dAtA[i] = 0x51
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.IntentBytes))
	i += 8
	dAtA[i] = 0x59
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.IntentCount))
	i += 8
	dAtA[i] = 0x61
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.SysBytes))
	i += 8
	dAtA[i] = 0x69
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.SysCount))
	i += 8
	dAtA[i] = 0x70
	i++
	if m.ContainsEstimates {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func encodeVarintMvcc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedMVCCMetadata(r randyMvcc, easy bool) *MVCCMetadata {
	this := &MVCCMetadata{}
	if r.Intn(10) != 0 {
		this.Txn = NewPopulatedTxnMeta(r, easy)
	}
	v1 := cockroach_util_hlc1.NewPopulatedLegacyTimestamp(r, easy)
	this.Timestamp = *v1
	this.Deleted = bool(bool(r.Intn(2) == 0))
	this.KeyBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.KeyBytes *= -1
	}
	this.ValBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.ValBytes *= -1
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(100)
		this.RawBytes = make([]byte, v2)
		for i := 0; i < v2; i++ {
			this.RawBytes[i] = byte(r.Intn(256))
		}
	}
	if r.Intn(10) != 0 {
		this.MergeTimestamp = cockroach_util_hlc1.NewPopulatedLegacyTimestamp(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMVCCStats(r randyMvcc, easy bool) *MVCCStats {
	this := &MVCCStats{}
	this.LastUpdateNanos = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.LastUpdateNanos *= -1
	}
	this.IntentAge = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.IntentAge *= -1
	}
	this.GCBytesAge = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.GCBytesAge *= -1
	}
	this.LiveBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.LiveBytes *= -1
	}
	this.LiveCount = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.LiveCount *= -1
	}
	this.KeyBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.KeyBytes *= -1
	}
	this.KeyCount = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.KeyCount *= -1
	}
	this.ValBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.ValBytes *= -1
	}
	this.ValCount = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.ValCount *= -1
	}
	this.IntentBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.IntentBytes *= -1
	}
	this.IntentCount = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.IntentCount *= -1
	}
	this.SysBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.SysBytes *= -1
	}
	this.SysCount = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.SysCount *= -1
	}
	this.ContainsEstimates = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMvcc interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMvcc(r randyMvcc) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMvcc(r randyMvcc) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneMvcc(r)
	}
	return string(tmps)
}
func randUnrecognizedMvcc(r randyMvcc, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMvcc(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMvcc(dAtA []byte, r randyMvcc, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMvcc(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateMvcc(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateMvcc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMvcc(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMvcc(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMvcc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMvcc(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *MVCCMetadata) Size() (n int) {
	var l int
	_ = l
	if m.Txn != nil {
		l = m.Txn.Size()
		n += 1 + l + sovMvcc(uint64(l))
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovMvcc(uint64(l))
	n += 2
	n += 1 + sovMvcc(uint64(m.KeyBytes))
	n += 1 + sovMvcc(uint64(m.ValBytes))
	if m.RawBytes != nil {
		l = len(m.RawBytes)
		n += 1 + l + sovMvcc(uint64(l))
	}
	if m.MergeTimestamp != nil {
		l = m.MergeTimestamp.Size()
		n += 1 + l + sovMvcc(uint64(l))
	}
	return n
}

func (m *MVCCStats) Size() (n int) {
	var l int
	_ = l
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 9
	n += 2
	return n
}

func sovMvcc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMvcc(x uint64) (n int) {
	return sovMvcc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MVCCMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMvcc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MVCCMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MVCCMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMvcc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Txn == nil {
				m.Txn = &TxnMeta{}
			}
			if err := m.Txn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMvcc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyBytes", wireType)
			}
			m.KeyBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValBytes", wireType)
			}
			m.ValBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMvcc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawBytes = append(m.RawBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.RawBytes == nil {
				m.RawBytes = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MergeTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMvcc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MergeTimestamp == nil {
				m.MergeTimestamp = &cockroach_util_hlc1.LegacyTimestamp{}
			}
			if err := m.MergeTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMvcc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMvcc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MVCCStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMvcc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MVCCStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MVCCStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateNanos", wireType)
			}
			m.LastUpdateNanos = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdateNanos = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentAge", wireType)
			}
			m.IntentAge = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.IntentAge = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GCBytesAge", wireType)
			}
			m.GCBytesAge = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.GCBytesAge = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiveBytes", wireType)
			}
			m.LiveBytes = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.LiveBytes = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiveCount", wireType)
			}
			m.LiveCount = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.LiveCount = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyBytes", wireType)
			}
			m.KeyBytes = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyBytes = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyCount", wireType)
			}
			m.KeyCount = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyCount = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValBytes", wireType)
			}
			m.ValBytes = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.ValBytes = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValCount", wireType)
			}
			m.ValCount = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.ValCount = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentBytes", wireType)
			}
			m.IntentBytes = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.IntentBytes = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 11:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentCount", wireType)
			}
			m.IntentCount = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.IntentCount = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysBytes", wireType)
			}
			m.SysBytes = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.SysBytes = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 13:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysCount", wireType)
			}
			m.SysCount = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.SysCount = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainsEstimates", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ContainsEstimates = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMvcc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMvcc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMvcc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMvcc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMvcc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMvcc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMvcc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMvcc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMvcc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMvcc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage/engine/enginepb/mvcc.proto", fileDescriptorMvcc) }

var fileDescriptorMvcc = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xb1, 0x6e, 0xd3, 0x40,
	0x18, 0xc7, 0x7b, 0x4d, 0x68, 0xec, 0x4b, 0x68, 0xa9, 0xc5, 0x10, 0x15, 0xc9, 0x49, 0x93, 0x81,
	0x88, 0xc1, 0x41, 0x94, 0xa9, 0x62, 0x69, 0x22, 0xd4, 0xa5, 0x65, 0x08, 0x85, 0x81, 0xc5, 0xba,
	0x5e, 0x3e, 0x5d, 0xac, 0x38, 0xe7, 0xc8, 0x77, 0x49, 0xeb, 0xb7, 0xe0, 0x11, 0xfa, 0x18, 0xbc,
	0x00, 0x52, 0x46, 0x46, 0xa6, 0x0a, 0xcc, 0xc2, 0xc0, 0x43, 0xa0, 0xbb, 0xb3, 0x1d, 0x27, 0x48,
	0xc0, 0x94, 0xcb, 0xf7, 0xfd, 0xee, 0xe7, 0xef, 0xfe, 0x67, 0xe3, 0x8e, 0x90, 0x51, 0x4c, 0x18,
	0xf4, 0x81, 0xb3, 0x80, 0xe7, 0x3f, 0xf3, 0xeb, 0xfe, 0x6c, 0x49, 0xa9, 0x37, 0x8f, 0x23, 0x19,
	0x39, 0xc7, 0x34, 0xa2, 0xd3, 0x38, 0x22, 0x74, 0xe2, 0x65, 0xb4, 0x67, 0x30, 0x2f, 0xa7, 0x8f,
	0xba, 0x7f, 0xd3, 0x9c, 0x18, 0xcf, 0x51, 0x6b, 0x21, 0x83, 0xb0, 0x3f, 0x09, 0x69, 0x3f, 0x04,
	0x46, 0x68, 0xe2, 0xcb, 0x60, 0x06, 0x42, 0x92, 0xd9, 0x3c, 0x03, 0x1e, 0xb3, 0x88, 0x45, 0x7a,
	0xd9, 0x57, 0x2b, 0x53, 0xed, 0xfc, 0xda, 0xc5, 0x8d, 0xcb, 0xf7, 0xc3, 0xe1, 0x25, 0x48, 0x32,
	0x26, 0x92, 0x38, 0xaf, 0x70, 0x45, 0xde, 0xf2, 0x26, 0x6a, 0xa3, 0x5e, 0xfd, 0xc5, 0x33, 0xef,
	0x9f, 0xd3, 0x79, 0x57, 0xb7, 0x5c, 0x6d, 0x1e, 0xa9, 0x6d, 0xce, 0x39, 0xb6, 0x8b, 0xe7, 0x36,
	0x77, 0xb5, 0xa3, 0x5b, 0x72, 0xa8, 0x19, 0xbd, 0x49, 0x48, 0xbd, 0x0b, 0x3d, 0xe3, 0x55, 0x8e,
	0x0e, 0xaa, 0xab, 0xfb, 0xd6, 0xce, 0x68, 0xbd, 0xd7, 0x71, 0x71, 0x6d, 0x0c, 0x21, 0x48, 0x18,
	0x37, 0x2b, 0x6d, 0xd4, 0xb3, 0x32, 0x22, 0x2f, 0x3a, 0xc7, 0xd8, 0x9e, 0x42, 0xe2, 0x5f, 0x27,
	0x12, 0x44, 0xb3, 0xda, 0x46, 0xbd, 0x4a, 0x46, 0x58, 0x53, 0x48, 0x06, 0xaa, 0xaa, 0x90, 0x25,
	0x09, 0x33, 0xe4, 0x41, 0x19, 0x59, 0x92, 0xd0, 0x20, 0x4f, 0xb0, 0x1d, 0x93, 0x9b, 0x0c, 0xd9,
	0x6b, 0xa3, 0x5e, 0x63, 0x64, 0xc5, 0xe4, 0xc6, 0x34, 0x2f, 0xf0, 0xc1, 0x0c, 0x62, 0x06, 0xeb,
	0x24, 0x9b, 0xb5, 0xff, 0x3e, 0xd1, 0x68, 0x5f, 0xef, 0x2d, 0xfe, 0x9f, 0x56, 0x3f, 0xdd, 0xb5,
	0x50, 0xe7, 0x73, 0x15, 0xdb, 0x2a, 0xee, 0xb7, 0x92, 0x48, 0xe1, 0x3c, 0xc7, 0x87, 0x21, 0x11,
	0xd2, 0x5f, 0xcc, 0xc7, 0x44, 0x82, 0xcf, 0x09, 0x8f, 0x84, 0x4e, 0xfe, 0x51, 0x36, 0xe9, 0x81,
	0x6a, 0xbf, 0xd3, 0xdd, 0x37, 0xaa, 0xe9, 0x74, 0x31, 0x0e, 0xb8, 0x04, 0x2e, 0x7d, 0xc2, 0x40,
	0x07, 0x9c, 0xa3, 0xb6, 0xa9, 0x9f, 0x31, 0x70, 0x5e, 0xe2, 0x06, 0xa3, 0xe6, 0x50, 0x1a, 0xab,
	0x68, 0xcc, 0x51, 0x58, 0x7a, 0xdf, 0xc2, 0xe7, 0x43, 0x7d, 0xbe, 0x33, 0x06, 0x23, 0xcc, 0x68,
	0xbe, 0x56, 0xea, 0x30, 0x58, 0x42, 0x29, 0xd2, 0x42, 0xad, 0xea, 0x26, 0x93, 0x1c, 0xa2, 0xd1,
	0x82, 0x4b, 0x1d, 0xea, 0x06, 0x34, 0x54, 0xe5, 0xcd, 0xbb, 0xd9, 0x2b, 0x31, 0x1b, 0x77, 0xa3,
	0x10, 0xa3, 0xa9, 0x6d, 0x21, 0x85, 0x65, 0x7d, 0x7d, 0x56, 0x19, 0x29, 0xae, 0x2f, 0x43, 0x8c,
	0xc5, 0xde, 0x42, 0x8c, 0xe5, 0x29, 0x6e, 0x64, 0x81, 0x19, 0x11, 0x2e, 0x51, 0x75, 0xd3, 0x31,
	0xae, 0x35, 0x68, 0x74, 0xf5, 0x3f, 0xc1, 0x62, 0x2e, 0x91, 0x88, 0x4c, 0xd7, 0x28, 0x3f, 0x54,
	0x24, 0xa2, 0x98, 0x4b, 0x21, 0x46, 0xf4, 0x70, 0x0b, 0x31, 0x96, 0x13, 0xec, 0xd0, 0x88, 0x4b,
	0x12, 0x70, 0xe1, 0x83, 0x90, 0xc1, 0x8c, 0x28, 0xdd, 0x7e, 0xe9, 0x55, 0x3f, 0xcc, 0xfb, 0xaf,
	0xf3, 0xf6, 0xa9, 0xa5, 0xde, 0xa1, 0x9f, 0x77, 0x2d, 0x34, 0xe8, 0xac, 0xbe, 0xbb, 0x3b, 0xab,
	0xd4, 0x45, 0x5f, 0x52, 0x17, 0x7d, 0x4d, 0x5d, 0xf4, 0x2d, 0x75, 0xd1, 0xc7, 0x1f, 0xee, 0xce,
	0x07, 0x2b, 0xff, 0x30, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xbc, 0xaf, 0xa8, 0x7e, 0x04,
	0x00, 0x00,
}