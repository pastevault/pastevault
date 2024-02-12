import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { ServiceClient as _proto_ServiceClient, ServiceDefinition as _proto_ServiceDefinition } from './proto/Service';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  proto: {
    Id: MessageTypeDefinition
    Paste: MessageTypeDefinition
    PasteRequest: MessageTypeDefinition
    Service: SubtypeConstructor<typeof grpc.Client, _proto_ServiceClient> & { service: _proto_ServiceDefinition }
  }
}

