<script>
    import { login } from "../api.js";
    import { navigate } from "svelte-routing";

    let username = "";
    let password = "";

    const doLogin = async () => {
        const tokenData = await login({ username, password });
        navigate("/");
    };
</script>

<div class="container">
    <form on:submit|preventDefault={doLogin}>
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
        <div class="field">
            <div class="control">
                <button class="button is-link">Login</button>
            </div>
        </div>
    </form>
</div>
