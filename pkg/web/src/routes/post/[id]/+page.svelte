<script lang="ts">
    import { onMount } from "svelte";
    import { getPost, postDelete } from "$lib/api";
    import { afterNavigate, goto } from "$app/navigation";
    import EditPostPanel from "$lib/components/panels/EditPostPanel.svelte";
    import ViewPostPanel from "$lib/components/panels/ViewPostPanel.svelte";

    import { page } from "$app/stores";
    import ShiorikoImage from "$lib/components/ui/ShiorikoImage.svelte";
    const { id } = $page.params;

    let post: any = $state();
    const getData = async () => {
        const data = await getPost({ id });
        post = data;
        imagePercentage = ((1000 * 100) / post.width).toFixed(0) + "%";
    };
    let loading = $state(false);
    let isOriginal = $state(false);

    const trimUrl = (str: string) => {
        if (str.length > 30) {
            return str.substring(0, 30) + "...";
        }
        return str;
    };


    let deleteMenuShown = $state(false);

    const onSubmitEdit = () => {
        getData();
        editMenuShown = false;
    }
    onMount(() => {
        getData();
    });

    const deletePost = async (e) => {
        e.preventDefault();
        toggleDeleteMenu(e);
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
</script>

<section class="section">
    <div class="container">
        <div class="columns">
            <div class="column is-one-third">
                {#if post}
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
                            onSubmit={onSubmitEdit}
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
                                    onclick={toggleDeleteMenu}>Cancel</button
                                >
                            </div>
                        </div>
                    {/if}
                {:else}
                    <div class="skeleton-block"></div>
                {/if}
            </div>
            <div class="column box">
                {#if post}
                    {#if post.width > 1000 && isOriginal == false}
                        <div class="notification is-info">
                            Resized to {imagePercentage} of the original image.
                            <a
                                onclick={() => {
                                    isOriginal = true;
                                }}>View original</a
                            >
                        </div>
                        <figure class="image">
                            <ShiorikoImage
                                alt={post.id}
                                src={post.preview_path}
                            />
                        </figure>
                    {:else}
                        <div class="notification is-primary">
                            Currently viewing original image.
                        </div>
                        <figure class="image">
                            <ShiorikoImage
                                alt={post.id}
                                src={post.image_path}
                            />
                        </figure>
                    {/if}
                {:else}
                    <div class="skeleton-block"></div>
                {/if}
            </div>
        </div>
    </div>
</section>
