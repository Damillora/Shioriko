<script lang="ts">
    import "../app.scss";
    import Navbar from "$lib/components/ui/Navbar.svelte";
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

<footer class="footer">
    <div class="content has-text-centered">
      <p>
        <strong><a href="https://github.com/Damillora/Shioriko">Shioriko</a></strong>: a booru-style image gallery written in Go and Svelte
      </p>
    </div>
  </footer>