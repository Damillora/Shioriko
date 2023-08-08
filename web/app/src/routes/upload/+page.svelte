<script>
    import { uploadBlob, postCreate, getTagAutocomplete } from "$lib/api";
    import { goto } from "$app/navigation";
    import Tags from "svelte-tags-input";
    import AuthRequired from "$lib/components/AuthRequired.svelte";

    let currentProgress = 0;

    let fileName = "";
    let similar = [];

    let form = {
        blob_id: "",
        source_url: "",
        tags: [],
    };

    const onProgress = (e) => {
        var percentCompleted = Math.round((e.loaded * 100) / e.total);
        currentProgress = percentCompleted;
    };

    const onFileChange = async (e) => {
        fileName = "";
        similar = [];
        var file = e.target.files[0];
        if (file) {
            var response = await uploadBlob({ file, onProgress });
            if (response.similar) {
                similar = response.similar;
            }
            form.blob_id = response.id;
            fileName = file.name;
        }
    };

    const onTagChange = (value) => {
        form.tags = value.detail.tags;
    };

    const onAutocomplete = async (tag) => {
        const list = await getTagAutocomplete({ tag, positive: true });
        return list;
    };

    const onSubmit = async () => {
        const response = await postCreate(form);
        goto(`/post/${response.id}`);
    };
</script>

<AuthRequired />

<section class="section">
    <div class="container">
        <h1 class="title">Upload Image</h1>
        <form on:submit|preventDefault={onSubmit}>
            <div class="field">
                <label for="file" class="label">Image File</label>
                <div class="control">
                    <div class="file">
                        <label class="file-label">
                            <input
                                id="file"
                                class="file-input"
                                type="file"
                                name="resume"
                                on:change={onFileChange}
                            />
                            <span class="file-cta">
                                <span class="file-icon" />
                                <span class="file-label"> Choose a file… </span>
                            </span>
                        </label>
                    </div>
                </div>
                {#if currentProgress > 0 && currentProgress < 100}
                    <p class="help">{currentProgress}%</p>
                {/if}
                {#if fileName !== ""}
                    <p class="help">{fileName} uploaded</p>
                {/if}
                {#if similar.length > 0}
                    <p class="help">
                        Similar posts:
                        {#each similar as post, i}
                            <a href="/post/{post.id}">{post.id}</a>
                            {#if i < similar.length - 1}
                                ,
                            {/if}
                        {/each}
                    </p>
                {/if}
            </div>
            <div class="field">
                <label for="source" class="label">Source URL</label>
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
            <div class="field">
                <label for="tags" class="label">Tags</label>
                <div class="control" id="tags">
                    <Tags
                        addKeys={[9, 32]}
                        on:tags={onTagChange}
                        autoComplete={onAutocomplete}
                        autoCompleteFilter={false}
                    />
                </div>
            </div>
            <div class="control">
                <button type="submit" class="button is-primary">Submit</button>
            </div>
        </form>
    </div>
</section>
