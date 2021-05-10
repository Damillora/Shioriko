<script>
import { getTags } from "../api";


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
                    <th>Tag</th>
                    <th>Tag Type</th>
                </tr>
            </thead>
            <tbody>
                {#each tags as tag}
                    <tr>
                        <td>{tag.name}</td>
                        <td>{tag.tagType}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</section>