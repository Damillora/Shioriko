<script lang="ts">
  import Tags from "svelte-tags-input";
  import { getTagAutocomplete } from "$lib/api";

	import { goto } from '$app/navigation';

  let searchTerms: string[] = [];

  const onTagChange = (value) => {
    searchTerms = value.detail.tags;
  };

  const onAutocomplete = async (tag) => {
        const list = await getTagAutocomplete({ tag });
        return list;
    };

  const onSearch = (i) => {
    if (searchTerms.length > 0) {
      goto(`/posts?tags=${searchTerms.join("+")}`);
    } else {
      goto(`/posts`);
    }
  };
</script>

<section class="hero is-small">
  <div class="hero-body">
    <div class="container has-text-centered">
      <p class="title">Shioriko</p>
      <p class="subtitle">Booru-style gallery written in Go and Svelte</p>
    </div>
  </div>
  <div class="hero-foot">
    <div class="container has-text-centered">
      <form on:submit|preventDefault={onSearch}>
        <div class="field has-addons">
          <div class="control has-text-left is-expanded">
            <div class="control" id="tags">
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
            <button type="submit" class="button is-primary"> Search </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</section>
