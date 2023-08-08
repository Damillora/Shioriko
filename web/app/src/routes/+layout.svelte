<script>
    import "../app.scss";
    import Navbar from "$lib/components/Navbar.svelte";

    export const ssr = false;
</script>

<svelte:head>
    <title>Shioriko</title>
</svelte:head>

<Navbar />

<slot />
