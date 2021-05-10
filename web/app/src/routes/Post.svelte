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

    const trimUrl = (str) => {
        if(str.length > 30) {  
            return str.substring(0,30) + "...";
        }
        return str;
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
        <div class="columns">
            <div class="column is-one-third box">
                    <p>
                        Source URL: <a href="{post.source_url}">{trimUrl(post.source_url)}</a>
                    </p>
                    <p>
                        Tags: 
                        {#each post.tags as tag (tag)}
                        <ul>
                            <li>
                                <Link to="/tag/{tag}">{tag}</Link>
                            </li>
                        </ul>
                        {/each}
                    </p>
            </div>
            <div class="column">
                <figure class="image">
                    <img src="{post.image_path}">
                </figure>
            </div>
        </div>
    </section>
</div>
{/if}