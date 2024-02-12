import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { PasteServiceClient as _proto_PasteServiceClient, PasteServiceDefinition as _proto_PasteServiceDefinition } from './proto/PasteService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  proto: {
    Id: MessageTypeDefinition
    Paste: MessageTypeDefinition
    PasteRequest: MessageTypeDefinition
    PasteService: SubtypeConstructor<typeof grpc.Client, _proto_PasteServiceClient> & { service: _proto_PasteServiceDefinition }
  }
}

