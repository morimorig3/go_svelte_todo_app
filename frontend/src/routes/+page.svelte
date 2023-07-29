<script lang="ts">
	import InputTask from '$lib/components/InputTask.svelte';
	import TaskList from '$lib/components/TaskList.svelte';
	import type { Task } from '$lib/types';
	import { onMount } from 'svelte';

	let tasks: Task[] = [];

	async function loadTasks() {
		tasks = await fetch(`/api/task/list`).then((res) => res.json());
	}

	onMount(() => {
		loadTasks();
	});
</script>

<div class="flex flex-col gap-y-10">
	<InputTask {loadTasks} />
	{#if tasks.length > 0}
		<TaskList {tasks} {loadTasks} />
	{/if}
</div>
