<script lang="ts">
    import type { Snippet } from 'svelte';

    type TableItem = {
        id: number | string;
        label: string;
    };

    let {
        items = $bindable([]),
        labelText = "Manage Category",
        placeholder = "Type and press Enter...",
        onAdd,
        onUpdate,
        actionCell
    }: {
        items: TableItem[];
        labelText?: string;
        placeholder?: string;
        /** Callback when a new item is submitted. Should return the created item with its DB ID. */
        onAdd?: (label: string) => Promise<TableItem | null> | TableItem | null;
        /** Callback when an existing item is updated. Should return true on success. */
        onUpdate?: (id: number | string, label: string) => Promise<boolean | null> | boolean | null;
        /** Svelte Snippet for injecting your custom action button (e.g., delete) into the row */
        actionCell?: Snippet<[TableItem]>;
    } = $props();

    let inputValue = $state('');
    let editingId = $state<number | string | null>(null);
    let isProcessing = $state(false);
    let inputRef: HTMLInputElement | null = $state(null);

    async function handleKeydown(e: KeyboardEvent) {
        if (e.key === 'Escape') {
            cancelEdit();
            return;
        }

        if (e.key === 'Enter') {
            e.preventDefault();
            const trimmed = inputValue.trim();
            if (!trimmed || isProcessing) return;

            isProcessing = true;

            try {
                if (editingId !== null) {
                    // --- UPDATE MODE ---
                    const success = onUpdate ? await onUpdate(editingId, trimmed) : true;
                    if (success) {
                        const index = items.findIndex(i => i.id === editingId);
                        if (index !== -1) items[index].label = trimmed;
                        cancelEdit();
                    }
                } else {
                    // --- CREATE MODE ---
                    // Check for exact duplicates locally before adding
                    if (items.some(i => i.label.toLowerCase() === trimmed.toLowerCase())) {
                        isProcessing = false;
                        return;
                    }

                    if (onAdd) {
                        const newItem = await onAdd(trimmed);
                        if (newItem) {
                            items = [...items, newItem];
                            inputValue = '';
                        }
                    } else {
                        // Fallback local mode (if no DB callback provided, assign a temp ID)
                        items = [...items, { id: Date.now(), label: trimmed }];
                        inputValue = '';
                    }
                }
            } catch (err) {
                console.error("Action failed:", err);
            } finally {
                isProcessing = false;
            }
        }
    }

    function selectForEdit(item: TableItem) {
        editingId = item.id;
        inputValue = item.label;
        inputRef?.focus();
    }

    function cancelEdit() {
        editingId = null;
        inputValue = '';
    }
</script>

<div class="flex flex-col w-full font-serif h-full">
    <div class="flex justify-between items-end mb-1 ml-2 pr-2">
        <label class="block text-sm font-cinzel text-[var(--gold-text)] tracking-wider uppercase text-left">
            {labelText}
        </label>
        {#if editingId !== null}
            <button
                type="button"
                class="text-xs italic text-[var(--wood-light)] hover:text-[var(--gold-text)] transition-colors"
                onclick={cancelEdit}
            >
                Cancel Edit (Esc)
            </button>
        {/if}
    </div>

    <!-- Input Area -->
    <div class="fantasy-input-wrapper flex items-center relative transition-all {editingId !== null ? 'ring-1 ring-[var(--gold-accent)]' : ''}">
        <input
            bind:this={inputRef}
            type="text"
            bind:value={inputValue}
            onkeydown={handleKeydown}
            {placeholder}
            disabled={isProcessing}
            class="fantasy-input pr-8 {isProcessing ? 'opacity-70' : ''}"
        />
        <!-- Indicator icon for edit mode -->
        {#if editingId !== null}
            <div class="absolute right-3 top-1/2 -translate-y-1/2 text-[var(--gold-accent)] opacity-80 pointer-events-none">
                ✎
            </div>
        {/if}
    </div>

    <!-- Table Area -->
    <div class="mt-2 flex-grow overflow-y-auto custom-scrollbar border border-[var(--wood-dark)] bg-[var(--parchment-base)] rounded-sm shadow-inner min-h-[140px] max-h-[140px]">
        {#if items.length > 0}
            <table class="w-full text-left text-sm border-collapse">
                <tbody>
                    {#each items as item}
                        <tr
                            class="group border-b border-dashed border-[var(--parchment-dark)]/50 transition-colors cursor-pointer
                                   {editingId === item.id ? 'bg-[var(--gold-text)]/20' : 'hover:bg-[var(--gold-text)]/10'}"
                            onclick={() => selectForEdit(item)}
                        >
                            <td class="px-3 py-2 text-[var(--ink-color)] truncate max-w-0 w-full">
                                {item.label}
                            </td>

                            <!-- Custom Action Button Column -->
                            {#if actionCell}
                                <td class="px-2 py-1 w-10 text-right align-middle" onclick={(e) => e.stopPropagation()}>
                                    {@render actionCell(item)}
                                </td>
                            {/if}
                        </tr>
                    {/each}
                </tbody>
            </table>
        {:else}
            <div class="p-4 text-center text-sm italic !text-[var(--wood-light)]/70 h-full flex items-center justify-center">
                No entries inscribed yet.
            </div>
        {/if}
    </div>
</div>

<style>
    .custom-scrollbar::-webkit-scrollbar { width: 6px; }
    .custom-scrollbar::-webkit-scrollbar-track { background: var(--parchment-base); }
    .custom-scrollbar::-webkit-scrollbar-thumb { background: var(--wood-light); border-radius: 4px; }
</style>
