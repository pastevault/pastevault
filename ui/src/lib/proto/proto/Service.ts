// Original file: ../proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Id as _proto_Id, Id__Output as _proto_Id__Output } from '../proto/Id';
import type { Paste as _proto_Paste, Paste__Output as _proto_Paste__Output } from '../proto/Paste';
import type { PasteRequest as _proto_PasteRequest, PasteRequest__Output as _proto_PasteRequest__Output } from '../proto/PasteRequest';

export interface ServiceClient extends grpc.Client {
  CreatePaste(argument: _proto_PasteRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  CreatePaste(argument: _proto_PasteRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  CreatePaste(argument: _proto_PasteRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  CreatePaste(argument: _proto_PasteRequest, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  createPaste(argument: _proto_PasteRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  createPaste(argument: _proto_PasteRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  createPaste(argument: _proto_PasteRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  createPaste(argument: _proto_PasteRequest, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  
  GetPasteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  GetPasteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  GetPasteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  GetPasteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  getPasteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  getPasteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  getPasteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  getPasteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Paste__Output>): grpc.ClientUnaryCall;
  
}

export interface ServiceHandlers extends grpc.UntypedServiceImplementation {
  CreatePaste: grpc.handleUnaryCall<_proto_PasteRequest__Output, _proto_Paste>;
  
  GetPasteById: grpc.handleUnaryCall<_proto_Id__Output, _proto_Paste>;
  
}

export interface ServiceDefinition extends grpc.ServiceDefinition {
  CreatePaste: MethodDefinition<_proto_PasteRequest, _proto_Paste, _proto_PasteRequest__Output, _proto_Paste__Output>
  GetPasteById: MethodDefinition<_proto_Id, _proto_Paste, _proto_Id__Output, _proto_Paste__Output>
}
