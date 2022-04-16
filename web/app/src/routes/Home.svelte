<script>
  import Tags from "svelte-tags-input";
  import { getTagAutocomplete } from "../api.js";
  import { navigate } from "svelte-routing";

  let searchTerms = [];

  const onTagChange = (value) => {
    searchTerms = value.detail.tags;
  };

  const onAutocomplete = async (tag) => {
        const list = await getTagAutocomplete({ tag });
        return list;
    };

  const onSearch = (i) => {
    if (searchTerms.length > 0) {
      navigate(`/posts?tags=${searchTerms.join("+")}`);
    } else {
      navigate(`/posts`);
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
