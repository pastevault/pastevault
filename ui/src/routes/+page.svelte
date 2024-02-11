<script lang="ts">
	import type { PageData } from './$types';
	import { superForm } from 'sveltekit-superforms/client';
	import SuperDebug from 'sveltekit-superforms/client/SuperDebug.svelte';
	import { encryptPaste } from '$lib/encryption/encrypt';
	import { onMount } from 'svelte';


	export let data: PageData;
	let paste: string;
	let disabled = true;

	onMount(() => {
		disabled = false;
	});

	const { form, errors, constraints, message, enhance } = superForm(data.form, {
		onSubmit: async (form) => {
			if (form.formData.get('encrypted') === 'on') {
				const password = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
				const content = form.formData.get('content') as string;
				paste = content;
				const encryptedContent = await encryptPaste(content, password);

				form.formData.set('content', btoa(encryptedContent.iv + encryptedContent.encrypted));
			}
		},
		onResult: async ({ result }) => {
			if (result.type === 'failure') {
				if ($form.encrypted) {
					if (result.data) {
						result.data.form.data.content = paste;
					}
				}
			}
		}
	});
</script>


<h1>{data.title}</h1>

{#if $message}<h3>{$message}</h3>{/if}

<form method="POST" use:enhance>
	<label for="content">Content</label>
	{#if $errors.content}<span class="invalid">{$errors.content}</span>{/if}
	<input
		type="text" name="content"
		aria-invalid={$errors.content ? 'true' : undefined}
		bind:value={$form.content}
		{...$constraints.content} />

	<label for="expiration">Expiration</label>
	{#if $errors.expiration}<span class="invalid">{$errors.expiration}</span>{/if}
	<input
		type="text" name="expiration"
		aria-invalid={$errors.expiration ? 'true' : undefined}
		bind:value={$form.expiration}
		{...$constraints.expiration} />

	<label for="encrypted">
		Encrypted {#if disabled} (disabled: js required) {/if}
	</label>
	<input
		type="checkbox" name="encrypted"
		bind:checked={$form.encrypted}
		{...$constraints.encrypted}
		disabled={disabled} />

	<div>
		<button>Submit</button>
	</div>

	<SuperDebug data={$form} />
</form>

<style>
    .invalid {
        color: red;
    }
</style>