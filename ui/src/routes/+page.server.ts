import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => {
	return {
		title: 'This is a page'
	};
};