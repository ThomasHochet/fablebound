<script lang='ts'>
    import FactionOverview from "../layout/Overviews/FactionOverview.svelte";
    import LocationOverview from "../layout/Overviews/LocationOverview.svelte";
    import LoreOverview from "../layout/Overviews/LoreOverview.svelte";
    import ArticleForm from "./forms/ArticleForm.svelte";
    import CharacterForm from "./forms/CharacterForm.svelte";
    import FactionForm from "./forms/FactionForm.svelte";
    import LocationForm from "./forms/LocationForm.svelte";
    import LoreForm from "./forms/LoreForm.svelte";
    import ParametersForm from "./forms/ParametersForm.svelte";

    let { section } = $props();
</script>

<section>
    {#if section === 'character'}
        <CharacterForm />
    {:else if section === 'lore'}
        <LoreOverview />
    {:else if section.startsWith('lore/form')}
        {@const formId = section.split('lore/form/')[1]}
        <LoreForm id={formId} />
    {:else if section === 'location'}
        <LocationOverview />
    {:else if section.startsWith('location/form')}
        {@const formId = section.split('location/form/')[1]}
        <LocationForm id={formId} />
    {:else if section === 'faction'}
        <FactionOverview />
    {:else if section.startsWith('faction/form')}
        {@const formId = section.split('faction/form/')[1]}
        <FactionForm id={formId} />
    {:else if section === 'world'}
        <ArticleForm />
    {:else if section === 'params'}
        <ParametersForm />
    {:else}
        <ArticleForm />
    {/if}
</section>
