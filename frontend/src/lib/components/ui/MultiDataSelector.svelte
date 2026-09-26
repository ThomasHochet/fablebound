<script lang="ts">
    import { logError } from "$lib/logger";

    type SelectValue = number | string;

    type SelectItem = {
        id: SelectValue;
        label?: string;
        name?: string;
    };

    type SelectVariant = 'fantasy' | 'header' | 'default' | 'forge';

    let {
        items = [],
        value = $bindable([]),
        placeholder = "Search or forge tags...",
        labelText = "Categories",
        required = false,
        selectOnly = false,
        disabled = false,
        variant = 'fantasy',
        onCreate
    }: {
        items: SelectItem[];
        value: SelectValue[];
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

    const themes = {
        forge: {
            wrapper: "relative w-full font-serif grid grid-rows-[auto-auto]",
            label: "forge-input-label text-left!",
            inputWrapper: "relative flex flex-wrap items-center gap-1.5 min-h-[40px] bg-[#fcf8f2]/40 border border-[#b7a48d] rounded-md px-2 py-1 pr-9 shadow-inner focus-within:ring-1 focus-within:ring-[#b7a48d]",
            pill: "flex items-center gap-1 bg-[#e6d3b3] text-[#3b2a1e] px-2 py-0.5 rounded text-xs font-serif font-semibold border border-[#b7a48d]",
            pillRemove: "hover:text-red-700 focus:outline-none ml-1 transition-colors cursor-pointer text-(--ink-color)!",
            input: "flex-1 bg-transparent border-none outline-none text-[var(--ink-color)] placeholder-[#9c8974] min-w-[120px] py-0.5 text-base font-serif",
            icon: "absolute right-2.5 top-2.5 text-(--ink-color)/60! hover:text-[#523d2b] focus:outline-none cursor-pointer",
            dropdown: "absolute top-full left-0 z-50 w-full mt-1 bg-[#f5eedf] border border-[#b7a48d] rounded-md shadow-lg max-h-60 overflow-y-auto overflow-x-hidden custom-scrollbar backdrop-blur-sm",
            item: "w-full text-left px-3 py-2 text-sm font-serif text-(--ink-color)! hover:bg-[#e6d3b3]/50 transition-colors border-b border-[#b7a48d]/20 cursor-pointer first:rounded-t-md last:rounded-b-md",
            createBg: "border-t border-[#b7a48d]/40 py-1 bg-[#f5eedf] rounded-b-md",
            createBtn: "w-full text-left px-3 py-1.5 text-sm font-serif text-[#3b2a1e] hover:bg-[#e6d3b3] font-bold italic transition-colors flex justify-between items-center cursor-pointer rounded-b-md",
            empty: "px-3 py-2 text-xs italic text-[#8c7661]"
        },
        fantasy: {
            wrapper: "relative w-full text-[var(--ink-color)] font-serif",
            label: "block text-sm font-cinzel text-[var(--gold-text)] mb-1 tracking-wider uppercase text-left ml-2",
            inputWrapper: "fantasy-input-wrapper flex flex-wrap items-center gap-1.5 p-1.5 min-h-[42px] cursor-text relative pr-9",
            pill: "flex items-center gap-1 bg-[var(--wood-dark)] text-[var(--gold-text)] px-2 py-0.5 rounded text-sm font-bold border border-[var(--wood-light)]/50",
            pillRemove: "hover:text-red-400 focus:outline-none ml-1 transition-colors cursor-pointer",
            input: "flex-1 bg-transparent border-none outline-none text-[var(--ink-color)] placeholder:text-[var(--gold-text)]/50 min-w-[120px] px-1 py-0.5 font-serif",
            icon: "absolute right-3 top-2.5 !text-[var(--wood-base)] hover:text-[var(--wood-dark)] focus:outline-none cursor-pointer",
            dropdown: "absolute z-50 w-full mt-1 bg-[var(--parchment-base)] border-2 border-[var(--wood-dark)] rounded-b-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar",
            item: "dropdown-item-btn w-full text-left px-4 py-2 hover:bg-[var(--gold-text)]/20 focus:bg-[var(--gold-text)]/20 transition-colors border-b border-dashed border-[var(--parchment-dark)]/50 !text-[var(--ink-color)] cursor-pointer",
            createBg: "border-t border-[var(--parchment-dark)]/50 py-1 bg-[var(--parchment-base)]",
            createBtn: "w-full text-left px-4 py-2 !text-[var(--wood-light)] hover:bg-[var(--gold-text)]/25 font-bold italic transition-colors flex justify-between items-center cursor-pointer",
            empty: "px-4 py-3 text-sm italic !text-[var(--wood-light)]/70"
        },
        header: {
            wrapper: "relative w-full font-serif",
            label: "block text-xs font-serif text-[var(--ink-color)] mb-1 text-left ml-1",
            inputWrapper: "relative flex flex-wrap items-center gap-1.5 bg-[#fcf8f2]/60 border border-[#b7a48d] rounded-md px-2 py-1 pr-8 text-sm font-serif text-[var(--ink-color)] shadow-inner min-h-[36px]",
            pill: "flex items-center gap-1 bg-[#e6d3b3] text-[#3b2a1e] px-2 py-0.5 rounded text-xs font-semibold border border-[#b7a48d]",
            pillRemove: "hover:text-red-700 focus:outline-none ml-1 transition-colors cursor-pointer",
            input: "flex-1 bg-transparent border-none outline-none text-[var(--ink-color)] placeholder-[#9c8974] min-w-[100px] py-0.5 text-xs font-serif",
            icon: "absolute right-2.5 top-2 text-[#8c7661] hover:text-[#523d2b] focus:outline-none cursor-pointer",
            dropdown: "absolute z-50 w-full mt-1 bg-[#f5eedf] border border-[#b7a48d] rounded-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar backdrop-blur-sm",
            item: "w-full text-left px-3 py-1.5 text-sm font-serif text-[#3b2a1e] hover:bg-[#e6d3b3] transition-colors border-b border-[#b7a48d]/20 cursor-pointer",
            createBg: "border-t border-[#b7a48d]/40 py-1 bg-[#f5eedf]",
            createBtn: "w-full text-left px-3 py-1.5 text-sm font-serif text-[#3b2a1e] hover:bg-[#e6d3b3] font-bold italic transition-colors flex justify-between items-center cursor-pointer",
            empty: "px-3 py-2 text-xs italic text-[#8c7661]"
        },
        default: {
            wrapper: "relative w-full text-gray-800 font-sans",
            label: "block text-xs font-medium text-gray-700 mb-1 text-left ml-0.5",
            inputWrapper: "relative flex flex-wrap items-center gap-1.5 bg-white border border-gray-300 rounded-md px-2 py-1 pr-8 text-sm shadow-sm focus-within:ring-2 focus-within:ring-amber-500/50 focus-within:border-amber-500 min-h-[38px]",
            pill: "flex items-center gap-1 bg-amber-100 text-amber-900 px-2 py-0.5 rounded text-xs font-medium border border-amber-200",
            pillRemove: "hover:text-amber-700 focus:outline-none ml-1 transition-colors cursor-pointer",
            input: "flex-1 bg-transparent border-none outline-none text-gray-900 placeholder-gray-400 min-w-[100px] py-0.5 text-sm",
            icon: "absolute right-2.5 top-2.5 text-gray-400 hover:text-gray-600 focus:outline-none cursor-pointer",
            dropdown: "absolute z-50 w-full mt-1 bg-white border border-gray-200 rounded-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar",
            item: "w-full text-left px-3 py-1.5 text-sm text-gray-700 hover:bg-gray-100 transition-colors border-b border-gray-100 cursor-pointer",
            createBg: "border-t border-gray-100 py-1 bg-gray-50",
            createBtn: "w-full text-left px-3 py-1.5 text-sm text-amber-600 hover:bg-amber-50 font-medium italic transition-colors flex justify-between items-center cursor-pointer",
            empty: "px-3 py-2 text-xs italic text-gray-400"
        }
    };

    let currentTheme = $derived(themes[variant] ?? themes.fantasy);

    function getItemLabel(item: SelectItem): string {
        return item.label ?? item.name ?? String(item.id);
    }

    let availableItems = $derived(
        items.filter(item =>
            !value.includes(item.id) &&
            getItemLabel(item).toLowerCase().includes(searchQuery.toLowerCase())
        )
    );

    let selectedObjects = $derived(
        value.map(val => {
            const found = items.find(i => i.id === val);
            if (found) {
                return { id: found.id, label: getItemLabel(found) };
            }
            return { id: val, label: String(val) };
        })
    );

    let exactMatch = $derived(
        items.some(item => getItemLabel(item).toLowerCase() === searchQuery.trim().toLowerCase()) ||
        value.some(val => String(val).toLowerCase() === searchQuery.trim().toLowerCase())
    );

    function handleSelect(item: SelectItem) {
        if (!value.includes(item.id)) {
            value = [...value, item.id];
        }
        searchQuery = '';
        inputRef?.focus();
    }

    function handleRemove(idToRemove: SelectValue) {
        value = value.filter(id => id !== idToRemove);
    }

    async function handleCreate() {
        const trimmed = searchQuery.trim();
        if (!trimmed || exactMatch || isCreating) return;

        isCreating = true;
        try {
            if (onCreate) {
                const newId = await onCreate(trimmed);
                if (newId !== null && newId !== undefined && !value.includes(newId)) {
                    value = [...value, newId];
                    searchQuery = '';
                }
            } else {
                if (!value.includes(trimmed)) {
                    value = [...value, trimmed];
                    searchQuery = '';
                }
            }
        } catch (err) {
            logError("Failed to forge new entry", err);
        } finally {
            isCreating = false;
        }
    }

    function handleWindowClick(e: MouseEvent) {
        if (inputRef && !inputRef.parentElement?.parentElement?.contains(e.target as Node)) {
            isOpen = false;
        }
    }
</script>

<svelte:window onclick={handleWindowClick} />

<div class={currentTheme.wrapper}>
    {#if labelText}
        <label class={currentTheme.label}>{labelText}</label>
    {/if}

    <div class={currentTheme.inputWrapper} onclick={() => inputRef?.focus()}>
        <!-- Hidden input for form validation -->
        <input
            type="text"
            class="sr-only absolute bottom-0 left-1/2 opacity-0 pointer-events-none"
            tabindex="-1"
            value={value.length > 0 ? 'valid' : ''}
            {required}
            {disabled}
            oninvalid={() => inputRef?.focus()}
        />

        <!-- Selected Pills -->
        {#each selectedObjects as item}
            <span class={currentTheme.pill}>
                {item.label}
                <button
                    type="button"
                    class={currentTheme.pillRemove}
                    disabled={disabled}
                    onclick={(e) => { e.stopPropagation(); handleRemove(item.id); }}
                >
                    &times;
                </button>
            </span>
        {/each}

        <!-- Input Field -->
        <input
            bind:this={inputRef}
            type="text"
            bind:value={searchQuery}
            onfocus={() => isOpen = true}
            onkeydown={(e) => {
                if (e.key === 'Enter') {
                    e.preventDefault();
                    if (!exactMatch && searchQuery && !selectOnly) {
                        handleCreate();
                    } else if (availableItems.length > 0) {
                        handleSelect(availableItems[0]);
                    }
                }
                if (e.key === 'Escape') {
                    isOpen = false;
                }
                if (e.key === 'Backspace' && searchQuery === '' && value.length > 0) {
                    handleRemove(value[value.length - 1]);
                }
            }}
            placeholder={value.length === 0 ? placeholder : ''}
            class={currentTheme.input}
            disabled={disabled || isCreating}
        />

        <!-- Dropdown Toggle Button -->
        <button
            type="button"
            onclick={(e) => { e.stopPropagation(); isOpen = !isOpen; }}
            class={currentTheme.icon}
            disabled={disabled}
        >
            <svg class="w-4 h-4 transition-transform duration-200 {isOpen ? 'rotate-180' : ''}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
            </svg>
        </button>
    </div>

    <!-- Dropdown Menu -->
    {#if isOpen}
        <div class={currentTheme.dropdown}>
            {#if availableItems.length > 0}
                <ul class="py-0 z-500">
                    {#each availableItems as item}
                        <li class="z-200">
                            <button
                                type="button"
                                class={currentTheme.item}
                                onclick={(e) => { e.stopPropagation(); handleSelect(item); }}
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
            {:else if availableItems.length === 0 && (exactMatch || selectOnly)}
                <div class={currentTheme.empty}>
                    No matching categories found.
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
