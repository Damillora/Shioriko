<script>
    import { token } from "./stores.js";
    import { isTokenExpired } from "./login-check.js";

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = !isTokenExpired(value);
    });
</script>

{#if loggedIn == true}
    <slot />
{/if}
