<script>
    import { run } from 'svelte/legacy';

    import { getTags } from "$lib/api";
    import { afterNavigate } from '$app/navigation';
    import TagTypeIndicator from '$lib/components/ui/TagTypeIndicator.svelte';

    let tags = $state([]);

    const getData = async () => {
        const data = await getTags();
        tags = data;
    };
    
    afterNavigate(() => {
        getData();
    })
</script>


<section class="section">
    <div class="container">
        <h1 class="title">Tag List</h1>
        <table class="table is-fullwidth">
            <thead>
                <tr>
                    <th >Tag</th>
                    <th style="width: 30%;">Tag Type</th>
                    <th style="width: 10%;">Post Count</th>
                </tr>
            </thead>
            <tbody>
                {#each tags as tag}
                    <tr>
                        <td>
                          <a href="/tags/{tag.tagName}">{tag.tagName}</a>
                        </td>
                        <td><TagTypeIndicator tagType={tag.tagType} /></td>
                        <td>{tag.postCount}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</section>