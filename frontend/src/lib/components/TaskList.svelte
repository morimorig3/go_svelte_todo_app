<script lang="ts">
	import type { Task as TaskType } from '$lib/types';
	import Task from './Task.svelte';

	export let tasks: TaskType[] = [];
	$: console.log(tasks);
</script>

<ul>
	{#each tasks as task (task.id)}
		<Task title={task.title} id={task.id} />
	{/each}
</ul>
