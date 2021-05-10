<script>
    import { uploadBlob, postCreate } from "../api.js";
    import { navigate } from "svelte-routing";
    import Tags from "svelte-tags-input";

    let currentProgress = 0;

    let fileName = "";

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
        var file = e.target.files[0];
        if (file) {
            var response = await uploadBlob({ file, onProgress });
            form.blob_id = response.id;
            fileName = file.name;
        }
    };

    const onTagChange = (value) => {
        form.tags = value.detail.tags;
    }

    const onSubmit = async () => {
        const response = await postCreate(form);
        navigate(`/post/${response.id}`);
    };
</script>

<section class="hero is-primary">
    <div class="hero-body">
        <p class="title">Upload</p>
    </div>
</section>

<section class="section">
    <div class="container">
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
                    <Tags on:tags={onTagChange} />
                </div>
            </div>
            <div class="control">
                <button type="submit" class="button is-primary">Submit</button>
            </div>
        </form>
    </div>
</section>
