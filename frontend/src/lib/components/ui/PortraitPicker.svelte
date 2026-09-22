<script lang="ts">
    let { characterId, currentPortrait = $bindable("") } = $props();

    let isSaving = $state(false);
    let fileInputRef: HTMLInputElement | null = $state(null);

    function handleSelectClick() {
        fileInputRef?.click();
    }

    function handleFileChange(e: Event) {
        const input = e.target as HTMLInputElement;
        const file = input.files?.[0];
        if (!file || !file.type.startsWith("image/")) return;

        isSaving = true;

        const reader = new FileReader();
        reader.onloadend = async () => {
            currentPortrait = reader.result as string;
            isSaving = false;

            // You can either trigger a direct Go save here,
            // or let the parent form handle the save via formState.portrait
        };
        reader.readAsDataURL(file);
    }
</script>

<div class="flex flex-col items-center gap-2">
    <input
        bind:this={fileInputRef}
        type="file"
        accept="image/png, image/jpeg, image/webp"
        class="hidden"
        onchange={handleFileChange}
    />

    <div
        role="button"
        tabindex="0"
        onkeydown={(e) => e.key === 'Enter' && handleSelectClick()}
        onclick={handleSelectClick}
        class="relative w-full h-52 border-2 border-dashed border-[var(--wood-dark)] bg-transparent rounded overflow-hidden flex flex-col items-center justify-center cursor-pointer hover:border-[var(--gold-text)] transition-colors"
    >
        {#if currentPortrait}
            <img src={currentPortrait} alt="Portrait" class="w-full h-full object-contain p-2" />

            <div class="absolute inset-0 bg-black/40 opacity-0 hover:opacity-100 flex items-center justify-center transition-opacity">
                <span class="text-[var(--gold-text)] text-xs uppercase font-bold tracking-wider">Change</span>
            </div>
        {:else}
            <span class="text-xs text-[var(--ink-color)] uppercase tracking-wider text-center px-4">
                Click to Select<br/>Image
            </span>
        {/if}

        {#if isSaving}
            <div class="absolute inset-0 bg-black/60 flex items-center justify-center">
                <span class="text-white text-xs animate-pulse">Loading...</span>
            </div>
        {/if}
    </div>
</div>
