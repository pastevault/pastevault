<script lang="ts">
	import type { PageData } from './$types';
	import { superForm } from 'sveltekit-superforms/client';
	import SuperDebug from 'sveltekit-superforms/client/SuperDebug.svelte';


	export let data: PageData;

	const { form, errors, constraints, message } = superForm(data.form);
</script>


<h1>{data.title}</h1>

{#if $message}<h3>{$message}</h3>{/if}

<form method="POST">
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