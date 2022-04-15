<script>
    import { onMount } from "svelte";

    import { getTag } from "../../api";
    import EditTagNotesPanel from "../../components/TagNotes/EditTagNotesPanel.svelte";
    import ViewTagNotesPanel from "../../components/TagNotes/ViewTagNotesPanel.svelte";
    import ViewTagPanel from "../../components/Tag/ViewTagPanel.svelte";
    import EditTagPanel from "../../components/Tag/EditTagPanel.svelte";

    export let tag;
    let data;

    const getData = async () => {
        if (tag) {
            data = await getTag({ tag });
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
                </div>
            </div>
        {/if}
    </div>
</section>
