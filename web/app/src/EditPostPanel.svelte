<script>
    import Tags from "svelte-tags-input";
    import { onMount } from "svelte";
    import { getPost, postUpdate, getTagAutocomplete } from "./api.js";

    export let isActive = false;
    export let post;
    export let onSubmit;

    const toggleEditModal = () => {
        isActive = !isActive;
    };

    let form = {
        source_url: "",
        tags: [],
    };

    const getData = async () => {
        form.source_url = post.source_url;
        form.tags = post.tags.map(x => x.tagType+":"+x.tagName);
    };

    const onTagChange = (value) => {
        form.tags = value.detail.tags;
    };

    const onAutocomplete = async () => {
        const list = await getTagAutocomplete();
        return list;
    };

    const onFormSubmit = async () => {
        const response = await postUpdate(post.id, form);
        toggleEditModal();

        onSubmit();
    };

    onMount(() => {
        getData();
    });
</script>

<form on:submit|preventDefault={onFormSubmit}>
    <div class="panel is-warning">
        <p class="panel-heading">Edit Post</p>
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
                <label for="source" class="label">Source URL:</label>
            </div>
            <div class="row">
                <div class="field">
                    <div class="control">
                        <input
                            id="source"
                            class="input"
                            type="url"
                            placeholder="Source URL"
                            bind:value={form.source_url}
                        />
                    </div>
                </div>
            </div>
        </div>
        <div class="panel-block column">
            <div class="row">
                <label for="tags" class="label">Tags:</label>
            </div>
            <div class="row">
                <div class="field">
                    <div class="control" id="tags">
                        <Tags
                            tags={form.tags}
                            addKeys={[9, 32]}
                            on:tags={onTagChange}
                            autoComplete={onAutocomplete}
                        />
                    </div>
                </div>
            </div>
        </div>
        <div class="panel-block column">
            <button class="button is-primary" type="submit">Save</button>
            <button class="button" on:click|preventDefault={toggleEditModal}
                >Cancel</button
            >
        </div>
    </div>
</form>
