<script>
    import AuthCheck from "./AuthCheck.svelte";
    import TagLinkNumbered from "./TagLinkNumbered.svelte";
    export let post;
    export let toggleEditMenu;
    export let toggleDeleteMenu;

    const trimUrl = (str) => {
        if (str.length > 30) {
            return str.substring(0, 30) + "...";
        }
        return str;
    };
</script>

<div class="panel is-primary">
    <p class="panel-heading">Post</p>
    <div class="panel-block column">
        <div class="row">
            <strong>Uploader:</strong>
        </div>
        <div class="row">{post.uploader}</div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Source URL:</strong>
        </div>
        <div class="row">
            <a href={post.source_url}>{trimUrl(post.source_url)}</a>
        </div>
    </div>
    <div class="panel-block column">
        <div class="row">
            <strong>Original:</strong>
        </div>
        <div class="row">
            <a href={post.image_path}>Image</a>
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
            <p><strong>Tags:</strong></p>
        </div>
        <div class="row">
            <div class="menu">
                <ul class="menu-list">
                    {#if post.tags}
                        {#each post.tags as tag (tag)}
                            <li>
                                <TagLinkNumbered
                                    class=""
                                    tag={tag.tagType + ":" + tag.tagName}
                                    num={tag.postCount}
                                />
                            </li>
                        {/each}
                    {/if}
                </ul>
            </div>
        </div>
    </div>
    <AuthCheck>
        <p class="panel-block column">
            <button
                on:click|preventDefault={toggleEditMenu}
                class="button is-primary">Edit</button
            >
            <button
                on:click|preventDefault={toggleDeleteMenu}
                class="button is-danger">Delete</button
            >
        </p>
    </AuthCheck>
</div>
