<div class="flex flex-col gap-y-5">
	<div class="floating-wrap h-12 w-full">
		<input
			class="floating-input h-full w-full p-2 pt-6"
			id="taskInput"
			type="text"
			placeholder=" "
		/>
		<label class="floating-label" for="taskInput">タスク名を入力</label>
	</div>
	<button class="w-full h-12 font-bold text-sky-50 bg-sky-400 hover:opacity-80 transition-colors"
		>追加ボタン</button
	>
</div>

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
	}
	.floating-label {
		transition: 0.2s;
	}
	.floating-input:not(:placeholder-shown) + .floating-label {
		color: theme(colors.slate.400);
		transform: scale(0.6);
		transform-origin: left top;
	}
	.floating-input:focus + .floating-label {
		color: theme(colors.sky.600);
		transform: scale(0.6);
		transform-origin: left top;
	}
</style>
