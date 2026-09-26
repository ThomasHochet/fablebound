<script lang='ts'>
    import ArticleForm from "./forms/ArticleForm.svelte";
    import CharacterForm from "./forms/CharacterForm.svelte";
    import FactionForm from "./forms/FactionForm.svelte";
    import LocationForm from "./forms/LocationForm.svelte";
    import LoreForm from "./forms/LoreForm.svelte";
    import ParametersForm from "./forms/ParametersForm.svelte";

    let { params } = $props<{
        params: {
          section?: string
          id?: string
        }
      }>()

    let section = $derived(params?.section)
    let id = $derived(params?.id)

</script>

<section>
    {#if section === 'character'}
        <CharacterForm {id} />
    {:else if section === 'lore'}
        <LoreForm {id} />
    {:else if section === 'location'}
        <LocationForm {id} />
    {:else if section === 'faction'}
        <FactionForm {id} />
    {:else if section === 'world'}
        <ArticleForm {id} />
    {:else if section === 'params'}
        <ParametersForm />
    {/if}
</section>
