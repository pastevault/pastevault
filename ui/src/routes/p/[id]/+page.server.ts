import type { PageServerLoad } from './$types';
import { server } from '$lib/server/grpc';
import { Metadata } from '@grpc/grpc-js';
import { grpcSafe } from '$lib/safe';
import type { Safe } from '$lib/safe';
import type { Paste__Output } from '$lib/proto/proto/Paste';

export type Paste = {
	content: string;
	encrypted: boolean;
	expiration: string;
	uuid: string;
};


export const load: PageServerLoad = async ({ params }) => {
	const { id } = params;
	if (!id) return { status: 400, body: 'Missing Paste ID' };

	const req: Safe<Paste__Output> = await new Promise((r) => {
		const metadata = new Metadata();
		server.getPasteById({ id }, metadata, grpcSafe(r));
	});

	if (req.error) return { status: 404, body: req.msg };

	return {
		paste: req.data,
	}
};