<script>
    import { run } from "svelte/legacy";

    import { getTags } from "$lib/api";
    import { afterNavigate } from "$app/navigation";
    import TagTypeIndicator from "$lib/components/ui/TagTypeIndicator.svelte";

    let tags = $state([]);
    let loading = $state(false);
    let highestCount = $state(1);

    const getData = async () => {
        const data = await getTags();
        tags = data;
        loading = false;
        highestCount = Math.max(...data.map((x) => x.postCount));
        if (highestCount <= 0) {
            highestCount = 1;
        }
    };

    afterNavigate(() => {
        loading = true;
        getData();
    });
</script>

<section class="section">
    <div class="container">
        <h1 class="title">Tag List</h1>
        {#if !loading}
            <table class="table is-fullwidth">
                <thead>
                    <tr>
                        <th>Tag</th>
                        <th>Post Count</th>
                    </tr>
                </thead>
                <tbody>
                    {#each tags as tag}
                        <tr>
                            <td>
                                <a href="/tags/{tag.tagName}"
                                    >{tag.tagName}
                                    <TagTypeIndicator
                                        tagType={tag.tagType}
                                    /></a
                                >
                            </td>
                            <td>
                                <span class="is-pulled-right"
                                    >{tag.postCount}</span
                                ></td
                            >
                        </tr>
                    {/each}
                </tbody>
            </table>
        {:else}
            <div class="skeleton-block"></div>
        {/if}
    </div>
</section>
