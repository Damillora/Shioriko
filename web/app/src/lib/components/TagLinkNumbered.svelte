<script lang="ts">
    let { tag, num } = $props();

    let tagType = tag.split(":")[0] ?? "";
    let tagName = tag.split(":")[1] ?? "";
    let tagDisplay = tagName.split("_").join(" ");
</script>

<a href="/posts?tags={tagName}"
    >{tagDisplay} <span class="is-pulled-right">{num}</span></a
>
