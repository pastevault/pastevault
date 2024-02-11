<script lang="ts">
	import type { PageData } from './$types';
	import dayjs from 'dayjs';
	import relativeTime from 'dayjs/plugin/relativeTime';
	import { onMount } from 'svelte';
	import { decryptPaste } from '$lib/encryption/decrypt';

	export let data: PageData;
	let passwordNeeded = false;
	let password: string;

	dayjs.extend(relativeTime);

	let expires_in: string;
	const expirationDifference = dayjs(data.paste.expiration).diff(dayjs(), 'minutes');

	if (expirationDifference < 10) {
		expires_in = 'soon';
	} else {
		expires_in = dayjs(data.paste.expiration).fromNow();
	}

	async function decrypt() {
		password = window.location.hash.substring(1);
		if (password) {
			let decrypted = await decryptPaste(data.paste.content, password);
			console.log(decrypted);
			if (decrypted.length > 0) {
				data.paste.content = decrypted;
			} else {
				passwordNeeded = true;
			}
		} else {
			passwordNeeded = true;
		}
	}

	function includeHash() {
		window.location.hash = password;
		decrypt();
		passwordNeeded = false;
	}

	onMount(async () => {
		if (data.paste.encrypted) {
			await decrypt();
		}
	});
</script>

<div>

	{#if passwordNeeded}
		<div>
			This paste is encrypted. Please enter the password to decrypt it.
		</div>
		<form on:submit|preventDefault={includeHash}>
			<input type="password" placeholder="Password" bind:value={password} />
			<button type="submit">Decrypt</button>
		</form>
		{:else}
		{data.paste.content}
	{/if}

	<div>
		Expires in: {expires_in}
	</div>
</div>