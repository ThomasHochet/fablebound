<script lang='ts'>

    import ArticleForm from "./forms/ArticleForm.svelte";
    import CharacterForm from "./forms/CharacterForm.svelte";
    import FactionForm from "./forms/FactionForm.svelte";
    import LocationForm from "./forms/LocationForm.svelte";
    import LoreForm from "./forms/LoreForm.svelte";
    import ParametersForm from "./forms/ParametersForm.svelte";

    let { section } = $props();
</script>

<section>
    {#if section.startsWith('character/form')}
        {@const formId = section.split('character/form/')[1]}
        <CharacterForm id={formId} />
    {:else if section.startsWith('lore/form')}
        {@const formId = section.split('lore/form/')[1]}
        <LoreForm id={formId} />
    {:else if section.startsWith('location/form')}
        {@const formId = section.split('location/form/')[1]}
        <LocationForm id={formId} />
    {:else if section.startsWith('faction/form')}
        {@const formId = section.split('faction/form/')[1]}
        <FactionForm id={formId} />
    {:else if section.startsWith('world/form')}
        {@const formId = section.split('world/form/')[1]}
        <ArticleForm id={formId} />
    {:else if section === 'params'}
        <ParametersForm />
    {/if}
</section>
