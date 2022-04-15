<script>
    import { Link } from "svelte-routing";
    import AuthCheck from "../../AuthCheck.svelte";

    export let tag;
    export let data;
    export let toggleRenameMenu;
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
    <AuthCheck>
        <div class="panel-block column">
            <button
                on:click|preventDefault={toggleRenameMenu}
                class="button is-primary">Rename</button
            >
        </div>
    </AuthCheck>
</div>
