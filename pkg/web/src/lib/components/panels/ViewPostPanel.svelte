<script lang="ts">
    import AuthCheck from "$lib/components/checks/AuthCheck.svelte";
    import TagLinkNumbered from "$lib/components/ui/TagLinkNumbered.svelte";
    import { format, formatDistanceToNow } from "date-fns";
    let { post } = $props();

    let tabPage  = $state(1);

    const changeTab = (tab) => {
        tabPage = tab;
    }
    const trimUrl = (str) => {
        if (str.length > 30) {
            return str.substring(0, 30) + "...";
        }
        return str;
    };
</script>

<div class="panel is-primary">
    <p class="panel-heading">Post</p>
    <div class="panel-tabs">
      <a href={"#"} class:is-active="{tabPage == 1}" onclick={() => changeTab(1)}>Tags</a>
      <a href={"#"} class:is-active="{tabPage == 2}" onclick={() => changeTab(2)}>Information</a>
    </div>
    {#if tabPage === 1}
    
    {#if post.tags}
    {#each post.tags as tag (tag)}
        <TagLinkNumbered
            tag={tag.tagType + ":" + tag.tagName}
            num={tag.postCount}
        />
        {/each}
    {/if}
    {:else if tabPage == 2}
    <div class="panel-block column">
        <div class="row">
            <strong>Upload Date:</strong>
        </div>
        <div class="row"><time title={format(post.upload_date, "dd MMMM yyyy HH:mm:ss")} datetime={post.upload_date} >{formatDistanceToNow(post.upload_date, {addSuffix: true })}</time></div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Uploader:</strong>
        </div>
        <div class="row">{post.uploader}</div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Original:</strong>
        </div>
        <div class="row">
            <a href={post.image_path} target="_blank">Image</a>
        </div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Dimensions:</strong>
        </div>
        <div class="row">
            {post.width}x{post.height}
        </div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Source URL:</strong>
        </div>
        <div class="row">
            <a href={post.source_url}>{trimUrl(post.source_url)}</a>
        </div>
    </div>
    {/if}
</div>
