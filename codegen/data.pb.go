// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: data.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LANGUAGE int32

const (
	LANGUAGE_C          LANGUAGE = 0
	LANGUAGE_CC         LANGUAGE = 1
	LANGUAGE_CLANG      LANGUAGE = 2
	LANGUAGE_PYTHON     LANGUAGE = 3
	LANGUAGE_TYPESCRIPT LANGUAGE = 4
	LANGUAGE_GOLANG     LANGUAGE = 5
)

// Enum value maps for LANGUAGE.
var (
	LANGUAGE_name = map[int32]string{
		0: "C",
		1: "CC",
		2: "CLANG",
		3: "PYTHON",
		4: "TYPESCRIPT",
		5: "GOLANG",
	}
	LANGUAGE_value = map[string]int32{
		"C":          0,
		"CC":         1,
		"CLANG":      2,
		"PYTHON":     3,
		"TYPESCRIPT": 4,
		"GOLANG":     5,
	}
)

func (x LANGUAGE) Enum() *LANGUAGE {
	p := new(LANGUAGE)
	*p = x
	return p
}

func (x LANGUAGE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LANGUAGE) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[0].Descriptor()
}

func (LANGUAGE) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[0]
}

func (x LANGUAGE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LANGUAGE.Descriptor instead.
func (LANGUAGE) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

// Codeforces testlib checker.
type TestlibChecker int32

const (
	//*
	// Compare two doubles with maximal absolute error of 1.5E-6.
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/acmp.cpp
	TestlibChecker_ACMP TestlibChecker = 0
	//*
	// Checker to compare output and answer in the form:
	//
	// Case 1: <number>
	// Case 2: <number>
	// ...
	// Case n: <number>
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/caseicmp.cpp
	TestlibChecker_CASEICMP TestlibChecker = 1
	//*
	// Checker to compare output and answer in the form:
	//
	// Case 1: <number> <number> <number> ... <number>
	// Case 2: <number> <number> <number> ... <number>
	// ...
	// Case n: <number> <number> <number> ... <number>
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/casencmp.cpp
	TestlibChecker_CASENCMP TestlibChecker = 2
	//*
	// Checker to compare output and answer in the form:
	//
	// Case 1: <token> <token> ... <token>
	// Case 2: <token> <token> ... <token>
	// ...
	// Case n: <token> <token> ... <token>
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/casewcmp.cpp
	TestlibChecker_CASEWCMP TestlibChecker = 3
	//*
	// Compare two doubles with maximal absolute error of 1E-6.
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/dcmp.cpp
	TestlibChecker_DCMP TestlibChecker = 4
	//*
	// compare files as sequence of lines.
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/fcmp.cpp
	TestlibChecker_FCMP TestlibChecker = 5
	//*
	// compare two signed huge integers.
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/hcmp.cpp
	TestlibChecker_HCMP TestlibChecker = 6
	//*
	// compare two signed int32's.
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/icmp.cpp
	TestlibChecker_ICMP TestlibChecker = 7
	//*
	// compare files as sequence of tokens in lines.
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/lcmp.cpp
	TestlibChecker_LCMP TestlibChecker = 8
	//*
	// compare ordered sequences of signed int64 numbers.
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/ncmp.cpp
	TestlibChecker_NCMP TestlibChecker = 9
	//*
	// multiple " + YES + "/" + NO + " (case insensetive)
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/nyesno.cpp
	TestlibChecker_NYESNO TestlibChecker = 10
	//*
	// compare two sequences of doubles, max absolute or relative error 1E-4
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/rcmp4.cpp
	TestlibChecker_RCMP4 TestlibChecker = 11
	//*
	// compare two sequences of doubles, max absolute or relative error 1E-6
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/rcmp6.cpp
	TestlibChecker_RCMP6 TestlibChecker = 12
	//*
	// compare two sequences of doubles, max absolute or relative error 1E-9
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/rcmp9.cpp
	TestlibChecker_RCMP9 TestlibChecker = 13
	//*
	// compare two sequences of doubles, maximal absolute error 1.5E-5
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/rncmp.cpp
	TestlibChecker_RNCMP TestlibChecker = 14
	//*
	// compare unordered sequences of signed int64 numbers
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/uncmp.cpp
	TestlibChecker_UNCMP TestlibChecker = 15
	//*
	// compare sequences of tokens
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/wcmp.cpp
	TestlibChecker_WCMP TestlibChecker = 16
	//*
	// YES + " or " + NO + " (case insensetive)"
	// https://github.com/MikeMirzayanov/testlib/blob/master/checkers/yesno.cpp
	TestlibChecker_YESNO TestlibChecker = 17
)

// Enum value maps for TestlibChecker.
var (
	TestlibChecker_name = map[int32]string{
		0:  "ACMP",
		1:  "CASEICMP",
		2:  "CASENCMP",
		3:  "CASEWCMP",
		4:  "DCMP",
		5:  "FCMP",
		6:  "HCMP",
		7:  "ICMP",
		8:  "LCMP",
		9:  "NCMP",
		10: "NYESNO",
		11: "RCMP4",
		12: "RCMP6",
		13: "RCMP9",
		14: "RNCMP",
		15: "UNCMP",
		16: "WCMP",
		17: "YESNO",
	}
	TestlibChecker_value = map[string]int32{
		"ACMP":     0,
		"CASEICMP": 1,
		"CASENCMP": 2,
		"CASEWCMP": 3,
		"DCMP":     4,
		"FCMP":     5,
		"HCMP":     6,
		"ICMP":     7,
		"LCMP":     8,
		"NCMP":     9,
		"NYESNO":   10,
		"RCMP4":    11,
		"RCMP6":    12,
		"RCMP9":    13,
		"RNCMP":    14,
		"UNCMP":    15,
		"WCMP":     16,
		"YESNO":    17,
	}
)

func (x TestlibChecker) Enum() *TestlibChecker {
	p := new(TestlibChecker)
	*p = x
	return p
}

func (x TestlibChecker) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestlibChecker) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[1].Descriptor()
}

