<script>
	import Header from '$lib/components/Header.svelte';
import '../app.css';
</script>

<div class="">
	<Header />
	<main>
		メイン
		<slot />
	</main>
	<footer>フッター</footer>
</div>

<style lang="postcss">
	:global(html) {
	}
</style>
