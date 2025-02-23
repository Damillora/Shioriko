<script>
    import { token } from "$lib/stores";
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";

    onMount(() => {
        token.set("");
        goto("/");
    });
</script>
