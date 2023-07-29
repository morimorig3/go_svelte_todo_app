<script lang="ts">
	import { CircleCheckRegular } from 'svelte-awesome-icons';

	export let title: string = '';
	export let id: string = '';
</script>

<li>
	<form
		class="bg-white flex hover:shadow-md transition-shadow h-20 border p-4 items-center gap-x-2"
		method="POST"
		action="?/remove"
	>
		<input type="hidden" value={id} name="taskId" />
		<button><CircleCheckRegular class="transition-colors hover:text-green-400" /></button>
		<p class="font-bold">{title}</p>
	</form>
</li>
