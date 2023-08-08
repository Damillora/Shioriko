<script>
    import { token } from "$lib/stores";
    import { isTokenExpired } from "$lib/login-check";

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = !isTokenExpired(value);
    });
</script>

{#if loggedIn == true}
    <slot />
{/if}
