<script>
    import { onMount } from "svelte";
    import { getUserProfile } from "$lib/api";
    import AuthRequired from "$lib/components/AuthRequired.svelte";

    let user;

    const getData = async () => {
        user = await getUserProfile();
    };
    onMount(() => {
        getData();
    });
</script>

<AuthRequired />

<section class="section">
    <div class="container">
        {#if user}
            <h1 class="title">Welcome, {user.username}</h1>
            <p>Email: {user.email}</p>
            <p>Username: {user.username}</p>
        {/if}
    </div>
</section>
