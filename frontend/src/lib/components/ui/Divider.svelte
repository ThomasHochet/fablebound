<script>
    import divider from '$lib/../assets/images/divider-2_bcl.png'

    let { vertical = false, class: className = "", alt = "section divider" } = $props()
</script>

{#if vertical}
    <div class="relative flex items-center justify-center h-full min-h-56 w-8 py-4 mx-auto select-none overflow-hidden {className}">
        <img
            src={divider}
            alt={alt}
            class="rotate-90 max-w-none w-[18rem] sm:w-88 h-auto object-contain pointer-events-none filter drop-shadow-[0_2px_4px_rgba(0,0,0,0.6)]"
        >
    </div>
{:else}
    <div class="relative flex items-center justify-center w-full my-4 px-2 select-none {className}">
        <img
        src={divider}
        alt={alt}
        class="w-full max-w-2xl h-auto object-contain pointer-events-none filter drop-shadow-[0_2px_4px_rgba(0,0,0,0.6)]"
        />
    </div>
{/if}
