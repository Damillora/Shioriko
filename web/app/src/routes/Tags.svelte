<script>
    import { getTags } from "../api";
    import { Link } from "svelte-routing";

    let tags = [];

    const getData = async () => {
        const data = await getTags();
        tags = data;
    };
    $: {
        getData();
    }
</script>


<section class="hero is-primary">
    <div class="hero-body">
        <p class="title">Tags</p>
    </div>
</section>

<section class="section">
    <div class="container">
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
                          <Link to="/posts?tags={tag.tagType}:{tag.tagName}">{tag.tagName}</Link>
                        </td>
                        <td>{tag.tagType}</td>
                        <td>{tag.postCount}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</section>