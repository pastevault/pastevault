<script lang="ts">
	import type { PageData } from './$types';
	import dayjs from 'dayjs';
	import relativeTime from 'dayjs/plugin/relativeTime';
	import { onMount } from 'svelte';
	import { decryptPaste } from '$lib/encryption/decrypt';

	export let data: PageData;

	dayjs.extend(relativeTime);

	let expires_in: string;
	const expirationDifference = dayjs(data.paste.expiration).diff(dayjs(), 'minutes');

	if (expirationDifference < 10) {
		expires_in = 'soon';
	} else {
		expires_in = dayjs(data.paste.expiration).fromNow();
	}

	onMount(async () => {
		if (data.paste.encrypted) {
			const password = window.location.hash.substring(1);
			if (password) {
				data.paste.content = await decryptPaste(data.paste.content, password);
			}
		}
	});
</script>

<div>
	{data.paste.content}

	<div>
		Expires in: {expires_in}
	</div>
</div>