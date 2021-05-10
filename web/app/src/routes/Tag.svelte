<script>
    import { onMount } from "svelte";
    import { getPostsTag } from "../api.js";
    import PostPaginator from "../PostPaginator.svelte";
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

<PostPaginator url="/tag/{id}" posts={posts} page={page} totalPages={totalPages} handlePage={handlePage} />