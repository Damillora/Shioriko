<script>
    import { login } from "$lib/api";
    import { goto } from "$app/navigation";

    let username = $state("");
    let password = $state("");
    let error = $state("");

    const doLogin = async (e) => {
        e.preventDefault();
        error = "";
        try {
            const tokenData = await login({ username, password });
            goto("/");
        } catch (e) {
            error = "We had trouble logging you in";
            return;
        }
    };
</script>

<section class="hero is-primary">
    <div class="hero-body">
        <p class="title">Login</p>
    </div>
</section>

<section class="section">
    <div class="container">
        <form onsubmit={doLogin}>
            <div class="field">
                <label for="username" class="label">Username</label>
                <div class="control">
                    <input
                        id="username"
                        class="input"
                        type="text"
                        placeholder="Username"
                        bind:value={username}
                        required
                    />
                </div>
            </div>
            <div class="field">
                <label for="password" class="label">Password</label>
                <div class="control">
                    <input
                        id="password"
                        class="input"
                        type="password"
                        placeholder="Password"
                        bind:value={password}
                        required
                    />
                </div>
            </div>
            {#if error}
                <div class="field">
                    <p class="has-text-danger">{error}</p>
                </div>
            {/if}
            <div class="field">
                <div class="control">
                    <button class="button is-link">Login</button>
                </div>
            </div>
        </form>
    </div>
</section>
