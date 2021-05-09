<script>
    import { token } from "../stores.js"
    import { navigate } from "svelte-routing"
    import {onMount } from "svelte"

    onMount(() => {
        token.set("");
        navigate("/");
    });

</script>
