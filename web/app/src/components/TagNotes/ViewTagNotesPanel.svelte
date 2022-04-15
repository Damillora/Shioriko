<script>
    import { Link } from "svelte-routing";

    export let data;
    export let toggleEditMenu;

</script>


<div class="panel is-info">
    <p class="panel-heading">Notes</p>
    <div class="panel-block column">
        <div class="content pre-line">
            {data.tagNote}
        </div>
    </div>
    <div class="panel-block column">
        <button on:click|preventDefault={toggleEditMenu} class="button is-primary">Edit</button>
    </div>
</div>
<style>
    .pre-line {
        white-space: pre-line;
    }
</style>