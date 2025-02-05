<script lang="ts">
    import "../app.scss";
    import Navbar from "$lib/components/Navbar.svelte";
    interface Props {
        children?: import('svelte').Snippet;
    }

    let { children }: Props = $props();

    export const ssr = false;
</script>

<svelte:head>
    <title>Shioriko</title>
</svelte:head>

<Navbar />

{@render children?.()}
