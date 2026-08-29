<script lang="ts">
    import coverImage from "$lib/../assets/images/parchments/wooden-floor-background.jpg";
    import Modal from "../../Modal.svelte";
    import { Faction } from "$wails/world-builder/internal/models/models";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";

    let {
        open = $bindable(false),
        onSave,
        factionData
    }: {
        open: boolean;
        onSave?: (data: { title: string; description: string }) => void;
        factionData: Faction | null
    } = $props();

    let title = $state('');
    let description = $state('');

    function handleSubmit(e: SubmitEvent) {
        e.preventDefault();
        onSave?.({ title, description });
        open = false;
    }

    function handleFactionChange(html: string) {
        description = html;
    }

    $effect(() => {
      if (open) {
        title = factionData?.name ?? ''
        description = factionData?.description ?? 'Describe your faction...'
        console.log(factionData)
      }

      if (!open) {
        factionData = null
        title = ''
        description = 'Describe your faction...'
      }
    })
</script>

<Modal bind:open={open}>
    <div
        class="fantasy-border p-6 bg-taupe-800 text-black"
        style="--inlay-bg: url('{coverImage}') center/cover; --inlay-filter: blur(1px) brightness(1);"
    >
        <form onsubmit={handleSubmit}>
            <input type="text" name="title" bind:value={title} placeholder="Faction" />

            <div class="text-left!">
                <RichEditor value={description} />
            </div>

            <button type="submit" class="fantasy-btn fantasy-bone-n-coper">Save</button>
        </form>
    </div>
</Modal>
