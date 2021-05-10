<script>
    import { onMount } from "svelte";
    import { getPostSearchTag } from "../api.js";
    import { navigate } from "svelte-routing";
    import PostPaginator from "../PostPaginator.svelte";
    import queryString from "query-string";
    import Tags from "svelte-tags-input";

    export let location;

    let searchTerms = [];

    let page = 1;
    let totalPages = 1;
    let posts = [];
    const getData = async () => {
        const data = await getPostSearchTag({ page, q: searchTerms.join("+") });
        if (Array.isArray(data.posts)) {
            posts = data.posts;
            totalPages = data.totalPage;
        }
    };

    let queryParams;

    const onTagChange = (value) => {
        searchTerms = value.detail.tags;
    };

    $: {
        queryParams = queryString.parse(location.search);
        console.log(queryParams);
        if (queryParams.page) {
            page = parseInt(queryParams.page);
        }
        if (queryParams.tags) {
            searchTerms = queryParams.tags.split(" ");
        } else {
            searchTerms = [];
        }
        getData();
    }

    const handlePage = (i) => {
        return () => {
            page = i;
            getData();
        };
    };

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
                                addKeys={[9,32]}
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
            <PostPaginator
                url="/posts?tags={searchTerms.join('+')}&"
                {posts}
                {page}
                {totalPages}
                {handlePage}
            />
        </div>
    </div>
</section>
