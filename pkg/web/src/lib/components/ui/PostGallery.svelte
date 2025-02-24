<script lang="ts">
    import { onMount } from "svelte";
    import ImageView from "./ImageView.svelte";

    let { posts = [] } = $props();
</script>

<div class="columns is-multiline">
    {#each posts as post, i (post.id)}
        <div class="column is-one-quarter">
            <div class="block">
                <div class="card">
                    <div class="card-image">
                        <figure class="image">
                            <a href="/post/{post.id}">
                                <ImageView alt={post.id} src={post.thumbnail_path} />
                            </a>
                        </figure>
                    </div>
                </div>
            </div>
        </div>
    {/each}
</div>
