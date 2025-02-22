<script lang="ts">
    let { tag, num } = $props();

    let tagType = tag.split(":")[0] ?? "";
    let tagName = tag.split(":")[1] ?? "";
    let tagDisplay = tagName.split("_").join(" ");
</script>

<a href="/posts?tags={tagName}">
    {tagDisplay}
    {#if tagType == "character"}
    <span class="tag is-link is-light">character</span>
    {:else if tagType == "series"}
    <span class="tag is-warning is-light">series</span>
    {:else}
    <span class="tag">{tagType}</span>
    {/if}
    <span class="is-pulled-right">{num}</span>
</a>
