<script lang="ts">
    import { onMount } from "svelte";
    import { updateTagNotes } from "$lib/api";

    let {
        tag,
        data,
        toggleEditMenu,
        onSubmit
    } = $props();

    let form = $state({
        note: "",
    });

    const getData = async () => {
        form.note = data.tagNote;
    };

    const onFormSubmit = async (e) => {
        e.preventDefault();
        await updateTagNotes(tag, form);
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
        <div class="panel-block column">
            <textarea
                bind:value={form.note}
                class="textarea has-fixed-size"
></textarea>
            <div class="content"></div>
        </div>
        <div class="panel-block column">
            <button type="submit" class="button is-primary">Save</button>
            <button onclick={toggleEditMenu} class="button"
                >Cancel</button
            >
        </div>
    </div>
</form>
