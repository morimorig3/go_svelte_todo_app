<script>
	import Footer from '$lib/components/Footer.svelte';
	import Header from '$lib/components/Header.svelte';
	import '../app.css';
</script>

<div class="">
	<Header />
	<main>
		メイン
		<slot />
	</main>
	<Footer />
</div>

<style lang="postcss">
	:global(html) {
	}
</style>
