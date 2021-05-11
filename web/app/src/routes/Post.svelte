<script>
    import { onMount } from "svelte";
    import TagLink from "../TagLink.svelte";
    import { getPost, postCreate, postDelete } from "../api.js";
    import { Link, navigate } from "svelte-routing";
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

    let modal_shown = false;

    const deletePost = async () => {
        toggleModal();
        const success = await postDelete({ id });
        if (success) {
            navigate("/posts");
        }
    };
    const toggleModal = () => {
        modal_shown = !modal_shown;
    };
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
                            <button
                                on:click|preventDefault={toggleModal}
                                class="button is-danger">Delete</button
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
                            Original: <a href={post.image_path}>Image</a>
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
                <div class="column box">
                    {#if post.width > 1000}
                        <div class="notification is-info">
                            The image has been resized due to size. The original
                            image link is in the sidebar
                        </div>
                        <figure class="image">
                            <img alt={post.id} src={post.preview_path} />
                        </figure>
                    {:else}
                        <figure class="image">
                            <img alt={post.id} src={post.image_path} />
                        </figure>
                    {/if}
                </div>
            </div>
        </section>
    </div>

    <div class="modal" class:is-active={modal_shown}>
        <div class="modal-background" />
        <div class="modal-content">
            Are you sure to delete post {post.id}?
        </div>
        <div class="modal-card">
            <header class="modal-card-head">
                <p class="modal-card-title">Delete post?</p>
                <button class="delete" aria-label="close" />
            </header>
            <section class="modal-card-body" />
            <footer class="modal-card-foot">
                <button
                    on:click|preventDefault={deletePost}
                    class="button is-danger">Delete</button
                >
                <button class="button" on:click|preventDefault={toggleModal}
                    >Cancel</button
                >
            </footer>
        </div>
    </div>
{/if}
