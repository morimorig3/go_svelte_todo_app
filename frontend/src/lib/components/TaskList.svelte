<script lang="ts">
	import type { Task as TaskType } from '$lib/types';
	import Task from './Task.svelte';

	export let loadTasks: () => Promise<void>;
	export let tasks: TaskType[] = [];
</script>

<ul>
	{#each tasks as task (task.id)}
		<Task title={task.title} id={task.id} {loadTasks} />
	{/each}
</ul>
