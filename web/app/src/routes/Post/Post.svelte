<script>
    import { onMount } from "svelte";
    import { getPost, postDelete } from "../../api.js";
    import { navigate } from "svelte-routing";
    import EditPostPanel from "../../EditPostPanel.svelte";
    import ViewPostPanel from "../../ViewPostPanel.svelte";
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

    let deleteMenuShown = false;

    const deletePost = async () => {
        toggleDeleteMenu();
        const success = await postDelete({ id });
        if (success) {
            navigate("/posts");
        }
    };
    const toggleDeleteMenu = () => {
        deleteMenuShown = !deleteMenuShown;
    };

    let editMenuShown = false;

    const toggleEditMenu = () => {
        editMenuShown = !editMenuShown;
    };

    let imagePercentage = "0%";

    $: {
        if (post)
            imagePercentage = ((1000 * 100) / post.width).toFixed(0) + "%";
    }
</script>

{#if post}
    <div class="container">
        <section class="section">
            <div class="columns">
                <div class="column is-one-third">
                    {#if editMenuShown == false && deleteMenuShown == false}
                        <ViewPostPanel
                            {post}
                            {toggleDeleteMenu}
                            {toggleEditMenu}
                        />
                    {:else if editMenuShown == true}
                        <EditPostPanel
                            bind:isActive={editMenuShown}
                            {post}
                            onSubmit={getData}
                        />
                    {:else if deleteMenuShown == true}
                        <div class="panel is-danger">
                            <p class="panel-heading">Delete Post</p>
                            <div class="panel-block">
                                Are you sure to delete post {post.id}?
                            </div>
                            <div class="panel-block column">
                                <button
                                    on:click|preventDefault={deletePost}
                                    class="button is-danger">Delete</button
                                >
                                <button
                                    class="button"
                                    on:click|preventDefault={toggleDeleteMenu}
                                    >Cancel</button
                                >
                            </div>
                        </div>
                    {/if}
                </div>
                <div class="column box">
                    {#if post.width > 1000}
                        <div class="notification is-info">
                            Resized to {imagePercentage} of the original image.
                            <a href={post.image_path} target="_blank"
                                >View original</a
                            >
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
{/if}
