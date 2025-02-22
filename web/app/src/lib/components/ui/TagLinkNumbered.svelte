<script lang="ts">
    import TagTypeIndicator from "$lib/components/ui/TagTypeIndicator.svelte";

    let { tag, num } = $props();

    let tagType = tag.split(":")[0] ?? "";
    let tagName = tag.split(":")[1] ?? "";
    let tagDisplay = tagName.split("_").join(" ");
</script>

<a href="/posts?tags={tagName}">
    {tagDisplay}
    <TagTypeIndicator tagType={tagType}></TagTypeIndicator>
    <span class="is-pulled-right">{num}</span>
</a>
