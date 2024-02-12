import type { Actions, PageServerLoad } from './$types';
import { z } from 'zod';
import { superValidate } from 'sveltekit-superforms/server';
import { fail, redirect } from '@sveltejs/kit';
import { grpcSafe, type Safe } from '$lib/safe';
import type { Paste__Output } from '$lib/proto/proto/Paste';
import { Metadata } from '@grpc/grpc-js';
import { server } from '$lib/server/grpc';
import type { PasteRequest } from '$lib/proto/proto/PasteRequest';

const schema = z.object({
	content: z.string(),
	expiration: z.enum(['1h', '1d', '1w', '1m', '1y']),
	encrypted: z.boolean().default(false).optional()
});

export const load: PageServerLoad = async () => {
	const form = await superValidate(schema);

	return {
		title: 'Paste Vault',
		form
	};
};


export const actions: Actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, schema);

		if (!form.valid) {
			return fail(400, { form });
		}

		const paste: PasteRequest = {
			content: form.data.content,
			encrypted: form.data.encrypted,
			expiration: form.data.expiration
		};

		const req: Safe<Paste__Output> = await new Promise((r) => {
			const metadata = new Metadata();
			server.CreatePaste(paste, metadata, grpcSafe(r));
		});

		if (req.error) {
			if (req.fields) {
				return fail(400, { form: { ...form, errors: req.fields } });
			}
			return fail(500, { message: req.msg });
		}

		return redirect(303, `/p/${req.data.id}`);
	}
};