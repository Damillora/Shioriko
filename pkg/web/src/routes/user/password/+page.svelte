<script>
    import { updateToken, updateUserPassword,  } from "$lib/api";
    import UserActionsPanel from "$lib/components/panels/UserActionsPanel.svelte";
    import { onMount } from "svelte";

    let loading = $state(false);
    let updated = $state(false);
    let form = $state({
        old_password: "",
        new_password: "",
    });

    const getData = async () => {
        loading = false;
    };

    const submitForm = async () => {
        updated = false;
        await updateUserPassword({ old_password: form.old_password, new_password: form.new_password })
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
                    <h1 class="title">Change Password</h1>
                    {#if updated}
                    <div class="notification is-success">
                        Password updated!
                    </div>
                    {/if}
                    <form onsubmit={submitForm}>
                        <div class="field">
                            <label class="label" for="old_password">Current Password</label>
                            <div class="control">
                                <input
                                    id="old_password"
                                    class="input"
                                    class:is-skeleton="{loading}"
                                    type="password"
                                    placeholder="Current password"
                                    bind:value={form.old_password}
                                />
                            </div>
                        </div>
                        <div class="field">
                            <label class="label" for="new_password">New Password</label>
                            <div class="control">
                                <input
                                    id="new_password"
                                    class="input"
                                    class:is-skeleton="{loading}"
                                    type="password"
                                    placeholder="New password"
                                    bind:value={form.new_password}
                                />
                            </div>
                        </div>
                        <div class="field">
                            <button class="button is-primary is-fullwidth is-outlined" type="submit">Change Password</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</section>
