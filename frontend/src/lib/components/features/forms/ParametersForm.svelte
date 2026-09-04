<script lang="ts">
    import Divider from "$lib/components/ui/Divider.svelte";
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import LookupForms from "./SubForms/LookupForms.svelte";
    import TraitForms from "./SubForms/TraitForms.svelte";



</script>

<div class="h-screen overflow-y-auto">
    <div class="relative z-20 -mb-15 pointer-events-none">
        <ParchmentTitle title="Miscellaneous" subtitle="Lookups, Traits and more..." />
    </div>
    <div class="relative z-10">
        <LookupForms />
        <Divider />
        <TraitForms />
    </div>

</div>
