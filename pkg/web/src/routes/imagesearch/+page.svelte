<script>
    import { getTagAutocomplete, searchBlob } from "$lib/api";
    import PostGallery from "$lib/components/ui/PostGallery.svelte";
    import ShiorikoImage from "$lib/components/ui/ShiorikoImage.svelte";
    import { paginate } from "$lib/simple-pagination";

    let currentProgress = $state(0);

    let file = $state();
    let fileName = $state("");
    let similar = $state([]);
    let similarCount = $state(0);
    let loading = $state(false);
    let loaded = $state(false);

    let form = $state({
        blob_id: "",
        source_url: "",
        tags: [],
    });

    const onProgress = (e) => {
        var percentCompleted = Math.round((e.loaded * 100) / e.total);
        currentProgress = percentCompleted;
    };

    const onFileChange = async (e) => {
        file = e.target.files[0];
        fileName = file.name;
    };
    const onSubmit = async (e) => {
        e.preventDefault();
        loading = true;
        loaded = false;
        similar = [];
        if (file) {
            var response = await searchBlob({ file, onProgress });
            similar = response.similar;
            similarCount = similar.length;
        }
        loading = false;
        loaded = true;
    };
</script>

<section class="section">
    <div class="container">
        <div class="columns">
            <div class="column is-one-third">
                <div class="panel is-primary">
                    <form onsubmit={onSubmit}>
                        <p class="panel-heading">Image Search</p>
                        <div class="panel-block column">
                            <div class="row">
                                <label for="file" class="label">Image:</label>
                            </div>
                            <div class="row">
                                <div class="field">
                                    <div class="control">
                                        <div class="file">
                                            <label class="file-label">
                                                <input
                                                    id="file"
                                                    class="file-input"
                                                    type="file"
                                                    name="resume"
                                                    onchange={onFileChange}
                                                />
                                                <span class="file-cta">
                                                    <span class="file-icon"
                                                    ></span>
                                                    <span class="file-label">
                                                        Choose a file…
                                                    </span>
                                                </span>
                                            </label>
                                        </div>
                                    </div>
                                    {#if fileName}
                                        <p class="help">{fileName}</p>
                                    {/if}
                                </div>
                            </div>
                        </div>
                        {#if currentProgress > 0 && currentProgress < 100}
                            <div class="panel-block">
                                <progress
                                    class="progress is-primary is-small"
                                    value={currentProgress}
                                    max="100"
                                >
                                    {currentProgress}%
                                </progress>
                            </div>
                        {/if}
                        <div class="panel-block column">
                            <button
                                type="submit"
                                class="button is-primary is-fullwidth is-outlined"
                                >Search</button
                            >
                        </div>
                    </form>
                </div>
            </div>
            <div class="column is-two-thirds">
                {#if !loading}
                    {#if loaded}
                        {#if similarCount > 0}
                            <div class="block">
                                <div class="notification is-success">
                                    Found {similarCount} similar posts.
                                </div>
                            </div>
                        {:else}
                            <div class="block">
                                <div class="notification is-warning">
                                    Found no similar posts.
                                </div>
                            </div>
                        {/if}
                    {:else}
                        <div class="notification is-primary">
                            Similar posts will appear here.
                        </div>
                    {/if}
                {/if}
                <div class="columns is-multiline">
                    {#if !loading}
                        {#if similarCount > 0}
                            <div class="column is-full">
                                <PostGallery posts={similar} />
                            </div>
                        {/if}
                    {:else}
                        <div class="column">
                            <div class="skeleton-block"></div>
                        </div>
                    {/if}
                </div>
            </div>
        </div>
    </div>
</section>
