<script lang="ts">
    let {
      isReady = false,
      missingType = null
    }: {
      isReady: boolean,
      missingType?: 'general' | 'lore' | 'location' | 'character' | 'faction' | null
    } = $props();

    const generalWarnings = [
      "It seems the stars are not aligned today...",
      "A lack of perceptiong perhaps?",
      "A bit of ink is missing... check your fields.",
      "A blank parchment tells no tales."
    ]

    const loreWarnings = [
      "I cannot store this page correctly if you do not tell me where to put it!",
      "A chronicle without a category is lost to the void.",
      "A book without a name?! Do you know how awful it is to find it afterwards?!"
    ]

    const locationWarnings = [
      "A city in the middle of nowhere? Interesting...",
      "Where is this place of the map?",
      "I cannot seem to find this area on the map..."
    ]

    const characterWarnings = [
      ""
    ]

    const factionWarnings = [
      "What is it, then? A band of bards? A mercenary group?",
      "And... Who are they?"
    ]

    let currentMessage = $derived.by(() => {
      if (isReady) {
        return "Ready to be inscribed."
      }

      let pool = generalWarnings
      if (missingType === 'lore') {
        pool = [...generalWarnings, ...loreWarnings]
      } else if(missingType === 'location') {
        pool = [...generalWarnings, ...locationWarnings]
      } else if(missingType === 'faction') {
        pool = [...generalWarnings, ...factionWarnings]
      } else if(missingType === 'character') {
        pool = [...generalWarnings, ...characterWarnings]
      }

      const index = Math.abs(JSON.stringify(pool).length + (missingType?.length || 0)) % pool.length;
      return pool[index];
    })
</script>

<span class="ml-2 text-sm text-[#a38c71] italic font-cinzel">
    Status: {currentMessage}
</span>
