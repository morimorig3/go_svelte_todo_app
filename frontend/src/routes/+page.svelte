<script lang="ts">
	import InputTask from '$lib/components/InputTask.svelte';
	import TaskList from '$lib/components/TaskList.svelte';
	import type { Task } from '$lib/types';

	const tasks: Task[] = [
		{
			id: '1',
			title: 'タスク1'
		},
		{
			id: '2',
			title: 'タスク2'
		},
		{
			id: '3',
			title: 'タスク3'
		}
	];
</script>

<div class="flex flex-col gap-y-10">
	<InputTask />
	<TaskList {tasks} />
</div>
