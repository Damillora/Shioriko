<script>
    import { onMount } from "svelte";
    import { Link } from "svelte-routing";

    export let posts = [];
</script>

<div class="columns is-multiline">
    {#each posts as post, i (post.id)}
        <div class="column is-one-quarter">
            <div class="block">
                <div class="card">
                    <div class="card-image">
                        <figure class="image">
                            <Link to="/post/{post.id}">
                                <img alt={post.id} src={post.thumbnail_path} />
                            </Link>
                        </figure>
                    </div>
                </div>
            </div>
        </div>
    {/each}
</div>
