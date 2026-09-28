<script lang="ts">
  import rafter from "$lib/../assets/images/1.png"

  let {
    direction = 'right',
    class: className = '',
    handleClick
  } = $props<{
    direction?: 'top' | 'top-right' | 'right' | 'bottom-right' | 'bottom' | 'bottom-left' | 'left' | 'top-left'
    class?: string
    handleClick?: (() => void) | null
  }>();

  // Map the chosen direction to the corresponding Tailwind rotation utility
  const rotations = {
    'right': 'rotate-0',
    'bottom-right': 'rotate-45',
    'bottom': 'rotate-90',
    'bottom-left': 'rotate-[135deg]',
    'left': 'rotate-180',
    'top-left': '-rotate-[135deg]',
    'top': '-rotate-90',
    'top-right': '-rotate-45'
  };

  let rotationClass = $derived(rotations[direction]);
</script>

<img
  src="{rafter}"
  alt=""
  class={`${handleClick ? 'cursor-pointer' : 'pointer-events-none'} transform origin-center ${rotationClass} ${className}`}
  aria-hidden="true"
  onclick={handleClick}
/>
