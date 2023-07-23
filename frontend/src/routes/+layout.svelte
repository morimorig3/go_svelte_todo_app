<script>
	import Footer from '$lib/components/Footer.svelte';
	import Header from '$lib/components/Header.svelte';
	import '../app.css';
</script>

<div class="min-h-screen flex flex-col">
	<Header />
	<main class="flex-1">
		<slot />
	</main>
	<Footer />
</div>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.sky.50);
		color: theme(colors.slate.900);
	}
	:global(a) {
		color: theme(colors.slate.900);
	}
</style>
