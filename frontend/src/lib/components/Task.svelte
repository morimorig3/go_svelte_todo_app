<script lang="ts">
	import { CircleCheckRegular } from 'svelte-awesome-icons';

	import { enhance } from '$app/forms';

	export let loadTasks: () => Promise<void>;
	export let title: string = '';
	export let id: string = '';
</script>

<li>
	<form
		class="bg-white flex hover:shadow-md transition-shadow h-20 border p-4 items-center gap-x-2"
		method="POST"
		action="?/remove"
		use:enhance={() => {
			return async ({ update }) => {
				await update();
				await loadTasks();
			};
		}}
	>
		<input type="hidden" value={id} name="taskId" />
		<button><CircleCheckRegular class="transition-colors hover:text-green-400" /></button>
		<p class="font-bold">{title}</p>
	</form>
</li>
