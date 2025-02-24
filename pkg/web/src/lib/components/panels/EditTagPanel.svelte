<script lang="ts">
    import { onMount } from "svelte";

    import { getTagTypes, updateTag } from "$lib/api";
    import { goto } from "$app/navigation";

    let { tag, data, toggleRenameMenu, onSubmit } = $props();

    let tagTypes = $state([]);
    let form = $state({
        name: "",
        tagTypeId: 1,
    });
    let editTagLoading = $state(false);

    const getData = async () => {
        tagTypes = await getTagTypes();
        form.name = data.tagName;
        let tagType = tagTypes.filter((x) => x.name == data.tagType);
        form.tagTypeId = tagType[0].id;
    };

    const onFormSubmit = async (e) => {
        e.preventDefault();
        editTagLoading = true;
        await updateTag(tag, form);
        editTagLoading = false;
        onSubmit(form.name);
    };

    onMount(() => {
        getData();
    });
</script>

<form onsubmit={onFormSubmit}>
    <div class="panel is-warning">
        <p class="panel-heading">Edit Tag</p>
        {#if !editTagLoading}
            <div class="panel-block column">
                <div class="row">
                    <strong>Name:</strong>
                </div>
                <div class="row">
                    <div class="field">
                        <div class="control">
                            <input
                                class="input"
                                type="text"
                                bind:value={form.name}
                            />
                        </div>
                    </div>
                </div>
            </div>
            <div class="panel-block column">
                <div class="row">
                    <strong>Category:</strong>
                </div>
                <div class="row">
                    <div class="field">
                        <div class="select">
                            <select bind:value={form.tagTypeId}>
                                {#each tagTypes as tagType}
                                    <option
                                        value={tagType.id}
                                        selected={form.tagTypeId === tagType.id}
                                    >
                                        {tagType.name}
                                    </option>
                                {/each}
                            </select>
                        </div>
                    </div>
                </div>
            </div>
            <div class="panel-block column">
                <button class="button is-primary" type="submit">Submit</button>
                <button onclick={toggleRenameMenu} class="button">Cancel</button
                >
            </div>
        {:else}
            <div class="panel-block column">
                <progress class="progress is-small is-warning" max="100"
                ></progress>
            </div>
        {/if}
    </div>
</form>
