<script lang="ts">
    type SelectValue = number | string;

    type SelectItem = {
        id: SelectValue;
        label?: string;
        name?: string;
    };

    type SelectVariant = 'fantasy' | 'header' | 'default' | 'forge';

    let {
        items = [],
        value = $bindable(null),
        placeholder = "Select or forge a new category...",
        labelText = "Category",
        required = false,
        selectOnly = false,
        disabled = false,
        variant = 'fantasy',
        onCreate
    }: {
        items: SelectItem[];
        value: SelectValue | null;
        placeholder?: string;
        labelText?: string;
        required?: boolean;
        selectOnly?: boolean;
        disabled?: boolean;
        variant?: SelectVariant;
        onCreate?: (newLabel: string) => Promise<SelectValue | null> | SelectValue | null;
    } = $props();

    let searchQuery = $state('');
    let isOpen = $state(false);
    let isCreating = $state(false);
    let inputRef: HTMLInputElement | null = $state(null);
    let lastSyncedValue: SelectValue | null = null;

    const themes = {
        forge: {
            wrapper: "relative w-full font-serif grid grid-rows-[auto-auto]",
            // label: "block text-xs font-serif text-[var(--ink-color)] mb-1 text-left ml-1",
            label: "forge-input-label text-left!",
            inputWrapper: "relative flex items-center",
            input: "w-full appearance-none bg-[#fcf8f2]/40 border border-[#b7a48d] rounded-md pl-2 pr-8 py-1.5 text-base font-serif text-[var(--ink-color)] placeholder-[#9c8974] shadow-inner focus:outline-none",
            // input: "forge-input w-full!",
            icon: "absolute right-2.5 top-1/2 -translate-y-1/2 text-(--ink-color)/60! hover:text-[#523d2b] focus:outline-none cursor-pointer",
            dropdown: "absolute top-full left-0 z-50 w-full bg-[#f5eedf] border border-[#b7a48d] rounded-md shadow-lg max-h-60 overflow-y-auto overflow-x-hidden custom-scrollbar backdrop-blur-sm",
            item: "w-full text-left px-3 py-2 text-sm font-serif text-(--ink-color)! hover:bg-[#e6d3b3]/50 transition-colors border-b border-[#b7a48d]/20 cursor-pointer first:rounded-t-md last:rounded-b-md",
            createBg: "border-t border-[#b7a48d]/40 py-1 bg-[#f5eedf] rounded-b-md",
            createBtn: "w-full text-left px-3 py-1.5 text-sm font-serif text-[#3b2a1e] hover:bg-[#e6d3b3] font-bold italic transition-colors flex justify-between items-center cursor-pointer rounded-b-md",
            empty: "px-3 py-2 text-xs italic text-[#8c7661]"
        },
        fantasy: {
            wrapper: "relative w-full text-[var(--ink-color)] font-serif",
            label: "block text-sm font-cinzel text-[var(--gold-text)] mb-1 tracking-wider uppercase text-left ml-2",
            inputWrapper: "fantasy-input-wrapper flex items-center relative",
            input: "fantasy-input pr-8",
            icon: "absolute right-3 top-1/2 -translate-y-1/2 !text-[var(--wood-base)] hover:text-[var(--wood-dark)] focus:outline-none",
            dropdown: "absolute z-50 w-full mt-1 bg-[var(--parchment-base)] border-2 border-[var(--wood-dark)] rounded-b-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar",
            item: "dropdown-item-btn w-full text-left px-4 py-2 hover:bg-[var(--gold-text)]/20 focus:bg-[var(--gold-text)]/20 transition-colors border-b border-dashed border-[var(--parchment-dark)]/50 !text-[var(--ink-color)]",
            createBg: "border-t border-[var(--parchment-dark)]/50 py-1 bg-[var(--parchment-base)]",
            createBtn: "w-full text-left px-4 py-2 !text-[var(--wood-light)] hover:bg-[var(--gold-text)]/25 font-bold italic transition-colors flex justify-between items-center",
            empty: "px-4 py-3 text-sm italic !text-[var(--wood-light)]/70"
        },
        header: {
            wrapper: "relative w-full font-serif",
            label: "block text-xs font-serif text-[var(--ink-color)] mb-1 text-left ml-1",
            inputWrapper: "relative flex items-center",
            input: "w-full bg-[#fcf8f2]/60 border border-[#b7a48d] rounded-md pl-3 pr-8 py-1.5 text-sm font-serif text-[var(--ink-color)] shadow-inner focus:outline-none placeholder-[#9c8974]",
            icon: "absolute right-2.5 top-1/2 -translate-y-1/2 text-[#8c7661] hover:text-[#523d2b] focus:outline-none",
            dropdown: "absolute z-50 w-full mt-1 bg-[#f5eedf] border border-[#b7a48d] rounded-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar backdrop-blur-sm",
            item: "w-full text-left px-3 py-1.5 text-sm font-serif text-[#3b2a1e] hover:bg-[#e6d3b3] transition-colors border-b border-[#b7a48d]/20",
            createBg: "border-t border-[#b7a48d]/40 py-1 bg-[#f5eedf]",
            createBtn: "w-full text-left px-3 py-1.5 text-sm font-serif text-[#3b2a1e] hover:bg-[#e6d3b3] font-bold italic transition-colors flex justify-between items-center",
            empty: "px-3 py-2 text-xs italic text-[#8c7661]"
        },
        default: {
            wrapper: "relative w-full text-gray-800 font-sans",
            label: "block text-xs font-medium text-gray-700 mb-1 text-left ml-0.5",
            inputWrapper: "relative flex items-center",
            input: "w-full bg-white border border-gray-300 rounded-md pl-3 pr-8 py-1.5 text-sm text-gray-900 shadow-sm focus:ring-2 focus:ring-amber-500/50 focus:border-amber-500 focus:outline-none placeholder-gray-400",
            icon: "absolute right-2.5 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 focus:outline-none",
            dropdown: "absolute z-50 w-full mt-1 bg-white border border-gray-200 rounded-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar",
            item: "w-full text-left px-3 py-1.5 text-sm text-gray-700 hover:bg-gray-100 transition-colors border-b border-gray-100",
            createBg: "border-t border-gray-100 py-1 bg-gray-50",
            createBtn: "w-full text-left px-3 py-1.5 text-sm text-amber-600 hover:bg-amber-50 font-medium italic transition-colors flex justify-between items-center",
            empty: "px-3 py-2 text-xs italic text-gray-400"
        }
    };

    let currentTheme = $derived(themes[variant] ?? themes.fantasy);

    function getItemLabel(item: SelectItem): string {
        return item.label ?? item.name ?? '';
    }

    let selectedItemLabel = $derived.by(() => {
        if (value !== null && value !== '') {
            const selectedItem = items.find(i => i.id === value);
            return selectedItem ? getItemLabel(selectedItem) : String(value);
        }
        return '';
    });

    $effect(() => {
        if (value !== lastSyncedValue) {
            lastSyncedValue = value;
            searchQuery = selectedItemLabel;
        }
    });

    let filteredItems = $derived(
        items.filter(item => {
            if ((variant === 'header' || variant === 'forge') && searchQuery === selectedItemLabel) {
                return true;
            }
            return getItemLabel(item).toLowerCase().includes(searchQuery.toLowerCase());
        })
    );

    let exactMatch = $derived(
        items.some(item =>
            getItemLabel(item).toLowerCase() === searchQuery.trim().toLowerCase()
        )
    );

    function handleSelect(item: SelectItem) {
        value = item.id;
        searchQuery = getItemLabel(item);
        isOpen = false;
    }

    async function handleCreate() {
        const trimmed = searchQuery.trim();
        if (!trimmed || exactMatch || isCreating) return;

        isCreating = true;
        try {
            if (onCreate) {
                const newId = await onCreate(trimmed);
                if (newId !== null && newId !== undefined) {
                    value = newId;
                    isOpen = false;
                }
            } else {
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
            if (searchQuery.trim() === '') {
              value = null;
            }
            searchQuery = selectedItemLabel;
        }
    }
</script>

<svelte:window onclick={handleWindowClick} />

<div class={currentTheme.wrapper}>
    {#if labelText}
        <label class={currentTheme.label}>{labelText}</label>
    {/if}

    <div class={currentTheme.inputWrapper}>
        <input
            type="text"
            class="sr-only absolute bottom-0 left-1/2 opacity-0 pointer-events-none"
            tabindex="-1"
            value={value ?? ''}
            {required}
            {disabled}
            oninvalid={() => inputRef?.focus()}
        />

        <input
            bind:this={inputRef}
            type="text"
            bind:value={searchQuery}
            oninput={() => {
              if (searchQuery.trim() === '') {
                value = null
              }
            }}
            onfocus={(e) => {
                isOpen = true;
                if (variant === 'header' || variant === 'forge') e.currentTarget.select();
            }}
            onkeydown={(e) => {
                if (e.key === 'Enter') {
                    e.preventDefault();
                    if (!exactMatch && searchQuery && !selectOnly) handleCreate();
                    else if (filteredItems.length > 0) handleSelect(filteredItems[0]);
                }
                if (e.key === 'Escape') isOpen = false;
            }}
            {placeholder}
            class={currentTheme.input}
            disabled={disabled || isCreating}
        />

        <button
            type="button"
            onclick={(e) => { e.stopPropagation(); isOpen = !isOpen; }}
            class={currentTheme.icon}
        >
            <svg class="w-4 h-4 transition-transform duration-200 {isOpen ? 'rotate-180' : ''}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
            </svg>
        </button>
    </div>

    {#if isOpen}
        <div class={currentTheme.dropdown}>
            {#if filteredItems.length > 0}
                <ul class="py-0 z-500">
                    {#each filteredItems as item}
                        <li class="z-200">
                            <button
                                type="button"
                                class="z-200 {currentTheme.item} {value === item.id ? ((variant === 'header' || variant === 'forge') ? 'font-bold bg-[#e6d3b3]/60' : 'font-bold') : ''}"
                                onclick={() => handleSelect(item)}
                            >
                                {getItemLabel(item)}
                            </button>
                        </li>
                    {/each}
                </ul>
            {/if}

            {#if searchQuery.trim() !== '' && !exactMatch && !selectOnly}
                <div class={currentTheme.createBg}>
                    <button
                        type="button"
                        class={currentTheme.createBtn}
                        onclick={handleCreate}
                        disabled={isCreating}
                    >
                        {#if onCreate}
                            <span>+ Forge "{searchQuery}" now</span>
                            {#if isCreating}
                                <span class="text-xs opacity-70">Scribing...</span>
                            {/if}
                        {:else}
                            <span>+ Mark as "{searchQuery}" (forged upon saving)</span>
                        {/if}
                    </button>
                </div>
            {:else if filteredItems.length === 0}
                <div class={currentTheme.empty}>
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
        background: #f5eedf;
    }
    .custom-scrollbar::-webkit-scrollbar-thumb {
        background: #b7a48d;
        border-radius: 4px;
    }
</style>
