import type { PageServerLoad } from './$types';
import { error, type NumericRange } from '@sveltejs/kit';

export type Paste = {
	content: string;
	encrypted: boolean;
	expiration: string;
	uuid: string;
};


export const load: PageServerLoad = async ({ params, fetch }) => {
	const response = await fetch(`http://localhost:8080/v1/paste/${params.id}`);

	if (!response.ok) {
		throw error(response.status as NumericRange<400, 599>, 'Failed to load paste');
	}

	const paste = await response.json() as Paste;

	return {
		paste
	};
};