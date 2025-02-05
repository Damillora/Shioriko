<script lang="ts">
    import { run } from 'svelte/legacy';

    import { onMount } from "svelte";
    import { getPost, postDelete } from "$lib/api";
    import { goto } from "$app/navigation";
    import EditPostPanel from "$lib/components/EditPostPanel.svelte";
    import ViewPostPanel from "$lib/components/ViewPostPanel.svelte";
    
    import { page } from "$app/stores";
    const { id } = $page.params;

    let post: any = $state();
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

    let deleteMenuShown = $state(false);

    const deletePost = async (e) => {
        e.preventDefault();
        toggleDeleteMenu();
        const success = await postDelete({ id });
        if (success) {
            goto("/posts");
        }
    };
    const toggleDeleteMenu = (e) => {
        e.preventDefault();
        deleteMenuShown = !deleteMenuShown;
    };

    let editMenuShown = $state(false);

    const toggleEditMenu = (e) => {
        e.preventDefault();
        editMenuShown = !editMenuShown;
    };

    let imagePercentage = $state("0%");

    run(() => {
        if (post)
            imagePercentage = ((1000 * 100) / post.width).toFixed(0) + "%";
    });
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
                                    onclick={deletePost}
                                    class="button is-danger">Delete</button
                                >
                                <button
                                    class="button"
                                    onclick={toggleDeleteMenu}
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
