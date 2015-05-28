// Code generated by protoc-gen-go.
// source: github.com/google/battery-historian/pb/build_proto/build.proto
// DO NOT EDIT!

/*
Package build is a generated protocol buffer package.

It is generated from these files:
	github.com/google/battery-historian/pb/build_proto/build.proto

It has these top-level messages:
	Build
*/
package build

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Build struct {
	// Fingerprint, e.g. "google/mysid/toro:4.0.4/IMM06/243892:userdebug/dev-keys"
	Fingerprint *string `protobuf:"bytes,1,opt,name=fingerprint" json:"fingerprint,omitempty"`
	// Carrier, e.g. "google"
	Brand *string `protobuf:"bytes,2,opt,name=brand" json:"brand,omitempty"`
	// Product name, e.g. "mysid"
	Product *string `protobuf:"bytes,3,opt,name=product" json:"product,omitempty"`
	// Product name, e.g. "toro"
	Device *string `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
	// Release version, e.g. "4.0.4"
	Release *string `protobuf:"bytes,5,opt,name=release" json:"release,omitempty"`
	// Build id, e.g. "IMM06"
	BuildId *string `protobuf:"bytes,6,opt,name=build_id" json:"build_id,omitempty"`
	// Incremental build id, e.g. "243892"
	Incremental *string `protobuf:"bytes,7,opt,name=incremental" json:"incremental,omitempty"`
	// Type of build, e.g. "userdebug"
	Type *string `protobuf:"bytes,8,opt,name=type" json:"type,omitempty"`
	// Tags, e.g. "dev-keys"
	Tags             []string `protobuf:"bytes,9,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}

func (m *Build) GetFingerprint() string {
	if m != nil && m.Fingerprint != nil {
		return *m.Fingerprint
	}
	return ""
}

func (m *Build) GetBrand() string {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return ""
}

func (m *Build) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *Build) GetDevice() string {
	if m != nil && m.Device != nil {
		return *m.Device
	}
	return ""
}

func (m *Build) GetRelease() string {
	if m != nil && m.Release != nil {
		return *m.Release
	}
	return ""
}

func (m *Build) GetBuildId() string {
	if m != nil && m.BuildId != nil {
		return *m.BuildId
	}
	return ""
}

func (m *Build) GetIncremental() string {
	if m != nil && m.Incremental != nil {
		return *m.Incremental
	}
	return ""
}

func (m *Build) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Build) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
}