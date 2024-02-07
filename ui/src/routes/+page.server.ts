import type { Actions, PageServerLoad } from './$types';
import { z } from 'zod';
import { message, superValidate } from 'sveltekit-superforms/server';
import { fail } from '@sveltejs/kit';

const schema = z.object({
	content: z.string(),
	expiration: z.enum(['1h', '1d', '1w', '1m', '1y']),
	encrypted: z.boolean().default(false)
});

export const load: PageServerLoad = async () => {
	const form = await superValidate(schema);

	return {
		title: 'Paste Vault',
		form
	};
};


type Response = {
	expiration: string;
	uuid: string;
};


export const actions: Actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, schema);

		if (!form.valid) {
			return fail(400, { form });
		}

		const response = await fetch('http://localhost:8080/v1/paste', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(form.data)
		});

		if (!response.ok) {
			return fail(500, { form });
		}

		const json = await response.json() as Response;
		return message(form, `${json.uuid}`);
	}
};