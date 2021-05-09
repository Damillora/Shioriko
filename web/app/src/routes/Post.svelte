<script>
    import { onMount } from "svelte";
    import { Link } from "svelte-routing";
    import {getPost } from "../api.js";
    export let id;
    let post;
    const getData = async() => {
        const data = await getPost({id});
        post = data;
    }

    onMount(() => {getData()});
</script>


<section class="hero is-primary">
    <div class="hero-body">
        {#if post}
      <p class="title">
        Post ID: {post.id}
      </p>
      {/if}
    </div>
</section>
{#if post}
<div class="container">
    <section class="section">
        <div class="tile is-ancestor">
            <div class="tile is-child is-4 box">
                    <p>
                        Source URL: <a href="//{post.source_url}">{post.source_url}</a>
                    </p>
                    <p>
                        Tags: 
                        {#each post.tags as tag (tag)}
                        <Link to="/tag/{tag}">{tag}</Link>
                        {/each}
                    </p>
            </div>
            <div class="tile is-child">
                <figure class="image">
                    <img src="{post.image_path}">
                </figure>
            </div>
        </div>
    </section>
</div>
{/if}