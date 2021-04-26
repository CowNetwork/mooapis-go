// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// Returns the player or creates/updates it.
	GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error)
	// Returns the player.
	GetPlayerById(ctx context.Context, in *GetPlayerByIdRequest, opts ...grpc.CallOption) (*GetPlayerByIdResponse, error)
	// Returns the player.
	GetPlayerByName(ctx context.Context, in *GetPlayerByNameRequest, opts ...grpc.CallOption) (*GetPlayerByNameResponse, error)
	// Returns the players or creates/updates them.
	GetPlayers(ctx context.Context, in *GetPlayersRequest, opts ...grpc.CallOption) (*GetPlayersResponse, error)
	// Returns the players or creates/updates them.
	GetPlayersById(ctx context.Context, in *GetPlayersByIdRequest, opts ...grpc.CallOption) (*GetPlayersByIdResponse, error)
	// Returns the user by id.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Returns the player.
	GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*GetUserByNameResponse, error)
	// Returns the players assigned to this user.
	GetUserPlayers(ctx context.Context, in *GetUserPlayersRequest, opts ...grpc.CallOption) (*GetUserPlayersResponse, error)
	// Updates the metadata and returns the player.
	UpdatePlayerMetadata(ctx context.Context, in *UpdatePlayerMetadataRequest, opts ...grpc.CallOption) (*UpdatePlayerMetadataResponse, error)
	// Updates the metadata and returns the user.
	UpdateUserMetadata(ctx context.Context, in *UpdateUserMetadataRequest, opts ...grpc.CallOption) (*UpdateUserMetadataResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error) {
	out := new(GetPlayerResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPlayerById(ctx context.Context, in *GetPlayerByIdRequest, opts ...grpc.CallOption) (*GetPlayerByIdResponse, error) {
	out := new(GetPlayerByIdResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetPlayerById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPlayerByName(ctx context.Context, in *GetPlayerByNameRequest, opts ...grpc.CallOption) (*GetPlayerByNameResponse, error) {
	out := new(GetPlayerByNameResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetPlayerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPlayers(ctx context.Context, in *GetPlayersRequest, opts ...grpc.CallOption) (*GetPlayersResponse, error) {
	out := new(GetPlayersResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPlayersById(ctx context.Context, in *GetPlayersByIdRequest, opts ...grpc.CallOption) (*GetPlayersByIdResponse, error) {
	out := new(GetPlayersByIdResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetPlayersById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*GetUserByNameResponse, error) {
	out := new(GetUserByNameResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserPlayers(ctx context.Context, in *GetUserPlayersRequest, opts ...grpc.CallOption) (*GetUserPlayersResponse, error) {
	out := new(GetUserPlayersResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/GetUserPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePlayerMetadata(ctx context.Context, in *UpdatePlayerMetadataRequest, opts ...grpc.CallOption) (*UpdatePlayerMetadataResponse, error) {
	out := new(UpdatePlayerMetadataResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/UpdatePlayerMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserMetadata(ctx context.Context, in *UpdateUserMetadataRequest, opts ...grpc.CallOption) (*UpdateUserMetadataResponse, error) {
	out := new(UpdateUserMetadataResponse)
	err := c.cc.Invoke(ctx, "/cow.user.v1.UserService/UpdateUserMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// Returns the player or creates/updates it.
	GetPlayer(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error)
	// Returns the player.
	GetPlayerById(context.Context, *GetPlayerByIdRequest) (*GetPlayerByIdResponse, error)
	// Returns the player.
	GetPlayerByName(context.Context, *GetPlayerByNameRequest) (*GetPlayerByNameResponse, error)
	// Returns the players or creates/updates them.
	GetPlayers(context.Context, *GetPlayersRequest) (*GetPlayersResponse, error)
	// Returns the players or creates/updates them.
	GetPlayersById(context.Context, *GetPlayersByIdRequest) (*GetPlayersByIdResponse, error)
	// Returns the user by id.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Returns the player.
	GetUserByName(context.Context, *GetUserByNameRequest) (*GetUserByNameResponse, error)
	// Returns the players assigned to this user.
	GetUserPlayers(context.Context, *GetUserPlayersRequest) (*GetUserPlayersResponse, error)
	// Updates the metadata and returns the player.
	UpdatePlayerMetadata(context.Context, *UpdatePlayerMetadataRequest) (*UpdatePlayerMetadataResponse, error)
	// Updates the metadata and returns the user.
	UpdateUserMetadata(context.Context, *UpdateUserMetadataRequest) (*UpdateUserMetadataResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetPlayer(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (UnimplementedUserServiceServer) GetPlayerById(context.Context, *GetPlayerByIdRequest) (*GetPlayerByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerById not implemented")
}
func (UnimplementedUserServiceServer) GetPlayerByName(context.Context, *GetPlayerByNameRequest) (*GetPlayerByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerByName not implemented")
}
func (UnimplementedUserServiceServer) GetPlayers(context.Context, *GetPlayersRequest) (*GetPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedUserServiceServer) GetPlayersById(context.Context, *GetPlayersByIdRequest) (*GetPlayersByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayersById not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByName(context.Context, *GetUserByNameRequest) (*GetUserByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedUserServiceServer) GetUserPlayers(context.Context, *GetUserPlayersRequest) (*GetUserPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPlayers not implemented")
}
func (UnimplementedUserServiceServer) UpdatePlayerMetadata(context.Context, *UpdatePlayerMetadataRequest) (*UpdatePlayerMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayerMetadata not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserMetadata(context.Context, *UpdateUserMetadataRequest) (*UpdateUserMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMetadata not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPlayer(ctx, req.(*GetPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPlayerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPlayerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetPlayerById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPlayerById(ctx, req.(*GetPlayerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPlayerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPlayerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetPlayerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPlayerByName(ctx, req.(*GetPlayerByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPlayers(ctx, req.(*GetPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPlayersById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayersByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPlayersById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetPlayersById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPlayersById(ctx, req.(*GetPlayersByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByName(ctx, req.(*GetUserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/GetUserPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserPlayers(ctx, req.(*GetUserPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePlayerMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlayerMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePlayerMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/UpdatePlayerMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePlayerMetadata(ctx, req.(*UpdatePlayerMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cow.user.v1.UserService/UpdateUserMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserMetadata(ctx, req.(*UpdateUserMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cow.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayer",
			Handler:    _UserService_GetPlayer_Handler,
		},
		{
			MethodName: "GetPlayerById",
			Handler:    _UserService_GetPlayerById_Handler,
		},
		{
			MethodName: "GetPlayerByName",
			Handler:    _UserService_GetPlayerByName_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _UserService_GetPlayers_Handler,
		},
		{
			MethodName: "GetPlayersById",
			Handler:    _UserService_GetPlayersById_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _UserService_GetUserByName_Handler,
		},
		{
			MethodName: "GetUserPlayers",
			Handler:    _UserService_GetUserPlayers_Handler,
		},
		{
			MethodName: "UpdatePlayerMetadata",
			Handler:    _UserService_UpdatePlayerMetadata_Handler,
		},
		{
			MethodName: "UpdateUserMetadata",
			Handler:    _UserService_UpdateUserMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cow/user/v1/user_service_api.proto",
}
