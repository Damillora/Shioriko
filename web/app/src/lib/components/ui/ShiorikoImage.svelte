<script>
    import { onMount } from "svelte";

    let { alt, src } = $props();
    let loading = $state(false);

	onMount(() => {
			const img = new Image();
			img.src = src;
			loading = true;

			img.onload = () => {
					loading = false;
			};
			img.onerror = () => {
					loading = true;
			};
	})
</script>

<figure class:is-skeleton="{loading}">
	<img {src} {alt} />
</figure>
