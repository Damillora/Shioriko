<script lang="ts">
  import Tags from "svelte-tags-input";
  import logo from "$lib/assets/phoebe-logo.svg";
  import { getPostCount, getPosts, getTagAutocomplete } from "$lib/api";

  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import PostGallery from "$lib/components/ui/PostGallery.svelte";

  let searchTerms: string[] = $state([]);
  let postCount: number = $state(0);
  let postCountLoaded: boolean = $state(false);

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
    postCountLoaded = true;
  };
  onMount(() => {
    getCounts();
  });
</script>

<section class="hero is-fullheight-with-navbar">
  <div class="hero-body">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-12-tablet is-8-desktop is-8-widescreen">
          <div class="block is-flex is-justify-content-center">
            <figure class="image is-128x128">
              <img alt="phoebe logo" src={logo} />
            </figure>
          </div>
          <div class="block has-text-centered">
            <p class="title">phoebe</p>
            <p class="subtitle">a booru-style image gallery and organizer</p>
          </div>
          <div class="box has-text-centered">
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
                <div class="field">
                  <div class="control">
                    <button type="submit" class="button is-primary">
                      Search
                    </button>
                  </div>
                </div>
              </form>
            </div>
            {#if postCountLoaded}
              <p class="block">
                serving <span class="is-primary"
                  ><strong>{postCount}</strong></span
                > images
              </p>
            {:else}
              <p class="block">
                serving <span class="is-primary"><strong>...</strong></span> images
              </p>
            {/if}
          </div>
        </div>
      </div>
    </div>
  </div>
</section>
