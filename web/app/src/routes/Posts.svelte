<script>
    import { onMount } from "svelte";
    import { getPostSearchTag, getTagAutocomplete } from "../api.js";
    import { Link, navigate } from "svelte-routing";
    import InfiniteScroll from "svelte-infinite-scroll";
    import TagLinkNumbered from "../TagLinkNumbered.svelte";
    import queryString from "query-string";
    import Tags from "svelte-tags-input";
    import { add_attribute } from "svelte/internal";
    import { paginate } from "../simple-pagination.js";

    export let location;

    let searchTerms = [];

    let page = 1;
    let totalPages = 1;
    let pagination = [];
    let posts = [];
    let tags = [];
    let categorizedTags = {};

    const getData = async () => {
        const data = await getPostSearchTag({ page, q: searchTerms.join("+") });
        if (data.posts) {
            posts = data.posts;
            tags = data.tags;
            totalPages = data.totalPage;
            pagination = paginate(page, totalPages);
        } else {
            posts = [];
            tags = [];
            totalPages = 0;
            pagination = paginate(page, totalPages);
        }
    };
    $: {
        let catTags = tags.reduce(
            (acc, o) => ((acc[o] = (acc[o] || 0) + 1), acc),
            {}
        );
        categorizedTags = Object.entries(catTags).map(([k, v]) => ({
            tag: k,
            num: v,
        }));
        categorizedTags = categorizedTags.sort((a, b) => b.num - a.num);
    }
    let queryParams;

    const onTagChange = (value) => {
        searchTerms = value.detail.tags;
    };

    const onAutocomplete = async () => {
        const list = await getTagAutocomplete();
        return list;
    };

    $: {
        queryParams = queryString.parse(location.search);
        if (queryParams.tags) {
            searchTerms = queryParams.tags.split(" ");
        } else {
            searchTerms = [];
        }
        posts = [];
        page = 1;
        getData();
    }
    const onSearch = (i) => {
        if (searchTerms.length > 0) {
            navigate(`/posts?tags=${searchTerms.join("+")}`);
        } else {
            navigate(`/posts`);
        }
    };

    const changePage = (i) => {
        page = i;
        getData();
    }
</script>

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
                                autoComplete={onAutocomplete}
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
                <div class="column is-one-third">
                    <div class="panel is-primary">
                        <div class="panel-heading">Tags</div>
                        <div class="panel-block column">
                            <div class="menu">
                                <ul class="menu-list">
                                    {#each categorizedTags as tag (tag)}
                                        <li>
                                            <TagLinkNumbered
                                                class=""
                                                tag={tag.tag}
                                                num={tag.num}
                                            />
                                        </li>
                                    {/each}
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="column  is-two-thirds">
                    <div class="columns is-multiline">
                        <div class="column is-full">
                            <nav
                                class="pagination is-centered"
                                role="navigation"
                                aria-label="pagination"
                            >
                                <a
                                    href={null}
                                    on:click={changePage(page - 1)}
                                    class="pagination-previous">Previous</a
                                >
                                <a
                                    href={null}
                                    on:click={changePage(page + 1)}
                                    class="pagination-next">Next page</a
                                >
                                <ul class="pagination-list">
                                    {#each pagination as pageEntry}
                                        {#if pageEntry == "..."}
                                            <li>
                                                <span
                                                    class="pagination-ellipsis"
                                                    >&hellip;</span
                                                >
                                            </li>
                                        {:else}
                                            <li>
                                                <a
                                                    href={null}
                                                    on:click={() =>
                                                        (changePage(pageEntry))}
                                                    class="pagination-link"
                                                    class:is-current={page ==
                                                        pageEntry}
                                                    aria-label="Goto page {pageEntry}"
                                                    >{pageEntry}</a
                                                >
                                            </li>
                                        {/if}
                                    {/each}
                                </ul>
                            </nav>
                        </div>
                        <div class="column is-full">
                            <div class="columns is-multiline">
                                {#each posts as post, i (post.id)}
                                    <div class="column is-one-third">
                                        <div class="block">
                                            <div class="card">
                                                <div class="card-image">
                                                    <figure class="image">
                                                        <Link
                                                            to="/post/{post.id}"
                                                        >
                                                            <img
                                                                alt={post.id}
                                                                src={post.thumbnail_path}
                                                            />
                                                        </Link>
                                                    </figure>
                                                </div>
                                                <div class="card-content" />
                                            </div>
                                        </div>
                                    </div>
                                {/each}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        {#if page >= totalPages}
            <div class="notification is-primary">
                <p class="has-text-centered">End of posts</p>
            </div>
        {/if}
    </div>
</section>