func (TestlibChecker) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[1]
}

func (x TestlibChecker) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestlibChecker.Descriptor instead.
func (TestlibChecker) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

type ExcutionPolicy int32

const (
	// This will guarantee the sequential execution order of test cases.
	ExcutionPolicy_SEQUENTIAL_EXECUTION ExcutionPolicy = 0
	// This will not guarantee the execution order of test cases.
	ExcutionPolicy_OUT_OF_ORDER_EXECUTION ExcutionPolicy = 1
)

// Enum value maps for ExcutionPolicy.
var (
	ExcutionPolicy_name = map[int32]string{
		0: "SEQUENTIAL_EXECUTION",
		1: "OUT_OF_ORDER_EXECUTION",
	}
	ExcutionPolicy_value = map[string]int32{
		"SEQUENTIAL_EXECUTION":   0,
		"OUT_OF_ORDER_EXECUTION": 1,
	}
)

func (x ExcutionPolicy) Enum() *ExcutionPolicy {
	p := new(ExcutionPolicy)
	*p = x
	return p
}

func (x ExcutionPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExcutionPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[2].Descriptor()
}

func (ExcutionPolicy) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[2]
}

func (x ExcutionPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExcutionPolicy.Descriptor instead.
func (ExcutionPolicy) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

// This is the program running exit policy, will only take effect on the test
// cases judge, in other words, any exception thrown at like compile stage or
// file download stage will still cause exit.
type ExitPolicy int32

const (
	// Will exit when first test case broken.
	ExitPolicy_EXIT_WHEN_FIRST_FAIL ExitPolicy = 0
	// Enfore run every test case.
	ExitPolicy_ENFORCE_RUN_ALL ExitPolicy = 1
)

// Enum value maps for ExitPolicy.
var (
	ExitPolicy_name = map[int32]string{
		0: "EXIT_WHEN_FIRST_FAIL",
		1: "ENFORCE_RUN_ALL",
	}
	ExitPolicy_value = map[string]int32{
		"EXIT_WHEN_FIRST_FAIL": 0,
		"ENFORCE_RUN_ALL":      1,
	}
)

func (x ExitPolicy) Enum() *ExitPolicy {
	p := new(ExitPolicy)
	*p = x
	return p
}

func (x ExitPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExitPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[3].Descriptor()
}

func (ExitPolicy) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[3]
}

