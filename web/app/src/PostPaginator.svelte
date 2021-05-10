<script>
    import { Link } from "svelte-routing";
    import TagLink from "./TagLink.svelte";

    export let posts = [];
    let postChunks = [];
    // split posts into 4 columns
    $: postChunks = Array(Math.min(posts.length, 4))
        .fill()
        .map(function (_, i) {
            let chunkSize = Math.floor(posts.length / 4);
            if (chunkSize % 4 > i + 1) {
                chunkSize += 1;
            }
            chunkSize = Math.max(chunkSize, 1);

            return posts.slice(i * chunkSize, i * chunkSize + chunkSize);
        });
    export let page = 1;
    export let totalPages = 1;
    export let handlePage = (i) => {};
    export let url = "/posts";
</script>

<nav class="pagination" role="navigation" aria-label="pagination">
    {#if page > 1}
        <Link
            on:click={handlePage(page - 1)}
            to="{url}page={page - 1}"
            class="pagination-previous"
            aria-label="Previous">Previous</Link
        >
    {/if}
    {#if page < totalPages}
        <Link
            on:click={handlePage(page + 1)}
            to="{url}page={page + 1}"
            class="pagination-next"
            aria-label="Next">Next</Link
        >
    {/if}
    <ul class="pagination-list">
        {#if page > 3}
            <li>
                <Link
                    on:click={handlePage(1)}
                    to="{url}page={1}"
                    class="pagination-link"
                    aria-label="Goto page 1">1</Link
                >
            </li>
            <li>
                <span class="pagination-ellipsis">&hellip;</span>
            </li>
        {/if}
        {#each [...Array(5).keys()].map((x) => x + page - 2) as i}
            {#if i >= 1 && i <= totalPages}
                {#if i == page}
                    <li>
                        <Link
                            on:click={handlePage(i)}
                            to="{url}page={i}"
                            class="pagination-link is-current"
                            aria-label="Goto page {i}">{i}</Link
                        >
                    </li>
                {:else}
                    <li>
                        <Link
                            on:click={handlePage(i)}
                            to="{url}page={i}"
                            class="pagination-link"
                            aria-label="Goto page {i}">{i}</Link
                        >
                    </li>
                {/if}
            {/if}
        {/each}
        {#if totalPages - page > 2}
            <li>
                <span class="pagination-ellipsis">&hellip;</span>
            </li>
            <li>
                <Link
                    on:click={handlePage(totalPages)}
                    to="{url}page={totalPages}"
                    class="pagination-link"
                    aria-label="Goto page {totalPages}">{totalPages}</Link
                >
            </li>
        {/if}
    </ul>
</nav>
<div class="tile is-multiline is-ancestor">
    {#each postChunks as postChunk}
        <div class="tile is-parent is-vertical is-3">
            {#each postChunk as post, i (post.id)}
                <div class="tile is-child is-vertical card">
                    <div class="card-image">
                        <figure class="image">
                            <Link to="/post/{post.id}">
                                <img alt={post.id} src={post.image_path} />
                            </Link>
                        </figure>
                    </div>
                    <div class="card-content">
                        {#if post.tags}
                            {#each post.tags as tag (tag)}
                                <TagLink {tag} />
                            {/each}
                        {:else}
                            None
                        {/if}
                    </div>
                </div>
            {/each}
        </div>
    {/each}
</div>
