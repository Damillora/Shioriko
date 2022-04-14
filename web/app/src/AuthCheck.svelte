<script>
    import { token } from "./stores.js";

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = value !== "";
    });
</script>

{#if loggedIn == true}
    <slot />
{/if}
