<script>
    import { getPost, postUpdate } from "../api.js";
    import { navigate } from "svelte-routing";
    import Tags from "svelte-tags-input";
    import { onMount } from "svelte";
    import { Link } from "svelte-routing";
    import AuthRequired from "../AuthRequired.svelte";

    export let id;

    let image_path = "";

    let form = {
        source_url: "",
        tags: [],
    };

    const getData = async () => {
        const data = await getPost({ id });
        form.source_url = data.source_url;
        form.tags = data.tags;
        image_path = data.image_path;
    };

    const onTagChange = (value) => {
        form.tags = value.detail.tags;
    };

    const onSubmit = async () => {
        const response = await postUpdate(id, form);
        navigate(`/post/${response.id}`);
    };

    onMount(() => {
        getData();
    });
</script>

<AuthRequired />

<section class="hero is-primary">
    <div class="hero-body">
        <p class="title">Edit Post: {id}</p>
    </div>
</section>

<div class="container">
    <section class="section">
        <div class="columns">
            <div class="column is-one-third box">
                <p>
                    <Link class="button is-primary" to="/post/{id}">Back</Link>
                </p>
                <form on:submit|preventDefault={onSubmit}>
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
                                tags={form.tags}
                                addKeys={[9, 32]}
                                on:tags={onTagChange}
                            />
                        </div>
                    </div>
                    <div class="control">
                        <button type="submit" class="button is-primary"
                            >Submit</button
                        >
                    </div>
                </form>
            </div>
            <div class="column">
                <figure class="image">
                    <img alt={id} src={image_path} />
                </figure>
            </div>
        </div>
    </section>
</div>
