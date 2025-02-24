<script lang="ts">
    import Navbar from "$lib/components/ui/Navbar.svelte";
    interface Props {
        children?: import('svelte').Snippet;
    }

    let { children }: Props = $props();
</script>


<section class="hero is-fullheight-with-navbar">
    <div class="hero-body">
        {@render children?.()}
    </div>
</section>
