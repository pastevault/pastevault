export async function decryptPaste(paste: string, password: string) {
	const pasteBuffer = Uint8Array.from(atob(paste), c => c.charCodeAt(0));
	const iv = pasteBuffer.slice(0, 12);
	const encrypted = pasteBuffer.slice(12);
	const encodedPassword = new TextEncoder().encode(password);
	const passwordHash = await crypto.subtle.digest('SHA-256', encodedPassword);
	const key = await crypto.subtle.importKey('raw', passwordHash, {
		name: 'AES-GCM'
	}, false, ['decrypt']);
	try {
		const decrypted = await crypto.subtle.decrypt({
			name: 'AES-GCM',
			iv
		}, key, encrypted);
		return new TextDecoder().decode(decrypted);
	} catch (e) {
		return "";
	}
}