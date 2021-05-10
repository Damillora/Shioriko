<script> 
    import { onMount } from "svelte";
    import { getPosts } from "../api.js";
    import { Link} from "svelte-routing";
    import queryString from "query-string";

    export let location;

    let page = 1;
    let totalPages = 1;
    let posts = [];
    const getData = async () => {
        const data = await getPosts({page});
        if(Array.isArray(data.posts)) {
            posts = data.posts;
        }
    }
    onMount(() => { 
      let queryParams;
      $: queryParams = queryString.parse(location.search);
      if(queryParams.page) {
          page = queryParams.page;
      }
        getData();
     })
    
    const handlePage = (i) => {
        return () => {
            page = 1;
            getData();
        }
    }
</script>

<section class="hero is-primary">
    <div class="hero-body">
      <p class="title">
        Posts
      </p>
    </div>
</section>

<section class="section">
    <div class="container">
        <nav class="pagination" role="navigation" aria-label="pagination">
            {#if page > 1}
            <a class="pagination-previous">Previous</a>
            {/if}
            {#if page < totalPages}
            <a class="pagination-next">Next page</a>
            {/if}
            <ul class="pagination-list">
            {#if page > 2}
            <li>
                <Link on:click="{handlePage(1)}" to="/posts?page={1}" class="pagination-link" aria-label="Goto page 1">1</Link>
            </li>
            <li>
                <span class="pagination-ellipsis">&hellip;</span>
            </li>
            {/if}
              {#each [...Array(5).keys()].map(x => x + page - 2) as i }
              {#if i >= 1 && i <= totalPages}
                {#if i == page}
                <li>
                <Link on:click="{handlePage(i)}" to="/posts?page={i}" class="pagination-link is-current" aria-label="Goto page {i}">{i}</Link>
            </li>
                {:else}
                <li>
                    <Link on:click="{handlePage(i)}" to="/posts?page={i}" class="pagination-link" aria-label="Goto page {i}">{i}</Link>
                </li>
                {/if}
              {/if}
              {/each}
              {#if (totalPages - page) > 2}
              <li>
                  <span class="pagination-ellipsis">&hellip;</span>
              </li>
              <li>
                <Link on:click="{handlePage(totalPages)}" to="/posts?page={totalPages}" class="pagination-link" aria-label="Goto page {totalPages}">{totalPages}</Link>
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
                        <img src="{post.image_path}">
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
