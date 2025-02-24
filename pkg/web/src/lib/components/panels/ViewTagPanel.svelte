<script lang="ts">
    import { onMount } from "svelte";

    import { getRelatedTags } from "$lib/api";
    import AuthCheck from "$lib/components/checks/AuthCheck.svelte";
    import TagLinkNumbered from "$lib/components/ui/TagLinkNumbered.svelte";
    import TagTypeIndicator from "../ui/TagTypeIndicator.svelte";

    let { tag, data, toggleRenameMenu } = $props();
    let related_tags = $state([]);

    let tabPage = $state(1);

    const changeTab = (tab) => {
        tabPage = tab;
    };
    const getData = async () => {
        related_tags = await getRelatedTags({ tag });
        related_tags = related_tags
            .filter((x) => x.tagName != tag)
            .sort((a, b) => b.postCount - a.postCount);
    };
    onMount(() => {
        getData();
    });
</script>

<div class="panel is-primary">
    <p class="panel-heading">Tag</p>
    <div class="panel-tabs">
        <a
            href={"#"}
            class:is-active={tabPage == 1}
            onclick={() => changeTab(1)}>Information</a
        >
        <a
            href={"#"}
            class:is-active={tabPage == 2}
            onclick={() => changeTab(2)}>Related Tags</a
        >
    </div>
    {#if tabPage === 1}
        <div class="panel-block column">
            <div class="row">
                <strong>Name:</strong>
            </div>
            <div class="row">{data.tagName}</div>
        </div>
        <div class="panel-block column">
            <div class="row">
                <strong>Category:</strong>
            </div>
            <div class="row"><TagTypeIndicator tagType={data.tagType} /></div>
        </div>
        <div class="panel-block column">
            <div class="row">
                <strong>Posts:</strong>
            </div>
            <div class="row">
                {data.postCount}
            </div>
        </div>
    {:else if tabPage === 2}
        {#each related_tags as tag (tag)}
            <TagLinkNumbered
                tag={tag.tagType + ":" + tag.tagName}
                num={tag.postCount}
            />
        {/each}
    {/if}
</div>
