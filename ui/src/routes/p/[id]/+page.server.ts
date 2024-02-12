import type { PageServerLoad } from './$types';
import { server } from '$lib/server/grpc';
import { Metadata, type ServiceError } from '@grpc/grpc-js';
import type { Paste__Output } from '$lib/proto/proto/Paste';
import { error } from '@sveltejs/kit';

export type Paste = {
	content: string;
	encrypted: boolean;
	expiration: string;
	uuid: string;
};


export const load: PageServerLoad = async ({ params }) => {
	const { id } = params;
	if (!id) throw error(404, 'Not Found');

	const paste: Paste__Output | ServiceError = await new Promise((r) => {
		const metadata = new Metadata();
		server.getPasteById({ id }, metadata, (err, data) => {
			if (err) {
				r(err);
			}
			r(data as Paste__Output);
		});

	});

	if (paste instanceof Error) {
		throw error(404, "Paste not found");
	}

	return {
		paste
	}
};