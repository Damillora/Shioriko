<script>
    import { onMount } from "svelte";
    import { getPostsTag } from "../api.js";
    import { Link } from "svelte-routing";
    import queryString from "query-string";

    export let location;

    export let id;

    let page = 1;
    let totalPages = 1;
    let posts = [];
    const getData = async () => {
        const data = await getPostsTag({ page, tag: id });
        if (Array.isArray(data.posts)) {
            posts = data.posts;
            totalPages = data.totalPage;
        }
    };
    onMount(() => {
        let queryParams;
        queryParams = queryString.parse(location.search);
        if (queryParams.page) {
            page = parseInt(queryParams.page);
        }
        getData();
    });

    const handlePage = (i) => {
        return () => {
            page = i;
            getData();
        };
    };
</script>

<section class="hero is-primary">
    <div class="hero-body">
        <p class="title">
            {id}
        </p>
        <p class="subtitle">Tag</p>
    </div>
</section>

<section class="section">
    <div class="container">
        <nav class="pagination" role="navigation" aria-label="pagination">
            {#if page > 1}
                <Link
                    on:click={handlePage(page - 1)}
                    to="/tag/{id}?page={page - 1}"
                    class="pagination-previous"
                    aria-label="Previous">Previous</Link
                >
            {/if}
            {#if page < totalPages}
                <Link
                    on:click={handlePage(page + 1)}
                    to="/tag/{id}?page={page + 1}"
                    class="pagination-next"
                    aria-label="Next">Next</Link
                >
            {/if}
            <ul class="pagination-list">
                {#if page > 3}
                    <li>
                        <Link
                            on:click={handlePage(1)}
                            to="/tag/{id}?page={1}"
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
                                    to="/tag/{id}?page={i}"
                                    class="pagination-link is-current"
                                    aria-label="Goto page {i}">{i}</Link
                                >
                            </li>
                        {:else}
                            <li>
                                <Link
                                    on:click={handlePage(i)}
                                    to="/tag/{id}?page={i}"
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
                            to="/tag/{id}?page={totalPages}"
                            class="pagination-link"
                            aria-label="Goto page {totalPages}"
                            >{totalPages}</Link
                        >
                    </li>
                {/if}
            </ul>
        </nav>
        <div class="columns is-multiline">
            {#each posts as post (post.id)}
                <div class="column is-one-quarter card">
                    <div class="card-image">
                        <figure class="image">
                            <Link to="/post/{post.id}">
                                <img alt={post.id} src={post.image_path} />
                            </Link>
                        </figure>
                    </div>
                    <div class="card-content">
                        <div class="content">
                            {#each post.tags as tag (tag)}
                                <p>
                                    <Link to="/tag/{tag}">{tag}</Link>
                                </p>
                            {/each}
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </div>
</section>
