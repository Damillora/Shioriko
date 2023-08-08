<script>
    import { onMount } from "svelte";
    import { updateTagNotes } from "$lib/api";

    export let tag;
    export let data;
    export let toggleEditMenu;
    export let onSubmit;

    let form = {
        note: "",
    };

    const getData = async () => {
        form.note = data.tagNote;
    };

    const onFormSubmit = async () => {
        await updateTagNotes(tag, form);
        toggleEditMenu();

        onSubmit();
    };

    onMount(() => {
        getData();
    });
</script>

<form on:submit|preventDefault={onFormSubmit}>
    <div class="panel is-warning">
        <p class="panel-heading">Edit Notes</p>
        <div class="panel-block column">
            <textarea
                bind:value={form.note}
                class="textarea has-fixed-size"
            />
            <div class="content" />
        </div>
        <div class="panel-block column">
            <button type="submit" class="button is-primary">Save</button>
            <button on:click|preventDefault={toggleEditMenu} class="button"
                >Cancel</button
            >
        </div>
    </div>
</form>
