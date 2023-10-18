<script lang="ts">
    import { getPostSearchTag, getTag, getTagAutocomplete } from "$lib/api";
    import TagLinkNumbered from "$lib/components/TagLinkNumbered.svelte";
    import PostGallery from "$lib/components/Post/PostGallery.svelte";
    import queryString from "query-string";
    import Tags from "svelte-tags-input";
    import { paginate } from "$lib/simple-pagination";
    import { beforeNavigate, goto } from "$app/navigation";
    import { page as currentPage } from '$app/stores';

    $: url = $currentPage.url;

    let searchTerms = [];

    let page = 1;
    let totalPages = 1;
    let pagination = [];
    let posts = [];
    let postCount = 0;
    let tags = [];
    let tagInfo = null;
    let categorizedTags = {};

    const getData = async () => {
        const data = await getPostSearchTag({ page, q: searchTerms.join("+") });
        if (data.posts) {
            posts = data.posts;
            tags = data.tags
                .filter(
                    (x) =>
                        !searchTerms.includes(x.tagName) &&
                        !searchTerms.includes(x.tagType + ":" + x.tagName)
                )
                .sort((a, b) => b.postCount - a.postCount);
            totalPages = data.totalPage;
            postCount = data.postCount;
            pagination = paginate(page, totalPages);
        } else {
            posts = [];
            tags = [];
            totalPages = 0;
            postCount = 0;
            pagination = paginate(page, totalPages);
        }

        if (searchTerms.filter(x => !x.startsWith("-")).length == 1) {
            tagInfo = await getTag({ tag: searchTerms[0] });
        }
    };
    let tagQuery;

    const onTagChange = (value) => {
        searchTerms = value.detail.tags;
    };

    const onAutocomplete = async (tag) => {
        const list = await getTagAutocomplete({ tag });
        return list;
    };

    $: {
        tagQuery = url.searchParams.get('tags');
        if (tagQuery) {
            searchTerms = tagQuery.split(" ");
        } else {
            searchTerms = [];
            tagInfo = null;
        }
        posts = [];
        page = 1;
        getData();
    }
    const onSearch = (i) => {
        if (searchTerms.length > 0) {
            goto(`/posts?tags=${searchTerms.join("+")}`);
        } else {
            goto(`/posts`);
        }
    };

    const changePage = (i) => {
        if (i >= 1 && i <= totalPages) {
            page = i;
            getData();
        }
    };
</script>

<section class="section">
    <div class="container">
        <div class="block">
            <div class="columns is-multiline">
                <div class="column is-full">
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
                                            autoCompleteFilter={false}
                                        />
                                    </div>
                                </div>
                                <div class="control">
                                    <button
                                        type="submit"
                                        class="button is-primary"
                                    >
                                        Search
                                    </button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
                <div class="column is-one-third">
                    {#if tagInfo}
                        <div class="panel is-info">
                            <p class="panel-heading">
                                Tag:
                                {tagInfo.tagName.split("_").join(" ")}
                            </p>
                            {#if tagInfo.tagNote}
                                <div class="panel-block column">
                                    <div class="content pre-line">
                                        {tagInfo.tagNote}
                                    </div>
                                </div>
                            {/if}
                            <div class="panel-block column">
                                <a
                                    class="button is-primary"
                                    href="/tags/{tagInfo.tagName}">View Tag</a
                                >
                            </div>
                        </div>
                    {/if}
                    <div class="panel is-primary">
                        <div class="panel-heading">Tags</div>
                        <div class="panel-block column">
                            <div class="menu">
                                <ul class="menu-list">
                                    {#each tags as tag (tag)}
                                        <li>
                                            <TagLinkNumbered
                                                class=""
                                                tag={tag.tagType +
                                                    ":" +
                                                    tag.tagName}
                                                num={tag.postCount}
                                            />
                                        </li>
                                    {/each}
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="column is-two-thirds">
                    <div class="columns is-multiline">
                        <div class="column is-full">
                            <PostGallery {posts} />
                        </div>

                        <div class="column is-full">
                            <nav
                                class="pagination is-centered"
                                aria-label="pagination"
                            >
                                <a
                                    href={null}
                                    on:click={changePage(page - 1)}
                                    class="pagination-previous"
                                    class:is-disabled={page == 1}>Previous</a
                                >
                                <a
                                    href={null}
                                    on:click={changePage(page + 1)}
                                    class="pagination-next"
                                    class:is-disabled={page == totalPages}
                                    >Next</a
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
                                                        changePage(pageEntry)}
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
                    </div>
                </div>
            </div>
        </div>
    </div>
</section>
