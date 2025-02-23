<script>
    import { getUserProfile, updateToken, updateUserProfile } from "$lib/api";
    import UserActionsPanel from "$lib/components/panels/UserActionsPanel.svelte";
    import { onMount } from "svelte";

    let loading = $state(false);
    let updated = $state(false);
    let form = $state({
        email: "",
        username: "",
    });

    const getData = async () => {
        const user = await getUserProfile();
        form.email = user.email;
        form.username = user.username;
        loading = false;
    };

    const submitForm = async () => {
        updated = false;
        await updateUserProfile({ email: form.email, username: form.username })
        await updateToken();
        updated = true;
    };

    onMount(() => {
        loading = true;
        getData();
    });
</script>

<section class="section">
    <div class="container">
        <div class="columns">
            <div class="column is-one-third">
                <UserActionsPanel />
            </div>
            <div class="column is-two-thirds">
                <div class="box">
                    <h1 class="title">Profile</h1>
                    {#if updated}
                    <div class="notification is-success">
                        Profile updated!
                    </div>
                    {/if}
                    <form onsubmit={submitForm}>
                        <div class="field">
                            <label class="label" for="email">Email</label>
                            <div class="control">
                                <input
                                    id="email"
                                    class="input"
                                    class:is-skeleton="{loading}"
                                    type="text"
                                    placeholder="Email"
                                    bind:value={form.email}
                                />
                            </div>
                        </div>
                        <div class="field">
                            <label class="label" for="username">Username</label>
                            <div class="control">
                                <input
                                    id="username"
                                    class="input"
                                    class:is-skeleton="{loading}"
                                    type="text"
                                    placeholder="Username"
                                    bind:value={form.username}
                                />
                            </div>
                        </div>
                        <div class="field">
                            <button class="button is-primary is-fullwidth is-outlined" type="submit">Update</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</section>
