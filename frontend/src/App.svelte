<script lang="ts">
    // import Header from '$lib/components/layout/Header.svelte'
    import Editor from '$lib/components/features/Editor.svelte';
    import { OpenEditorWindow } from '$wails/world-builder/app.js';
    import { onMount } from 'svelte';
    import Overview from '$lib/components/layout/Overview.svelte';

    let hash = $state(window.location.hash)

    onMount(() => {
      const handleHashChange = () => {
        hash = window.location.hash
      }
      handleHashChange()
      window.addEventListener('hashchange', handleHashChange)
      return () => removeEventListener('hashchange', handleHashChange)
    })

    let isEditorWindow = $derived(hash.startsWith('#/editor/'))
    let currentSection = $derived(hash.split('#/editor/')[1] || '')

    function handleViewChange(section: string) {
      OpenEditorWindow(section).catch(err => {
        console.error("Failed to open window:", err)
      })

    }
</script>

<main class="bg">
    <svg width="0" height="0" style="position: absolute; z-index: -1;">
        <defs>
            <!-- Heavy Wood Grain Filter -->
            <filter id="wood-grain" x="0%" y="0%" width="100%" height="100%">
                <!-- Stretched frequency on X axis creates wood fibers -->
                <feTurbulence type="fractalNoise" baseFrequency="0.01 0.2" numOctaves="4" result="noise" />
                <feColorMatrix type="matrix" values="0.3 0 0 0 0  0.15 0 0 0 0  0.05 0 0 0 0  0 0 0 0.5 0" in="noise" result="coloredNoise" />
                <feComposite operator="in" in="coloredNoise" in2="SourceGraphic" result="composite" />
                <feBlend mode="multiply" in="composite" in2="SourceGraphic" />
            </filter>
        </defs>
    </svg>
    {#if isEditorWindow}
        <Editor section={currentSection} />
    {:else}
        <!-- <Header onSelect={handleViewChange} /> -->
        <Overview />
    {/if}
</main>
