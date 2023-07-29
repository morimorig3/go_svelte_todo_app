<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import InputTask from '$lib/components/InputTask.svelte';
	import TaskList from '$lib/components/TaskList.svelte';
	import type { Task } from '$lib/types';

	let taskRequest = new Promise<Task[]>(() => {});

	afterNavigate(() => {
		taskRequest = fetch(`/api/task/list`).then((res) => res.json());
	});
</script>

<div class="flex flex-col gap-y-10">
	<InputTask />
	{#await taskRequest then tasks}
		{#if tasks.length > 0}
			<TaskList {tasks} />
		{/if}
	{/await}
</div>
