<script>
    import { onMount } from "svelte";

    import { getTag, getPostSearchTag } from "$lib/api";
    import EditTagNotesPanel from "$lib/components/TagNotes/EditTagNotesPanel.svelte";
    import ViewTagNotesPanel from "$lib/components/TagNotes/ViewTagNotesPanel.svelte";
    import ViewTagPanel from "$lib/components/Tag/ViewTagPanel.svelte";
    import EditTagPanel from "$lib/components/Tag/EditTagPanel.svelte";
    import PostGallery from "$lib/components/Post/PostGallery.svelte";


    import { page } from "$app/stores";
    let { tag } = $page.params;

    let data;
    let posts = [];

    const getData = async () => {
        if (tag) {
            data = await getTag({ tag });
            const response = await getPostSearchTag({
                page: 1,
                q: tag,
            });
            if (response.posts) {
                posts = response.posts.slice(0, 4);
            }
        }
    };

    let renameMenuShown = false;
    const toggleRenameMenu = () => {
        renameMenuShown = !renameMenuShown;
    };

    let editMenuShown = false;
    const toggleEditMenu = () => {
        editMenuShown = !editMenuShown;
    };

    const onTagSubmit = (newName) => {
        tag = newName;
        getData();
    };

    onMount(() => {
        getData();
    });
</script>

<section class="section">
    <div class="container">
        {#if data}
            <div class="columns">
                <div class="column is-one-third">
                    {#if renameMenuShown}
                        <EditTagPanel
                            {tag}
                            {data}
                            {toggleRenameMenu}
                            onSubmit={onTagSubmit}
                        />
                    {:else}
                        <ViewTagPanel {tag} {data} {toggleRenameMenu} />
                    {/if}
                </div>
                <div class="column is-two-thirds">
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
                </div>
            </div>
        {/if}
    </div>
</section>
