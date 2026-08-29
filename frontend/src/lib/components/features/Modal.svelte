<script lang="ts">
    import type { Snippet } from 'svelte';

    let {
        open = $bindable(false),
        children,
        onclose
    }: {
        open: boolean;
        children?: Snippet;
        onclose?: () => void;
    } = $props();

    function close() {
        open = false;
        onclose?.();
    }

    function handleKeydown(e: KeyboardEvent) {
        if (e.key === 'Escape') close();
    }
</script>

<svelte:window onkeydown={handleKeydown} />

{#if open}
    <!-- Fixed Backdrop Overlay -->
    <div
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm p-4"
        onclick={close}
        role="presentation"
    >
        <!-- Unstyled Content Container (Stops click propagation so clicking inside won't close) -->
        <div
            class="relative max-h-[90vh] w-full max-w-6xl overflow-y-auto"
            onclick={(e) => e.stopPropagation()}
            role="dialog"
            aria-modal="true"
            tabindex="0"
        >
            {@render children?.()}
        </div>
    </div>
{/if}
