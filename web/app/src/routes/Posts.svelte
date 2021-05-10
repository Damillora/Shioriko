<script>
    import { onMount } from "svelte";
    import { getPosts } from "../api.js";
    import { Link } from "svelte-routing";
    import queryString from "query-string";
    import PostPaginator from "../PostPaginator.svelte";

    export let location;

    let page = 1;
    let totalPages = 1;
    let posts = [];
    const getData = async () => {
        const data = await getPosts({ page });
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
        <p class="title">Posts</p>
    </div>
</section>

<PostPaginator url="/posts" posts={posts} page={page} totalPages={totalPages} handlePage={handlePage} />