<script>
    import { token } from "$lib/stores";
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { isTokenExpired } from "$lib/login-check";

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = !isTokenExpired(value);
    });

    onMount(() => {
        if (loggedIn === false) {
            goto("/auth/login");
        }
    });
</script>
