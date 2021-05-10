<script>
    import { token } from "./stores.js";
    import { navigate } from "svelte-routing";
    import { onMount } from "svelte";

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = value !== "";
    });

    onMount(() => {
        if (loggedIn === false) {
            navigate("/auth/login");
        }
    });
</script>
