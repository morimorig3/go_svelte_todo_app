<script lang="ts">
	import { CircleCheckRegular } from 'svelte-awesome-icons';

	export let id: string = '';
	export let title: string = '';
</script>

<li class="bg-white flex hover:shadow-md transition-shadow h-20 border p-4 items-center gap-x-2">
	<button><CircleCheckRegular class="transition-colors hover:text-green-400" /></button>
	<p class="font-bold">{title}{id}</p>
</li>
