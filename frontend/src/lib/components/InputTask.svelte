<script lang="ts">
	import { enhance } from '$app/forms';

	export let loadTasks: () => Promise<void>;
</script>

<form
	method="POST"
	action="?/add"
	use:enhance={() => {
		return async ({ update }) => {
			await update();
			await loadTasks();
		};
	}}
>
	<div class="flex flex-col gap-y-5">
		<div class="floating-wrap h-12 w-full">
			<input
				class="floating-input h-full w-full p-2 pt-6"
				id="taskInput"
				type="text"
				placeholder=" "
				name="taskTitle"
			/>
			<label class="floating-label" for="taskInput">タスク名を入力</label>
		</div>
		<button class="w-full h-12 font-bold text-sky-50 bg-sky-400 hover:opacity-80 transition-colors"
			>追加ボタン</button
		>
	</div>
</form>

<style lang="postcss">
	.floating-wrap {
		position: relative;
	}
	.floating-label {
		position: absolute;
		display: grid;
		place-items: center;
		top: 0;
		left: 10px;
		bottom: 0;
		transition: 0.2s;
		color: theme(colors.slate.400);
	}
	.floating-input:not(:placeholder-shown) + .floating-label {
		transform: scale(0.6);
		transform-origin: left top;
	}
	.floating-input:focus + .floating-label {
		color: theme(colors.sky.600);
		transform: scale(0.6);
		transform-origin: left top;
	}
</style>
