<script>
    import { onMount } from "svelte";
    import { getPostSearchTag } from "../api.js";
    import { Link, navigate } from "svelte-routing";
    import InfiniteScroll from "svelte-infinite-scroll";
    import TagLink from "../TagLink.svelte";
    import queryString from "query-string";
    import Tags from "svelte-tags-input";
    import { add_attribute } from "svelte/internal";

    export let location;

    let searchTerms = [];

    let page = 1;
    let posts = [];
    let newBatch = [];

    const splitToChunks = (array, parts) => {
        let result = [];
        for (let i = 0; i < parts; i++) {
            let currentColumn = [];
            for (let j = i; j < array.length; j += parts) {
                currentColumn.push(array[j]);
            }
            result.push(currentColumn);
        }
        return result;
    };

    let postChunks = [];
    // split posts into 4 columns
    $: {
        postChunks = splitToChunks(posts, 5);
    }

    const getData = async () => {
        const data = await getPostSearchTag({ page, q: searchTerms.join("+") });
        if (data.posts) {
            newBatch = data.posts;
        } else {
            newBatch = [];
        }
    };
    $: {
        posts = [...posts, ...newBatch];
    }
    let queryParams;

    const onTagChange = (value) => {
        searchTerms = value.detail.tags;
    };

    $: {
        queryParams = queryString.parse(location.search);
        if (queryParams.tags) {
            searchTerms = queryParams.tags.split(" ");
        } else {
            searchTerms = [];
        }
        posts = [];
        getData();
    }

    const onSearch = (i) => {
        if (searchTerms.length > 0) {
            navigate(`/posts?tags=${searchTerms.join("+")}`);
        } else {
            navigate(`/posts`);
        }
    };
</script>

<section class="hero is-primary">
    <div class="hero-body">
        <p class="title">Posts</p>
    </div>
</section>

<section class="section">
    <div class="container">
        <div class="block">
            <form on:submit|preventDefault={onSearch}>
                <div class="field has-addons">
                    <div class="control is-expanded">
                        <div class="control" id="tags">
                            <Tags
                                tags={searchTerms}
                                addKeys={[9, 32]}
                                on:tags={onTagChange}
                            />
                        </div>
                    </div>
                    <div class="control">
                        <button type="submit" class="button is-primary">
                            Search
                        </button>
                    </div>
                </div>
            </form>
        </div>
        <div class="block">
            <div class="columns">
                {#each postChunks as postChunk}
                    <div class="column is-one-fifth">
                        {#each postChunk as post, i (post.id)}
                            <div class="block">
                                <div class="card">
                                    <div class="card-image">
                                        <figure class="image">
                                            <Link to="/post/{post.id}">
                                                <img
                                                    alt={post.id}
                                                    src={post.image_path}
                                                />
                                            </Link>
                                        </figure>
                                    </div>
                                    <div class="card-content">
                                        {#if post.tags}
                                            {#each post.tags as tag (tag)}
                                                <TagLink {tag} />
                                            {/each}
                                        {:else}
                                            <TagLink tag="tagme" />
                                        {/if}
                                    </div>
                                </div>
                            </div>
                        {/each}
                    </div>
                {/each}
            </div>
            <InfiniteScroll
                hasMore={newBatch.length}
                elementScroll={document}
                on:loadMore={() => {
                    page++;
                    getData();
                }}
            />
        </div>
        {#if newBatch.length == 0}
            <div class="notification is-primary">
                <p class="has-text-centered">End of posts</p>
            </div>
        {/if}
    </div>
</section>
