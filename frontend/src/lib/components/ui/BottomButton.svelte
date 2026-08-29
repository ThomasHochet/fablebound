<script lang="ts">
    import type { HTMLButtonAttributes } from 'svelte/elements';
    import type { Snippet } from 'svelte';

    interface Props extends HTMLButtonAttributes {
        children?: Snippet;
        wrapperClass?: string;
    }

    let {
        children,
        type = 'button',
        class: customClass = '',
        wrapperClass = '',
        ...restProps
    }: Props = $props();
</script>

<div class={`fantasy-dock-root ${wrapperClass}`}>
    <div class="fantasy-dock">
        <div class="btn-wrapper">
            <button
                {type}
                class={`fantasy-btn-md fantasy-bone-n-coper ${customClass}`}
                {...restProps}
            >
                {#if children}
                    {@render children()}
                {/if}
            </button>
        </div>
    </div>
</div>

<style>
    .fantasy-dock-root {
        --banner-bg: #1e1610;
        --border-color: #e6c06f;
        --border-w: 2.5px;

        position: fixed;
        bottom: 0;
        left: 50%;
        transform: translateX(-50%);
        z-index: 40;
        pointer-events: none;
    }

    .fantasy-dock {
        position: relative;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 18px 36px 14px;
        background-color: var(--banner-bg);
        border: var(--border-w) solid var(--border-color);
        border-bottom: none;
        border-radius: 16px 16px 0 0;
        box-shadow: 0 -6px 20px rgba(0, 0, 0, 0.7);
        pointer-events: auto;
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .fantasy-dock:hover {
        padding: 22px 44px 18px;
        box-shadow: 0 -10px 28px rgba(0, 0, 0, 0.85);
    }

    .btn-wrapper {
        position: relative;
        z-index: 2;
    }

    .btn-wrapper :global(button) {
        transform: scale(0.9);
        transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .fantasy-dock:hover .btn-wrapper :global(button) {
        transform: scale(1.2);
    }

    .btn-wrapper :global(button):active {
        transform: scale(1.05);
    }
</style>
