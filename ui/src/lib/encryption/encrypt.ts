export async function encryptPaste(content: string, password: string) {
	const encodedPassword = new TextEncoder().encode(password);
	const passwordHash = await crypto.subtle.digest('SHA-256', encodedPassword);

	// Encrypt content
	const encodedContent = new TextEncoder().encode(content);
	const iv = crypto.getRandomValues(new Uint8Array(12));
	const key = await crypto.subtle.importKey('raw', passwordHash, {
		name: 'AES-GCM'
	}, false, ['encrypt']);
	const encrypted = await crypto.subtle.encrypt({
		name: 'AES-GCM',
		iv
	}, key, encodedContent);

// 	return iv and encrypted as stings
	return {
		iv: String.fromCharCode(...iv),
		encrypted: String.fromCharCode(...new Uint8Array(encrypted))
	};
}