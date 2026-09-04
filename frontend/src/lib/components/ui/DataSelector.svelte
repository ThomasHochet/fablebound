<script lang="ts">
    type SelectValue = number | string;

    type SelectItem = {
        id: SelectValue;
        label: string;
    };

    let {
        items = [],
        value = $bindable(null),
        placeholder = "Select or forge a new category...",
        labelText = "Category",
        required = false,
        onCreate
    }: {
        items: SelectItem[];
        value: SelectValue | null;
        placeholder?: string;
        labelText?: string;
        required?: boolean;
        onCreate?: (newLabel: string) => Promise<SelectValue | null> | SelectValue | null;
    } = $props();

    let searchQuery = $state('');
    let isOpen = $state(false);
    let isCreating = $state(false);
    let inputRef: HTMLInputElement | null = $state(null);
    let lastSyncedValue: SelectValue | null = null;

    $effect(() => {
        if (value !== lastSyncedValue) {
            lastSyncedValue = value;
            if (value !== null && value !== '') {
                const selectedItem = items.find(i => i.id === value);
                // Displays label if found in items list, or raw string if custom
                searchQuery = selectedItem ? selectedItem.label : String(value);
            } else {
                searchQuery = '';
            }
        }
    });

    let filteredItems = $derived(
        items.filter(item =>
            item.label.toLowerCase().includes(searchQuery.toLowerCase())
        )
    );

    let exactMatch = $derived(
        items.some(item =>
            item.label.toLowerCase() === searchQuery.trim().toLowerCase()
        )
    );

    function handleSelect(item: SelectItem) {
        value = item.id;
        searchQuery = item.label;
        isOpen = false;
    }

    async function handleCreate() {
        const trimmed = searchQuery.trim();
        if (!trimmed || exactMatch || isCreating) return;

        isCreating = true;
        try {
            if (onCreate) {
                // Scribe mode: immediate creation via parent callback
                const newId = await onCreate(trimmed);
                if (newId !== null && newId !== undefined) {
                    value = newId;
                    isOpen = false;
                }
            } else {
                // Local mode: assign raw string directly to form state
                value = trimmed;
                isOpen = false;
            }
        } catch (err) {
            console.error("Failed to forge new entry", err);
        } finally {
            isCreating = false;
        }
    }

    function handleWindowClick(e: MouseEvent) {
        if (inputRef && !inputRef.parentElement?.contains(e.target as Node)) {
            isOpen = false;
            if (value !== null && value !== '') {
                const selectedItem = items.find(i => i.id === value);
                searchQuery = selectedItem ? selectedItem.label : String(value);
            } else {
                searchQuery = '';
            }
        }
    }
</script>

<svelte:window onclick={handleWindowClick} />

<div class="relative w-full text-[var(--ink-color)] font-serif">
    <label class="block text-sm font-cinzel text-[var(--gold-text)] mb-1 tracking-wider uppercase text-left ml-2">{labelText}</label>

    <div class="fantasy-input-wrapper flex items-center">
        <input
            type="text"
            class="sr-only absolute bottom-0 left-1/2 opacity-0 pointer-events-none"
            tabindex="-1"
            value={value ?? ''}
            {required}
            oninvalid={() => {
                inputRef?.focus();
            }}
        />
        <input
            bind:this={inputRef}
            type="text"
            bind:value={searchQuery}
            onfocus={() => isOpen = true}
            onkeydown={(e) => {
                if (e.key === 'Enter') {
                    e.preventDefault();
                    if (!exactMatch && searchQuery) handleCreate();
                    else if (filteredItems.length > 0) handleSelect(filteredItems[0]);
                }
                if (e.key === 'Escape') isOpen = false;
            }}
            placeholder={placeholder}
            class="fantasy-input pr-8"
            disabled={isCreating}
        />

        <button
            type="button"
            onclick={(e) => { e.stopPropagation(); isOpen = !isOpen; }}
            class="absolute right-3 top-1/2 -translate-y-1/2 !text-[var(--wood-base)] hover:text-[var(--wood-dark)] focus:outline-none"
        >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
            </svg>
        </button>
    </div>

    {#if isOpen}
        <div class="absolute z-50 w-full mt-1 bg-[var(--parchment-base)] border-2 border-[var(--wood-dark)] rounded-b-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar">
            {#if filteredItems.length > 0}
                <ul class="py-1">
                    {#each filteredItems as item}
                        <li>
                            <button
                                type="button"
                                class="dropdown-item-btn w-full text-left px-4 py-2 hover:bg-[var(--gold-text)]/20 focus:bg-[var(--gold-text)]/20 transition-colors border-b border-dashed border-[var(--parchment-dark)]/50 !text-[var(--ink-color)]"
                                onclick={() => handleSelect(item)}
                            >
                                {item.label}
                            </button>
                        </li>
                    {/each}
                </ul>
            {/if}

            {#if searchQuery.trim() !== '' && !exactMatch}
                <div class="border-t border-[var(--parchment-dark)]/50 py-1 bg-[var(--parchment-base)]">
                    <button
                        type="button"
                        class="w-full text-left px-4 py-2 !text-[var(--wood-light)] hover:bg-[var(--gold-text)]/25 font-bold italic transition-colors flex justify-between items-center"
                        onclick={handleCreate}
                        disabled={isCreating}
                    >
                        {#if onCreate}
                            <span>+ Forge "{searchQuery}" now</span>
                            {#if isCreating}
                                <span class="text-sm opacity-70">Scribing...</span>
                            {/if}
                        {:else}
                            <span>+ Mark as "{searchQuery}" (forged upon saving)</span>
                        {/if}
                    </button>
                </div>
            {:else if filteredItems.length === 0}
                <div class="px-4 py-3 text-sm italic !text-[var(--wood-light)]/70">
                    Awaits your penstrokes to create...
                </div>
            {/if}
        </div>
    {/if}
</div>

<style>
    .dropdown-item-btn {
        color: var(--ink-color) !important;
    }

    .custom-scrollbar::-webkit-scrollbar {
        width: 6px;
    }
    .custom-scrollbar::-webkit-scrollbar-track {
        background: var(--parchment-base);
        border-left: 1px solid var(--parchment-dark);
    }
    .custom-scrollbar::-webkit-scrollbar-thumb {
        background: var(--wood-light);
        border-radius: 4px;
    }
</style>
