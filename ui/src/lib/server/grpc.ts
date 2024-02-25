import protoLoader from '@grpc/proto-loader';
import { loadPackageDefinition, credentials } from '@grpc/grpc-js';
import type { ProtoGrpcType } from '$lib/proto/main';
import { env } from "$env/dynamic/private";

export const packageDefinition = protoLoader.loadSync(
	'./src/lib/proto/main.proto',
	{
		keepCase: true,
		longs: String,
		defaults: true,
		oneofs: true,
		arrays: true
	}
);

// proto is of type ProtoGrpcType
const proto = loadPackageDefinition(packageDefinition) as unknown as ProtoGrpcType;

const cr =
	env.ENV === "production"
		? credentials.createSsl()
		: credentials.createInsecure();

export const server =  new proto.proto.PasteService(env.SERVER_GRPC, cr);