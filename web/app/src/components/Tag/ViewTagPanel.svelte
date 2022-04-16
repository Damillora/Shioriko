<script>
    import { onMount } from "svelte";

    import { Link } from "svelte-routing";
    import { getRelatedTags } from "../../api.js";
    import AuthCheck from "../../AuthCheck.svelte";
    import TagLinkNumbered from "../../TagLinkNumbered.svelte";

    export let tag;
    export let data;
    export let toggleRenameMenu;
    let related_tags = [];
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
        <div class="row">{data.tagType}</div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Posts:</strong>
        </div>
        <div class="row">
            {data.postCount} (<Link to="/posts?tags={tag}">Browse</Link>)
        </div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Related Tags:</strong>
        </div>
        <div class="row">
            <div class="menu">
                <ul class="menu-list">
                    {#each related_tags as tag (tag)}
                        <li>
                            <TagLinkNumbered
                                class=""
                                tag={tag.tagType + ":" + tag.tagName}
                                num={tag.postCount}
                            />
                        </li>
                    {/each}
                </ul>
            </div>
        </div>
    </div>
    <AuthCheck>
        <div class="panel-block column">
            <button
                on:click|preventDefault={toggleRenameMenu}
                class="button is-primary">Rename</button
            >
        </div>
    </AuthCheck>
</div>
