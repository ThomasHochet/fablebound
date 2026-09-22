<script lang="ts">
    type SelectValue = number | string;

    type SelectItem = {
        id: SelectValue;
        label: string;
    };

    let {
        items = [],
        value = $bindable(null),
        placeholder = "Select...",
        labelText = "",
        disabled = false
    }: {
        items: SelectItem[];
        value: SelectValue | null;
        placeholder?: string;
        labelText?: string;
        disabled?: boolean;
    } = $props();

    let isOpen = $state(false);
    let containerRef: HTMLDivElement | null = $state(null);

    let selectedItem = $derived(items.find(i => i.id === value));

    function toggleDropdown(e: MouseEvent) {
        e.stopPropagation()
        if (!disabled) {
            isOpen = !isOpen;
        }
    }

    function handleSelect(e: MouseEvent, id: SelectValue) {
        e.stopPropagation()
        value = id;
        isOpen = false;
    }

    function handleWindowClick(e: MouseEvent) {
        if (containerRef && !containerRef.contains(e.target as Node)) {
            isOpen = false;
        }
    }
</script>

<svelte:window onclick={handleWindowClick} />

<div bind:this={containerRef} class="relative w-40 font-serif">
    {#if labelText}
        <label class="block text-xs font-serif text-(--ink-color) mb-1 text-left ml-1">
            {labelText}
        </label>
    {/if}

    <button
        type="button"
        onclick={toggleDropdown}
        {disabled}
        class="w-full bg-[#fcf8f2]/60 border border-[#b7a48d] rounded-md px-3 py-1.5 text-sm font-serif text-(--ink-color)/60! shadow-inner focus:outline-none flex items-center justify-between cursor-pointer transition-colors hover:bg-[#f5eedf]"
    >
        <span class={selectedItem ? "font-bold" : "text-(--ink-color)! italic"}>
            {selectedItem ? selectedItem.label : placeholder}
        </span>
        <svg class="w-4 h-4 text-[#8c7661] transition-transform duration-200 {isOpen ? 'rotate-180' : ''}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
        </svg>
    </button>

    {#if isOpen}
        <div class="absolute z-50 w-full mt-1 bg-[#f5eedf] border border-[#b7a48d] rounded-md shadow-lg max-h-60 overflow-y-auto custom-scrollbar backdrop-blur-sm">
            {#if items.length > 0}
                <ul class="py-1">
                    {#each items as item (item.id)}
                        <li>
                            <button
                                type="button"
                                class="w-full text-left px-3 py-1.5 text-sm font-serif text-(--ink-color)/80! hover:bg-[#e6d3b3] transition-colors border-b border-[#b7a48d]/20 cursor-pointer {value === item.id ? 'font-bold bg-[#e6d3b3]/60' : ''}"
                                onclick={(e) => handleSelect(e, item.id)}
                            >
                                {item.label}
                            </button>
                        </li>
                    {/each}
                </ul>
            {:else}
                <div class="px-3 py-2 text-xs italic text-[#8c7661]">
                    No options available...
                </div>
            {/if}
        </div>
    {/if}
</div>

<style>
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
