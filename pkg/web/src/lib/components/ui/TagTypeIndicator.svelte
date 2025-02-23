<script>
    let { tagType } = $props();

</script>

{#if tagType == "character"}
<span class="tag is-link is-light">{tagType}</span>
{:else if tagType == "series"}
<span class="tag is-warning is-light">{tagType}</span>
{:else}
<span class="tag">{tagType}</span>
{/if}