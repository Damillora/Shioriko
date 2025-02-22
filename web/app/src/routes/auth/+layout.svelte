<script lang="ts">
    import Navbar from "$lib/components/ui/Navbar.svelte";
    interface Props {
        children?: import('svelte').Snippet;
    }

    let { children }: Props = $props();

    export const ssr = false;
</script>


<section class="hero is-primary is-fullheight-with-navbar">
    <div class="hero-body">
        {@render children?.()}
    </div>
</section>
