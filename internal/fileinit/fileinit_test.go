package fileinit_test

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"testing"

	proto "github.com/golang/protobuf/proto"
	testpb "github.com/golang/protobuf/v2/internal/testprotos/test"
	"github.com/golang/protobuf/v2/reflect/protodesc"
	"github.com/golang/protobuf/v2/reflect/protoreflect"
	descriptorpb "github.com/golang/protobuf/v2/types/descriptor"
)

func TestInit(t *testing.T) {
	// Compare the FileDescriptorProto for the same test file from two different sources:
	//
	// 1. The result of passing the fileinit-produced FileDescriptor through protodesc.
	// 2. The protoc-generated wire-encoded message.
	//
	// This serves as a test of both fileinit and protodesc.
	got := protodesc.ToFileDescriptorProto(testpb.File_test_test_proto)

	want := &descriptorpb.FileDescriptorProto{}
	zb, _ := (&testpb.TestAllTypes{}).Descriptor()
	r, _ := gzip.NewReader(bytes.NewBuffer(zb))
	b, _ := ioutil.ReadAll(r)
	if err := proto.Unmarshal(b, want); err != nil {
		t.Fatal(err)
	}

	if !proto.Equal(got, want) {
		t.Errorf("protodesc.ToFileDescriptorProto(testpb.Test_protoFile) is not equal to the protoc-generated FileDescriptorProto for internal/testprotos/test/test.proto")
	}

	// Verify that the test proto file provides exhaustive coverage of all descriptor fields.
	seen := make(map[protoreflect.FullName]bool)
	visitFields(want.ProtoReflect(), func(field protoreflect.FieldDescriptor) {
		seen[field.FullName()] = true
	})
	ignore := map[protoreflect.FullName]bool{
		// The protoreflect descriptors don't include source info.
		"google.protobuf.FileDescriptorProto.source_code_info": true,
		"google.protobuf.FileDescriptorProto.syntax":           true,

		// TODO: Test oneof and extension options. Testing these requires extending the
		// options messages (because they contain no user-settable fields), but importing
		// decriptor.proto from test.proto currently causes an import cycle. Add test
		// cases when that import cycle has been fixed.
		"google.protobuf.OneofDescriptorProto.options": true,
	}
	for _, messageName := range []protoreflect.Name{
		"FileDescriptorProto",
		"DescriptorProto",
		"FieldDescriptorProto",
		"OneofDescriptorProto",
		"EnumDescriptorProto",
		"EnumValueDescriptorProto",
		"ServiceDescriptorProto",
		"MethodDescriptorProto",
	} {
		message := descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName(messageName)
		for i, fields := 0, message.Fields(); i < fields.Len(); i++ {
			if name := fields.Get(i).FullName(); !seen[name] && !ignore[name] {
				t.Errorf("No test for descriptor field: %v", name)
			}
		}
	}

	// Verify that message descriptors for map entries have no Go type info.
	mapEntryName := protoreflect.FullName("goproto.proto.test.TestAllTypes.MapInt32Int32Entry")
	d := testpb.File_test_test_proto.DescriptorByName(mapEntryName)
	if _, ok := d.(protoreflect.MessageDescriptor); !ok {
		t.Errorf("message descriptor for %v not found", mapEntryName)
	}
	if _, ok := d.(protoreflect.MessageType); ok {
		t.Errorf("message descriptor for %v must not implement protoreflect.MessageType", mapEntryName)
	}
}

// visitFields calls f for every field set in m and its children.
func visitFields(m protoreflect.Message, f func(protoreflect.FieldDescriptor)) {
	typ := m.Type()
	k := m.KnownFields()
	k.Range(func(num protoreflect.FieldNumber, value protoreflect.Value) bool {
		field := typ.Fields().ByNumber(num)
		f(field)
		switch field.Kind() {
		case protoreflect.MessageKind, protoreflect.GroupKind:
			if field.Cardinality() == protoreflect.Repeated {
				for i, list := 0, value.List(); i < list.Len(); i++ {
					visitFields(list.Get(i).Message(), f)
				}
			} else {
				visitFields(value.Message(), f)
			}
		}
		return true
	})
}