func (x ExitPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExitPolicy.Descriptor instead.
func (ExitPolicy) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

// Use the HTTP GET method to retrieve test case.
type NetworkResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The url of the resource(file), the osiris will use HTTP GET method to fetch
	// it.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	//
	// sha256 sum of the file, the osiris will not cache the file if absent or not
	// equal to the osiris calculated result.
	// The osiris will calculate the file's sha256 via Golang's offical sha256
	// package(https://golang.org/pkg/crypto/sha256/), The client should keep
	// consistent with this library.
	Sha256Sum []byte `protobuf:"bytes,2,opt,name=sha256_sum,json=sha256Sum,proto3" json:"sha256_sum,omitempty"`
}

func (x *NetworkResource) Reset() {
	*x = NetworkResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkResource) ProtoMessage() {}

func (x *NetworkResource) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkResource.ProtoReflect.Descriptor instead.
func (*NetworkResource) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkResource) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NetworkResource) GetSha256Sum() []byte {
	if x != nil {
		return x.Sha256Sum
	}
	return nil
}

type TestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//	*TestCase_NetworkResource
	//	*TestCase_RawContent
	Source isTestCase_Source `protobuf_oneof:"source"`
	// Since the program running result is a stream response, so this field used
	// for client to track every case, osiris will copy this field into response.
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *TestCase) Reset() {
	*x = TestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCase) ProtoMessage() {}

func (x *TestCase) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCase.ProtoReflect.Descriptor instead.
func (*TestCase) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (m *TestCase) GetSource() isTestCase_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *TestCase) GetNetworkResource() *NetworkResource {
	if x, ok := x.GetSource().(*TestCase_NetworkResource); ok {
		return x.NetworkResource
	}
	return nil
}

func (x *TestCase) GetRawContent() []byte {
	if x, ok := x.GetSource().(*TestCase_RawContent); ok {
		return x.RawContent
	}
	return nil
}

func (x *TestCase) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

type isTestCase_Source interface {
	isTestCase_Source()
}

type TestCase_NetworkResource struct {
	// Give the test case as network resource.
	NetworkResource *NetworkResource `protobuf:"bytes,1,opt,name=network_resource,json=networkResource,proto3,oneof"`
}

type TestCase_RawContent struct {
	// Directly set the test case content.
	RawContent []byte `protobuf:"bytes,2,opt,name=raw_content,json=rawContent,proto3,oneof"`
}

func (*TestCase_NetworkResource) isTestCase_Source() {}

func (*TestCase_RawContent) isTestCase_Source() {}

type Program struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Program language.
	Language LANGUAGE `protobuf:"varint,1,opt,name=language,proto3,enum=LANGUAGE" json:"language,omitempty"`
	// Program content.
	RawContent string `protobuf:"bytes,2,opt,name=raw_content,json=rawContent,proto3" json:"raw_content,omitempty"`
}

func (x *Program) Reset() {
	*x = Program{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Program) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Program) ProtoMessage() {}

func (x *Program) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Program.ProtoReflect.Descriptor instead.
func (*Program) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *Program) GetLanguage() LANGUAGE {
	if x != nil {
		return x.Language
	}
	return LANGUAGE_C
}

func (x *Program) GetRawContent() string {
	if x != nil {
		return x.RawContent
	}
	return ""
}

type Checker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//	*Checker_TestlibChecker
	//	*Checker_CustomizedChecker
	Source isChecker_Source `protobuf_oneof:"source"`
}

func (x *Checker) Reset() {
	*x = Checker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checker) ProtoMessage() {}

func (x *Checker) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checker.ProtoReflect.Descriptor instead.
func (*Checker) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (m *Checker) GetSource() isChecker_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Checker) GetTestlibChecker() TestlibChecker {
	if x, ok := x.GetSource().(*Checker_TestlibChecker); ok {
		return x.TestlibChecker
	}
	return TestlibChecker_ACMP
}

