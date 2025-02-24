<script>
    import { onMount } from "svelte";

    import { getTag, getPosts } from "$lib/api";
    import EditTagNotesPanel from "$lib/components/panels/EditTagNotesPanel.svelte";
    import ViewTagNotesPanel from "$lib/components/panels/ViewTagNotesPanel.svelte";
    import ViewTagPanel from "$lib/components/panels/ViewTagPanel.svelte";
    import EditTagPanel from "$lib/components/panels/EditTagPanel.svelte";
    import PostGallery from "$lib/components/ui/PostGallery.svelte";

    import { page } from "$app/stores";
    import AuthCheck from "$lib/components/checks/AuthCheck.svelte";
    let { tag } = $state($page.params);

    let data = $state();
    let posts = $state([]);

    const getData = async () => {
        if (tag) {
            data = null;
            data = await getTag({ tag });
            const response = await getPosts({
                page: 1,
                q: tag,
            });
            if (response.posts) {
                posts = response.posts.slice(0, 4);
            }
        }
    };

    let renameMenuShown = $state(false);
    const toggleRenameMenu = (e) => {
        e.preventDefault();
        renameMenuShown = !renameMenuShown;
    };

    let editMenuShown = $state(false);
    const toggleEditMenu = () => {
        editMenuShown = !editMenuShown;
    };

    const onTagSubmit = (newName) => {
        tag = newName;
        toggleEditMenu();
        getData();
    };

    onMount(() => {
        getData();
    });
</script>

<section class="section">
    <div class="container">
        <div class="columns">
            <div class="column is-one-third">
                {#if data}
                    <div class="block">
                        <ViewTagPanel {tag} {data} />
                    </div>
                    {#if renameMenuShown}
                    <div class="block">
                        <EditTagPanel
                            {tag}
                            {data}
                            {toggleRenameMenu}
                            onSubmit={onTagSubmit}
                        />
                    </div>
                    {:else}
                    <AuthCheck>
                        <div class="panel is-info">
                            <div class="panel-heading">Tag Actions</div>
                            <a class="panel-block" href="/posts?tags={tag}">Browse Posts</a>
                            <a onclick={toggleRenameMenu} class="panel-block"
                                >Rename</a
                            >
                        </div>
                    </AuthCheck>
                    {/if}
                {:else}
                    <div class="skeleton-block"></div>
                {/if}
            </div>
            <div class="column is-two-thirds">
                {#if data}
                    {#if editMenuShown}
                        <EditTagNotesPanel
                            {tag}
                            {data}
                            {toggleEditMenu}
                            onSubmit={getData}
                        />
                    {:else}
                        <ViewTagNotesPanel {data} {toggleEditMenu} />
                    {/if}
                    <h1 class="title">Posts</h1>
                    <PostGallery {posts} />
                {:else}
                    <div class="skeleton-block"></div>
                {/if}
            </div>
        </div>
    </div>
</section>
