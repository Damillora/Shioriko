<script>
    import { token } from "./stores.js";
    import { navigate } from "svelte-routing";
    import { onMount } from "svelte";
    import { isTokenExpired } from "./login-check.js";

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = !isTokenExpired(value);
    });

    onMount(() => {
        if (loggedIn === false) {
            navigate("/auth/login");
        }
    });
</script>
