<script lang="ts">
    import AffiliationReader from "./readers/AffiliationReader.svelte";
    import CharacterReader from "./readers/CharacterReader.svelte";
    import FactionReader from "./readers/FactionReader.svelte";
    import LocationReader from "./readers/LocationReader.svelte";
    import LoreReader from "./readers/LoreReader.svelte";
    import WorldReader from "./readers/WorldReader.svelte";

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
    {#if section === 'world'}
        <WorldReader {id} />
    {:else if section === 'location'}
        <LocationReader {id} />
    {:else if section === 'character'}
        <CharacterReader {id} />
    {:else if section === 'faction'}
        <FactionReader {id} />
    {:else if section === 'affiliation'}
        <AffiliationReader {id} />
    {:else if section === 'lore'}
        <LoreReader {id} />
    {/if}
</section>
