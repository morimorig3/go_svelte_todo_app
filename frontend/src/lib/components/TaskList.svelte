<script lang="ts">
	import type { Task as TaskType } from '$lib/types';
	import Task from './Task.svelte';

	export let tasks: TaskType[] = [];
</script>

{#if tasks.length > 0}
	<ul>
		{#each tasks as { id, title }, i (`${id}${i}`)}
			<Task {id} {title} />
		{/each}
	</ul>
{:else}{/if}
