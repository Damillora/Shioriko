<script> 
    import { onMount } from "svelte";
    import { getPosts } from "../api.js";
    import { Link} from "svelte-routing";

    let page = 1;
    let totalPages = 1;
    let posts = [];
    const getData = async () => {
        const data = await getPosts({page});
        posts = data.posts;
    }
    onMount(() => { getData(); })
    
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
        <div class="tile is-ancestor">
            {#each posts as post (post.id)}
            <div class="tile is-child is-4 card">
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
                        <Link to="/tag/{tag}">{tag}</Link>
                        {/each}
                    </div>
                </div>
            </div>
            {/each}
        </div>
    </div>
</section>
