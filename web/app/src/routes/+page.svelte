<script lang="ts">
  import Tags from "svelte-tags-input";
  import { getPostCount, getPosts, getTagAutocomplete } from "$lib/api";

  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import PostGallery from "$lib/components/ui/PostGallery.svelte";

  let searchTerms: string[] = $state([]);
  let postCount: number = $state(0);

  const onTagChange = (value) => {
    searchTerms = value.detail.tags;
  };

  const onAutocomplete = async (tag) => {
    const list = await getTagAutocomplete({ tag });
    return list;
  };

  const onSearch = (e) => {
    e.preventDefault();
    if (searchTerms.length > 0) {
      goto(`/posts?tags=${searchTerms.join("+")}`);
    } else {
      goto(`/posts`);
    }
  };

  const getCounts = async () => {
    const response = await getPostCount();
    postCount = response.postCount;
  };
  onMount(() => {
    getCounts();
  });
</script>

<section class="hero is-primary is-fullheight-with-navbar">
  <div class="hero-body">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-12-tablet is-8-desktop is-8-widescreen">
          <div class="box has-text-centered">
            <p class="title">Shioriko</p>
            <p class="subtitle">Booru-style gallery written in Go and Svelte</p>
            <div class="block">
              <form onsubmit={onSearch}>
                <div class="field">
                  <div class="control has-text-left is-expanded" id="tags">
                    <Tags
                      tags={searchTerms}
                      addKeys={[9, 32]}
                      on:tags={onTagChange}
                      autoComplete={onAutocomplete}
                      autoCompleteFilter={false}
                    />
                  </div>
                </div>
                <div class="control">
                  <button type="submit" class="button is-primary">
                    Search
                  </button>
                </div>
              </form>
            </div>
            <p class="block">Serving <strong>{postCount}</strong> images</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>