func (x *Checker) GetCustomizedChecker() *Program {
	if x, ok := x.GetSource().(*Checker_CustomizedChecker); ok {
		return x.CustomizedChecker
	}
	return nil
}

type isChecker_Source interface {
	isChecker_Source()
}

type Checker_TestlibChecker struct {
	// With testlib checker.
	TestlibChecker TestlibChecker `protobuf:"varint,1,opt,name=testlib_checker,json=testlibChecker,proto3,enum=TestlibChecker,oneof"`
}

type Checker_CustomizedChecker struct {
	// Customized checker.
	CustomizedChecker *Program `protobuf:"bytes,2,opt,name=customized_checker,json=customizedChecker,proto3,oneof"`
}

func (*Checker_TestlibChecker) isChecker_Source() {}

func (*Checker_CustomizedChecker) isChecker_Source() {}

type Restriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Compile time restriction.
	CompileTimeRestrictionMillisecond int64 `protobuf:"varint,1,opt,name=compile_time_restriction_millisecond,json=compileTimeRestrictionMillisecond,proto3" json:"compile_time_restriction_millisecond,omitempty"`
	// Runtime memory restriction.
	MemoryRestrictionKib int64 `protobuf:"varint,2,opt,name=memory_restriction_kib,json=memoryRestrictionKib,proto3" json:"memory_restriction_kib,omitempty"`
	// Runtime cpu time restriction.
	CpuTimeRestrictionMillisecond int64 `protobuf:"varint,3,opt,name=cpu_time_restriction_millisecond,json=cpuTimeRestrictionMillisecond,proto3" json:"cpu_time_restriction_millisecond,omitempty"`
	// Output limit restriction.
	OutputRestrictionKib int64 `protobuf:"varint,4,opt,name=output_restriction_kib,json=outputRestrictionKib,proto3" json:"output_restriction_kib,omitempty"`
	// Checker runtime memory restriction.
	CheckerMemoryRestrictionKib int64 `protobuf:"varint,5,opt,name=checker_memory_restriction_kib,json=checkerMemoryRestrictionKib,proto3" json:"checker_memory_restriction_kib,omitempty"`
	// Checker runtime cpu time restriction.
	CheckerCpuTimeRestrictionMillisecond int64 `protobuf:"varint,6,opt,name=checker_cpu_time_restriction_millisecond,json=checkerCpuTimeRestrictionMillisecond,proto3" json:"checker_cpu_time_restriction_millisecond,omitempty"`
}

func (x *Restriction) Reset() {
	*x = Restriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restriction) ProtoMessage() {}

func (x *Restriction) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restriction.ProtoReflect.Descriptor instead.
func (*Restriction) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *Restriction) GetCompileTimeRestrictionMillisecond() int64 {
	if x != nil {
		return x.CompileTimeRestrictionMillisecond
	}
	return 0
}

func (x *Restriction) GetMemoryRestrictionKib() int64 {
	if x != nil {
		return x.MemoryRestrictionKib
	}
	return 0
}

func (x *Restriction) GetCpuTimeRestrictionMillisecond() int64 {
	if x != nil {
		return x.CpuTimeRestrictionMillisecond
	}
	return 0
}

func (x *Restriction) GetOutputRestrictionKib() int64 {
	if x != nil {
		return x.OutputRestrictionKib
	}
	return 0
}

func (x *Restriction) GetCheckerMemoryRestrictionKib() int64 {
	if x != nil {
		return x.CheckerMemoryRestrictionKib
	}
	return 0
}

