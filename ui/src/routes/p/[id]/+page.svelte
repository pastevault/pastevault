<script lang="ts">
	import type { PageData } from './$types';
	import dayjs from 'dayjs';
	import relativeTime from 'dayjs/plugin/relativeTime';

	export let data: PageData;

	dayjs.extend(relativeTime)

	let expires_in: string;
	const expirationDifference = dayjs(data.paste.expiration).diff(dayjs(), 'minutes');

	if (expirationDifference < 10) {
		expires_in = "soon";
	} else {
		expires_in = dayjs(data.paste.expiration).fromNow();
	}
</script>

<div>
	{data.paste.content}

	<div>
		Expires in: {expires_in}
	</div>
</div>