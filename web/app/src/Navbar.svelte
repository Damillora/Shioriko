<script>
    import { Link } from "svelte-routing";
    import { token } from "./stores.js";

    let menu_shown = false;

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = value !== "";
    });

    const toggleMenu = () => {
        menu_shown = !menu_shown;
    };
</script>

<nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
        <Link class="navbar-item" to="/">Shioriko</Link>

        <a
            href={"#"}
            on:click={toggleMenu}
            role="button"
            class="navbar-burger"
            aria-label="menu"
            aria-expanded="false"
        >
            <span aria-hidden="true" />
            <span aria-hidden="true" />
            <span aria-hidden="true" />
        </a>
    </div>

    <div class="navbar-menu" class:is-active={menu_shown}>
        <div class="navbar-start">
            <Link class="navbar-item" to="/posts">Posts</Link>
            {#if loggedIn}
                <Link class="navbar-item" to="/upload">Upload</Link>
            {/if}
        </div>

        <div class="navbar-end">
            {#if loggedIn}
                <div class="navbar-item">
                    <div class="buttons">
                        <Link to="/auth/logout" class="button is-light">
                            Log out
                        </Link>
                    </div>
                </div>
            {:else}
                <div class="navbar-item">
                    <div class="buttons">
                        <Link to="/auth/register" class="button is-primary">
                            <strong>Register</strong>
                        </Link>
                        <Link to="/auth/login" class="button is-light">
                            Log in
                        </Link>
                    </div>
                </div>
            {/if}
        </div>
    </div>
</nav>
