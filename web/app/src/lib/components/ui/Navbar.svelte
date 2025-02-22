<script>
    import { token } from "$lib/stores";
    import { isTokenExpired } from "$lib/login-check";

    let menu_shown = $state(false);
    
    let loggedIn = $state(false);
    token.subscribe((value) => {
        loggedIn = !isTokenExpired(value);
    });

    const toggleMenu = () => {
        menu_shown = !menu_shown;
    };
</script>

<nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
        <a class="navbar-item" href="/">Shioriko</a>

        <a
            href={"#"}
            onclick={toggleMenu}
            role="button"
            class="navbar-burger"
            aria-label="menu"
            aria-expanded="false"
        >
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
        </a>
    </div>

    <div class="navbar-menu" class:is-active={menu_shown}>
        <div class="navbar-start">
            <a class="navbar-item" href="/posts">Posts</a>
            <a class="navbar-item" href="/tags">Tags</a>
            {#if loggedIn}
                <a class="navbar-item" href="/upload">Upload</a>
            {/if}
        </div>

        <div class="navbar-end">
            {#if loggedIn}
                <div class="navbar-item">
                    <div class="buttons">
                        <a href="/user/profile" class="button is-primary">
                            Profile
                        </a>
                        <a href="/auth/logout" class="button is-light">
                            Log out
                        </a>
                    </div>
                </div>
            {:else}
                <div class="navbar-item">
                    <div class="buttons">
                        <a href="/auth/register" class="button is-primary">
                            Register
                        </a>
                        <a href="/auth/login" class="button is-light">
                            Log in
                        </a>
                    </div>
                </div>
            {/if}
        </div>
    </div>
</nav>
