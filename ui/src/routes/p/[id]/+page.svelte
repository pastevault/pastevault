<script lang="ts">
	import type { PageData } from './$types';
	import dayjs from 'dayjs';
	import relativeTime from 'dayjs/plugin/relativeTime';
	import { onMount } from 'svelte';
	import { decryptPaste } from '$lib/encryption/decrypt';
	import type { Paste__Output } from '$lib/proto/proto/Paste';

	export let data: PageData;
	let paste = data.paste as Paste__Output;
	let passwordNeeded = false;
	let password: string;

	dayjs.extend(relativeTime);

	let expires_in: string;
	const expirationDifference = dayjs(paste.expiration).diff(dayjs(), 'minutes');

	if (expirationDifference < 10) {
		expires_in = 'soon';
	} else {
		expires_in = dayjs(paste.expiration).fromNow();
	}

	async function decrypt() {
		password = window.location.hash.substring(1);
		if (password) {
			let decrypted = await decryptPaste(paste.content, password);
			console.log(decrypted);
			if (decrypted.length > 0) {
				paste.content = decrypted;
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
		if (paste.encrypted) {
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
		{paste.content}
		<a href="http://localhost:8080/v1/paste/{paste.id}/raw">Raw</a>
	{/if}

	<div>
		Expires in: {expires_in}
	</div>
</div>