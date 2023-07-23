<script>
	import '../app.css';
</script>

<div class="">
	<header>ヘッダー</header>
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
