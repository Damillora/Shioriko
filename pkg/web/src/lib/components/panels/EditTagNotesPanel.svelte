<script lang="ts">
    import { onMount } from "svelte";
    import { updateTagNotes } from "$lib/api";

    let { tag, data, toggleEditMenu, onSubmit } = $props();

    let form = $state({
        note: "",
    });
    let editNotesLoading = $state(false);

    const getData = async () => {
        form.note = data.tagNote;
    };

    const onFormSubmit = async (e) => {
        e.preventDefault();
        editNotesLoading = true;
        await updateTagNotes(tag, form);
        editNotesLoading = false;
        toggleEditMenu();

        onSubmit();
    };

    onMount(() => {
        getData();
    });
</script>

<form onsubmit={onFormSubmit}>
    <div class="panel is-warning">
        <p class="panel-heading">Edit Notes</p>
        {#if !editNotesLoading}
            <div class="panel-block column">
                <textarea bind:value={form.note} class="textarea has-fixed-size"
                ></textarea>
                <div class="content"></div>
            </div>
            <div class="panel-block column">
                <button type="submit" class="button is-primary">Save</button>
                <button onclick={toggleEditMenu} class="button">Cancel</button>
            </div>
        {:else}
            <div class="panel-block column">
                <progress class="progress is-small is-warning" max="100"
                ></progress>
            </div>
        {/if}
    </div>
</form>