func (x *Restriction) GetCheckerCpuTimeRestrictionMillisecond() int64 {
	if x != nil {
		return x.CheckerCpuTimeRestrictionMillisecond
	}
	return 0
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0f,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x5f, 0x73, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x53, 0x75, 0x6d,
	0x22, 0x97, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0b,
	0x72, 0x61, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41,
	0x47, 0x45, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x61, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x8a, 0x01,
	0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0f, 0x74, 0x65, 0x73,
	0x74, 0x6c, 0x69, 0x62, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x6c, 0x69, 0x62, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x69, 0x62, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x11, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xb0, 0x03, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x21, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x69, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69,
	0x62, 0x12, 0x47, 0x0a, 0x20, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x63, 0x70, 0x75,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6b, 0x69, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x62,
	0x12, 0x43, 0x0a, 0x1e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x69, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x69, 0x62, 0x12, 0x56, 0x0a, 0x28, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x5f, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x24, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x43, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2a, 0x4c, 0x0a,
	0x08, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x43, 0x43, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x41, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x59, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x10, 0x04, 0x12,
	0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x4c, 0x41, 0x4e, 0x47, 0x10, 0x05, 0x2a, 0xd8, 0x01, 0x0a, 0x0e,
	0x54, 0x65, 0x73, 0x74, 0x6c, 0x69, 0x62, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x08,
	0x0a, 0x04, 0x41, 0x43, 0x4d, 0x50, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x53, 0x45,
	0x49, 0x43, 0x4d, 0x50, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x53, 0x45, 0x4e, 0x43,
	0x4d, 0x50, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x53, 0x45, 0x57, 0x43, 0x4d, 0x50,
	0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x43, 0x4d, 0x50, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x43, 0x4d, 0x50, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x43, 0x4d, 0x50, 0x10, 0x06,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x43, 0x4d, 0x50, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x43,
	0x4d, 0x50, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x43, 0x4d, 0x50, 0x10, 0x09, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x59, 0x45, 0x53, 0x4e, 0x4f, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x43,
	0x4d, 0x50, 0x34, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x43, 0x4d, 0x50, 0x36, 0x10, 0x0c,
	0x12, 0x09, 0x0a, 0x05, 0x52, 0x43, 0x4d, 0x50, 0x39, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x52,
	0x4e, 0x43, 0x4d, 0x50, 0x10, 0x0e, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x43, 0x4d, 0x50, 0x10,
	0x0f, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x43, 0x4d, 0x50, 0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x59,
	0x45, 0x53, 0x4e, 0x4f, 0x10, 0x11, 0x2a, 0x46, 0x0a, 0x0e, 0x45, 0x78, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x51, 0x55,
	0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x2a, 0x3b,
	0x0a, 0x0a, 0x45, 0x78, 0x69, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x14,
	0x45, 0x58, 0x49, 0x54, 0x5f, 0x57, 0x48, 0x45, 0x4e, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x46, 0x4f, 0x52, 0x43,
	0x45, 0x5f, 0x52, 0x55, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x74, 0x65, 0x63, 0x65,
	0x2d, 0x61, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x2f, 0x6f, 0x73, 0x69, 0x72, 0x69, 0x73, 0x2d,
	0x6e, 0x65, 0x78, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_data_proto_goTypes = []interface{}{
	(LANGUAGE)(0),           // 0: LANGUAGE
	(TestlibChecker)(0),     // 1: TestlibChecker
	(ExcutionPolicy)(0),     // 2: ExcutionPolicy
	(ExitPolicy)(0),         // 3: ExitPolicy
	(*NetworkResource)(nil), // 4: NetworkResource
	(*TestCase)(nil),        // 5: TestCase
	(*Program)(nil),         // 6: Program
	(*Checker)(nil),         // 7: Checker
	(*Restriction)(nil),     // 8: Restriction
}
var file_data_proto_depIdxs = []int32{
	4, // 0: TestCase.network_resource:type_name -> NetworkResource
	0, // 1: Program.language:type_name -> LANGUAGE
	1, // 2: Checker.testlib_checker:type_name -> TestlibChecker
	6, // 3: Checker.customized_checker:type_name -> Program
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkResource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Program); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restriction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_data_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TestCase_NetworkResource)(nil),
		(*TestCase_RawContent)(nil),
	}
	file_data_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Checker_TestlibChecker)(nil),
		(*Checker_CustomizedChecker)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		EnumInfos:         file_data_proto_enumTypes,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}
