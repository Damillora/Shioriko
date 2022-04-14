<script>
    import { Link } from "svelte-routing";

    export let tag;

    let tagType = tag.split(":")[0] ?? "";
    let tagName = tag.split(":")[1] ?? "";
    let tagDisplay = tagName.split("_").join(" ");
</script>

<Link class="button is-rounded is-primary is-small m-1" to="/posts?tags={tagName}">{tagDisplay}</Link>
