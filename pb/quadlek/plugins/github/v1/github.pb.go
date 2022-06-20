// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: quadlek/plugins/github/v1/github.proto

package v1

import (
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

type AuthState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResponseUrl string `protobuf:"bytes,3,opt,name=response_url,json=responseUrl,proto3" json:"response_url,omitempty"`
	ExpireTime  int64  `protobuf:"varint,4,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *AuthState) Reset() {
	*x = AuthState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quadlek_plugins_github_v1_github_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthState) ProtoMessage() {}

func (x *AuthState) ProtoReflect() protoreflect.Message {
	mi := &file_quadlek_plugins_github_v1_github_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthState.ProtoReflect.Descriptor instead.
func (*AuthState) Descriptor() ([]byte, []int) {
	return file_quadlek_plugins_github_v1_github_proto_rawDescGZIP(), []int{0}
}

func (x *AuthState) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthState) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthState) GetResponseUrl() string {
	if x != nil {
		return x.ResponseUrl
	}
	return ""
}

func (x *AuthState) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresAt    int64  `protobuf:"varint,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quadlek_plugins_github_v1_github_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_quadlek_plugins_github_v1_github_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_quadlek_plugins_github_v1_github_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Token) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Scopes     []string `protobuf:"bytes,2,rep,name=scopes,proto3" json:"scopes,omitempty"`
	GithubUser string   `protobuf:"bytes,3,opt,name=github_user,json=githubUser,proto3" json:"github_user,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quadlek_plugins_github_v1_github_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_quadlek_plugins_github_v1_github_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_quadlek_plugins_github_v1_github_proto_rawDescGZIP(), []int{2}
}

func (x *AuthToken) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *AuthToken) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *AuthToken) GetGithubUser() string {
	if x != nil {
		return x.GithubUser
	}
	return ""
}

var File_quadlek_plugins_github_v1_github_proto protoreflect.FileDescriptor

var file_quadlek_plugins_github_v1_github_proto_rawDesc = []byte{
	0x0a, 0x26, 0x71, 0x75, 0x61, 0x64, 0x6c, 0x65, 0x6b, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x71, 0x75, 0x61, 0x64, 0x6c, 0x65,
	0x6b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x22, 0x78, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8d, 0x01,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x7c, 0x0a,
	0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x71, 0x75, 0x61, 0x64,
	0x6c, 0x65, 0x6b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x72, 0x77, 0x69, 0x6e,
	0x2f, 0x71, 0x75, 0x61, 0x64, 0x6c, 0x65, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quadlek_plugins_github_v1_github_proto_rawDescOnce sync.Once
	file_quadlek_plugins_github_v1_github_proto_rawDescData = file_quadlek_plugins_github_v1_github_proto_rawDesc
)

func file_quadlek_plugins_github_v1_github_proto_rawDescGZIP() []byte {
	file_quadlek_plugins_github_v1_github_proto_rawDescOnce.Do(func() {
		file_quadlek_plugins_github_v1_github_proto_rawDescData = protoimpl.X.CompressGZIP(file_quadlek_plugins_github_v1_github_proto_rawDescData)
	})
	return file_quadlek_plugins_github_v1_github_proto_rawDescData
}

var file_quadlek_plugins_github_v1_github_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_quadlek_plugins_github_v1_github_proto_goTypes = []interface{}{
	(*AuthState)(nil), // 0: quadlek.plugins.github.v1.AuthState
	(*Token)(nil),     // 1: quadlek.plugins.github.v1.Token
	(*AuthToken)(nil), // 2: quadlek.plugins.github.v1.AuthToken
}
var file_quadlek_plugins_github_v1_github_proto_depIdxs = []int32{
	1, // 0: quadlek.plugins.github.v1.AuthToken.token:type_name -> quadlek.plugins.github.v1.Token
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_quadlek_plugins_github_v1_github_proto_init() }
func file_quadlek_plugins_github_v1_github_proto_init() {
	if File_quadlek_plugins_github_v1_github_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quadlek_plugins_github_v1_github_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthState); i {
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
		file_quadlek_plugins_github_v1_github_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_quadlek_plugins_github_v1_github_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_quadlek_plugins_github_v1_github_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quadlek_plugins_github_v1_github_proto_goTypes,
		DependencyIndexes: file_quadlek_plugins_github_v1_github_proto_depIdxs,
		MessageInfos:      file_quadlek_plugins_github_v1_github_proto_msgTypes,
	}.Build()
	File_quadlek_plugins_github_v1_github_proto = out.File
	file_quadlek_plugins_github_v1_github_proto_rawDesc = nil
	file_quadlek_plugins_github_v1_github_proto_goTypes = nil
	file_quadlek_plugins_github_v1_github_proto_depIdxs = nil
}
