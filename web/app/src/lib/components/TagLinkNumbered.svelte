<script>
    export let tag;
    export let num;

    let tagType = tag.split(":")[0] ?? "";
    let tagName = tag.split(":")[1] ?? "";
    let tagDisplay = tagName.split("_").join(" ");
</script>

<a href="/posts?tags={tagName}"
    >{tagDisplay} <span class="is-pulled-right">{num}</span></a
>
