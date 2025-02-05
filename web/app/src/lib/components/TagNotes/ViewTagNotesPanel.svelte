<script lang="ts">
    import AuthCheck from "$lib/components/AuthCheck.svelte";

    let { data, toggleEditMenu } = $props();
</script>

<div class="panel is-info">
    <p class="panel-heading">Notes</p>
    <div class="panel-block column">
        <div class="content pre-line">
            {data.tagNote}
        </div>
    </div>
    <AuthCheck>
        <div class="panel-block column">
            <button
                onclick={toggleEditMenu}
                class="button is-primary">Edit</button
            >
        </div>
    </AuthCheck>
</div>
