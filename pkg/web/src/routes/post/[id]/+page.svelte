<script lang="ts">
    import { onMount } from "svelte";
    import { getPost, postDelete } from "$lib/api";
    import { afterNavigate, goto } from "$app/navigation";
    import EditPostPanel from "$lib/components/panels/EditPostPanel.svelte";
    import ViewPostPanel from "$lib/components/panels/ViewPostPanel.svelte";

    import { page } from "$app/stores";
    import ImageView from "$lib/components/ui/ImageView.svelte";
    import AuthCheck from "$lib/components/checks/AuthCheck.svelte";
    import DeletePostPanel from "$lib/components/panels/DeletePostPanel.svelte";
    const { id } = $page.params;

    let post: any = $state();
    const getData = async () => {
        post = null;
        const data = await getPost({ id });
        post = data;
        imagePercentage = ((1000 * 100) / post.width).toFixed(0) + "%";
    };
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
    };
    onMount(() => {
        getData();
    });

    const onDelete = async (success) => {
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
                    <div class="block">
                        <ViewPostPanel {post} />
                    </div>
                    {#if editMenuShown == true && deleteMenuShown == false}
                        <div class="block">
                            <EditPostPanel
                                bind:isActive={editMenuShown}
                                {post}
                                onSubmit={onSubmitEdit}
                            />
                        </div>
                    {:else if deleteMenuShown == true}
                        <DeletePostPanel
                            id={post.id}
                            {toggleDeleteMenu}
                            {onDelete}
                        />
                    {:else}
                        <AuthCheck>
                            <div class="panel is-info">
                                <div class="panel-heading">Post Actions</div>

                                <a
                                    class="panel-block"
                                    href={post.image_path}
                                    target="_blank">View Original</a
                                >
                                <a
                                    href={"#"}
                                    onclick={toggleEditMenu}
                                    class="panel-block">Edit</a
                                >
                                <a
                                    href={"#"}
                                    onclick={toggleDeleteMenu}
                                    class="panel-block">Delete</a
                                >
                            </div>
                        </AuthCheck>
                    {/if}
                {:else}
                    <div class="skeleton-block"></div>
                {/if}
            </div>
            <div class="column is-two-thirds">
                {#if post}
                    <div class="block">
                        {#if post.width > 1000 && isOriginal == false}
                            <div class="notification is-info">
                                Resized to {imagePercentage} of the original image.
                                <a
                                    href={"#"}
                                    onclick={() => {
                                        isOriginal = true;
                                    }}>View original</a
                                >
                            </div>
                            <div class="box">
                                <figure class="image">
                                    <ImageView
                                        alt={post.id}
                                        src={post.preview_path}
                                    />
                                </figure>
                            </div>
                        {:else}
                            <div class="notification is-primary">
                                Currently viewing original image.
                            </div>
                            <div class="box">
                                <figure class="image">
                                    <ImageView
                                        alt={post.id}
                                        src={post.image_path}
                                    />
                                </figure>
                            </div>
                        {/if}
                    </div>
                {:else}
                    <div class="skeleton-block"></div>
                {/if}
            </div>
        </div>
    </div>
</section>
