<script>
    import { register } from "../../api.js";
    import { navigate } from "svelte-routing";

    let username = "";
    let password = "";
    let email = "";
    let error = "";

    const doRegister = async () => {
        try {
            const tokenData = await register({ email, username, password });
            navigate("/");
        } catch (error) {
            error = "We had trouble registering you";
        }
    };
</script>

<section class="hero is-primary">
    <div class="hero-body">
        <p class="title">Register</p>
    </div>
</section>

<div class="container">
    <form on:submit|preventDefault={doRegister}>
        <div class="field">
            <label for="email" class="label">Email</label>
            <div class="control">
                <input
                    id="email"
                    class="input"
                    type="text"
                    placeholder="Email"
                    bind:value={email}
                    required
                />
            </div>
        </div>
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
