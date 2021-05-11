<script>
    import { onMount } from "svelte";
    import TagLink from "../TagLink.svelte";
    import { getPost, postCreate } from "../api.js";
    import { Link } from "svelte-routing";
    export let id;
    let post;
    const getData = async () => {
        const data = await getPost({ id });
        post = data;
    };

    const trimUrl = (str) => {
        if (str.length > 30) {
            return str.substring(0, 30) + "...";
        }
        return str;
    };

    onMount(() => {
        getData();
    });
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
                    <div class="content">
                        <p>
                            <Link
                                class="button is-primary"
                                to="/post/edit/{post.id}">Edit</Link
                            >
                        </p>
                        <p>
                            Uploader: {post.uploader}
                        </p>
                        <p>
                            Source URL: <a href={post.source_url}
                                >{trimUrl(post.source_url)}</a
                            >
                        </p>
                        <p>
                            Dimensions: {post.width}x{post.height}
                        </p>
                        <p>
                            Tags:<br />
                        </p>
                        <p>
                            {#if post.tags}
                                {#each post.tags as tag (tag)}
                                    <TagLink {tag} />
                                {/each}
                            {/if}
                        </p>
                    </div>
                </div>
                <div class="column">
                    <figure class="image">
                        <img alt={post.id} src={post.image_path} />
                    </figure>
                </div>
            </div>
        </section>
    </div>
{/if}
