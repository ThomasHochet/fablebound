<script lang="ts">
    import { Editor } from '@tiptap/core'
    import StarterKit from '@tiptap/starter-kit'
    import TextAlign from '@tiptap/extension-text-align'
    import { BulletList, OrderedList } from '@tiptap/extension-list'
    import { Details, DetailsContent, DetailsSummary } from '@tiptap/extension-details'
    import Image from '@tiptap/extension-image'
    import { onDestroy, onMount, untrack } from 'svelte';
    import type { Snippet } from 'svelte'
    import CornerGrey from "$lib/../assets/images/corners_grey.png"

    import TextLeftIcon from '@iconify-svelte/bi/components/t/text-left.svelte'
    import TextCenterIcon from '@iconify-svelte/bi/components/t/text-center.svelte'
    import TextRightIcon from '@iconify-svelte/bi/components/t/text-right.svelte'
    import TextJustifyIcon from '@iconify-svelte/bi/components/j/justify.svelte'
    import BulletListIcon from '@iconify-svelte/bi/components/l/list-task.svelte'
    import OrderedListIcon from '@iconify-svelte/bi/components/l/list-ol.svelte'
    import BlockquoteIcon from '@iconify-svelte/bi/components/b/blockquote-left.svelte'
    import BookIcon from '@iconify-svelte/bi/components/b/book.svelte'

    let {
      value = $bindable(''),
      disabled = false,
      children,
      classes = '',
    }: {
      value?: string,
      disabled?: boolean,
      children?: Snippet<[Snippet]>,
      classes?: string
    } = $props()

    let editor = $state<Editor | null>(null)
    let element = $state<HTMLDivElement>()
    let updatedTick = $state(0)

    onMount(() => {
      if (!element) return;

      editor = new Editor({
        element: element,
        extensions: [
          StarterKit.configure({
            blockquote: {
              HTMLAttributes: {
                class: 'fantasy-border fantasy-border-brown',
              }
            }
          }),
          Details.configure({
            HTMLAttributes: {
              class: 'fantasy-details fantasy-border fantasy-border-brown',
            }
          }),
          DetailsSummary,
          DetailsContent,
          TextAlign.configure({
            types: ['heading', 'paragraph']
          }),

          Image
        ],
        content: value,
        editable: !disabled,
        onTransaction: () => {
          updatedTick++
        },
        onUpdate: ({ editor: currentEditor}) => {
          const html = currentEditor.getHTML()
          if (html !== value) {
            value = html
          }
          updatedTick++
        },
        onSelectionUpdate: () => {
          updatedTick++
        }
      })
    })

    onDestroy(() => {
      editor?.destroy()
    })

    $effect(() => {
        const currentValue = value;

        untrack(() => {
          if (editor && currentValue !== editor.getHTML()) {
            editor.commands.setContent(currentValue, false);
          }
        });
      })

      $effect(() => {
        const currentDisabled = disabled;

        untrack(() => {
          if (editor) {
            editor.setEditable(!currentDisabled);
          }
        });
      })
</script>

{#if editor && !disabled}
    <div class="relative mb-1 text-left tool-bar rounded-lg z-10">
        <!-- <div class="wood-texture-bg"></div> -->

        <div class="iron-nail nail-tl"></div>
        <div class="iron-nail nail-tr"></div>
        <div class="iron-nail nail-bl"></div>
        <div class="iron-nail nail-br"></div>


        <button
            type="button"
            title="Bold (Ctrl+B)"
            onclick={() => editor?.chain().focus().toggleBold().run()}
            disabled={!editor?.can().chain().focus().toggleBold().run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive('bold')}
        >
            <span class="font-bold inline-block transition-transform duration-100 {updatedTick >= 0 && editor.isActive('bold') ? 'translate-y-[1px]' : ''}">
                B
            </span>
        </button>
        <button
            type="button"
            title="Ctrl + I"
            onclick={() => editor?.chain().focus().toggleItalic().run()}
            disabled={!editor?.can().chain().focus().toggleItalic().run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive('italic')}
        >
            <span class="italic inline-block transition-transform duration-100 {updatedTick >= 0 && editor.isActive('italic') ? 'translate-y-[1px]' : ''}">
                I
            </span>
        </button>
        <button
            type="button"
            title="Ctrl + U"
            onclick={() => editor?.chain().focus().toggleUnderline().run()}
            disabled={!editor?.can().chain().focus().toggleUnderline().run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive('underline')}
        >
            <span class="underline inline-block transition-transform duration-100 {updatedTick >= 0 && editor.isActive('underline') ? 'translate-y-[1px]' : ''}">
                U
            </span>
        </button>
        <button
            type="button"
            title="Ctrl + Shift + S"
            onclick={() => editor?.chain().focus().toggleStrike().run()}
            disabled={!editor?.can().chain().focus().toggleStrike().run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive('strike')}
        >
            <span class="line-through inline-block transition-transform duration-100 {updatedTick >= 0 && editor.isActive('strike') ? 'translate-y-[1px]' : ''}">
                abc
            </span>
        </button>
        <div class="tool-divider"></div>
        <button
            type="button"
            title="Align Left - Ctrl + Shift + L"
            onclick={() => editor?.chain().focus().setTextAlign('left').run()}
            disabled={!editor?.can().chain().focus().setTextAlign('left').run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive({ textAlign: 'left'})}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive({ textAlign: 'left' }) ? 'translate-y-[1px]' : ''}">
                <TextLeftIcon height="1em" />
            </span>
        </button>
        <button
            type="button"
            title="Align Center - Ctrl + Shift + E"
            onclick={() => editor?.chain().focus().setTextAlign('center').run()}
            disabled={!editor?.can().chain().focus().setTextAlign('center').run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive({ textAlign: 'center'})}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive({ textAlign: 'center'}) ? 'translate-y-[1px]' : ''}">
                <TextCenterIcon height="1em" />
            </span>
        </button>
        <button
            type="button"
            title="Align Right - Ctrl + Shift + R"
            onclick={() => editor?.chain().focus().setTextAlign('right').run()}
            disabled={!editor?.can().chain().focus().setTextAlign('right').run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive({ textAlign: 'right'})}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive({ textAlign: 'right'}) ? 'translate-y-[1px]' : ''}">
                <TextRightIcon height="1em" />
            </span>
        </button>
        <button
            type="button"
            title="Justify - Ctrl + Shift + J"
            onclick={() => editor?.chain().focus().setTextAlign('justify').run()}
            disabled={!editor?.can().chain().focus().setTextAlign('justify').run()}
            class="tool-btn text-sm"
            class:is-active={updatedTick >= 0 && editor.isActive({ textAlign: 'justify'})}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive({ textAlign: 'justify'}) ? 'translate-y-[1px]' : ''}">
                <TextJustifyIcon height="1em" />
            </span>
        </button>
        <div class="tool-divider"></div>
        <button
            type="button"
            title="Heading 1 - Ctrl + Alt + 1"
            onclick={() => editor?.chain().focus().toggleHeading({ level: 1}).run()}
            disabled={!editor?.can().chain().focus().toggleHeading({ level: 1}).run()}
            class="tool-btn text-sm px-1"
            class:is-active={updatedTick >= 0 && editor.isActive( 'heading', {level: 1})}
        >
            <span class="inline-block transition-transform duration-100 {updatedTick >= 0 && editor.isActive('heading', {level: 1}) ? 'translate-y-[1px]' : ''}">
                h1
            </span>
        </button>
        <button
            type="button"
            title="Heading 2 - Ctrl + Alt + 2"
            onclick={() => editor?.chain().focus().toggleHeading({ level: 2}).run()}
            disabled={!editor?.can().chain().focus().toggleHeading({ level: 2}).run()}
            class="tool-btn text-sm px-1"
            class:is-active={updatedTick >= 0 && editor.isActive( 'heading', {level: 2})}
        >
            <span class="inline-block transition-transform duration-100 {updatedTick >= 0 && editor.isActive('heading', {level: 2}) ? 'translate-y-[1px]' : ''}">
                h2
            </span>
        </button>
        <button
            type="button"
            title="Heading 3 - Ctrl + Alt + 3"
            onclick={() => editor?.chain().focus().toggleHeading({ level: 3}).run()}
            disabled={!editor?.can().chain().focus().toggleHeading({ level: 3}).run()}
            class="tool-btn text-sm px-1"
            class:is-active={updatedTick >= 0 && editor.isActive( 'heading', {level: 3})}
        >
            <span class="inline-block transition-transform duration-100 {updatedTick >= 0 && editor.isActive('heading', {level: 3}) ? 'translate-y-[1px]' : ''}">
                h3
            </span>
        </button>
        <div class="tool-divider"></div>
        <button
            type="button"
            title="Bullet list"
            onclick={() => editor?.chain().focus().toggleBulletList().run()}
            disabled={!editor?.can().chain().focus().toggleBulletList().run()}
            class="tool-btn text-sm px-1"
            class:is-active={updatedTick >= 0 && editor.isActive('bulletList')}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive('bulletList') ? 'translate-y-[1px]' : ''}">
                <BulletListIcon height="1em" />
            </span>
        </button>
        <button
            type="button"
            title="Ordered list"
            onclick={() => editor?.chain().focus().toggleOrderedList().run()}
            disabled={!editor?.can().chain().focus().toggleOrderedList().run()}
            class="tool-btn text-sm px-1"
            class:is-active={updatedTick >= 0 && editor.isActive('orderedList')}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive('orderedList') ? 'translate-y-[1px]' : ''}">
                <OrderedListIcon height="1em" />
            </span>
        </button>
        <button
            type="button"
            title="Blockquote"
            onclick={() => editor?.chain().focus().setBlockquote().run()}
            disabled={!editor?.can().chain().focus().setBlockquote().run()}
            class="tool-btn text-sm px-1"
            class:is-active={updatedTick >= 0 && editor.isActive('blockquote')}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive('blockquote') ? 'translate-y-[1px]' : ''}">
                <BlockquoteIcon height="1em" />
            </span>
        </button>
        <button
            type="button"
            title="Details"
            onclick={() => editor?.chain().focus().setDetails().run()}
            disabled={!editor?.can().chain().focus().setDetails().run()}
            class="tool-btn text-sm px-1"
            class:is-active={updatedTick >= 0 && editor.isActive('details')}
        >
            <span class="inline-block transition-transform duration-100 translate-y-0.75 {updatedTick >= 0 && editor.isActive('details') ? 'translate-y-[1px]' : ''}">
                <BookIcon height="1em" />
            </span>
        </button>
    </div>
{/if}

{#snippet editorBody()}
    <div bind:this={element} class="h-full overflow-y-auto prose prose-invert text-lg pt-2"></div>
{/snippet}

<div
    class={`px-5 parchment-base ${!disabled ? 'min-h-96 max-h-96 flex flex-col overflow-hidden cursor-text text-left' : (classes ?? '')}`}
    onclick={() => !disabled && editor?.chain().focus().run()}
>
    <!-- <img src="{CornerGrey}" alt="" class="absolute -top-2 -left-2 w-12 h-12 pointer-events-none z-10">
    <img src="{CornerGrey}" alt="" class="absolute -top-2 -right-2 w-12 h-12 pointer-events-none z-10 scale-x-[-1]" />
    <img src="{CornerGrey}" alt="" class="absolute -bottom-2 -left-2 w-12 h-12 pointer-events-none z-10 scale-y-[-1]" />
    <img src="{CornerGrey}" alt="" class="absolute -bottom-2 -right-2 w-12 h-12 pointer-events-none z-10 rotate-180" /> -->
<!-- <div
    class={`fantasy-border px-5 ${!disabled ? 'min-h-96 max-h-963 flex flex-col overflow-hidden cursor-text text-left' : (classes ?? '')}`}
    style="--inlay-bg: url('{coverImage}') center/cover; --inlay-filter: blur(1px) brightness(1);"
    onclick={() => !disabled && editor?.chain().focus().run()}
> -->
    {#if children}
        {@render children(editorBody)}
    {:else}
        {@render editorBody()}
    {/if}
</div>

<style>
  :global(.ProseMirror:focus) {
    outline: none;
  }

  :global(.ProseMirror) {
      outline: none;
  }

  :global(.ProseMirror h1) {
      font-size: 2.35rem;
  }

  :global(.ProseMirror h2) {
      font-size: 1.85rem;
  }

  :global(.ProseMirror h3) {
      font-size: 1.45rem;
  }

  :global(.ProseMirror ol), :global(.ProseMirror ul) {
      padding-left: 1.5rem;
      margin-bottom: 0.75rem;
  }

  :global(.ProseMirror ol) {
      list-style-type: decimal;
  }

  :global(.ProseMirror ul) {
      list-style-type: square;
  }

  :global(.ProseMirror ul li::marker) {
      color: #422419;
  }

  /*:global(.ProseMirror blockquote) {
      background: rgba(0, 0, 0, 60%);
      border-left: 10px solid #ccc;
      margin: 1.5em 10px;
      padding: 0.5em 10px;
      quotes: "\201C""\201D""\2018""\2019";
  }*/

  /*:global(.ProseMirror blockquote::before) {
      color: #ccc;
      content: open-quote;
      font-size: 4em;
      line-height: 0.1em;
      margin-right: 0.25em;
      vertical-align: -0.4em;
  }

  :global(.ProseMirror p) {
      display: inline;
  }*/

  :global(.ProseMirror blockquote) {
        position: relative;
        /*margin: 1rem 0;*/
        padding: 0.85rem 1rem 0.85rem 1.50rem;

        /* Background & Inset Shadow (Matches fantasy frame inner aesthetic) */
        background: transparent;
        border: none;
        box-shadow: none;

        /* 2. Delegate fill color to the fantasy border inlay */
        --inlay-bg: rgba(30, 22, 16, 0.5);

        /* Metallic Accent Bar on Left */
        border-left: 5px solid #e6c06f;
        border-radius: 0.25rem 0 0 0.25rem;

        /* Typography */
        font-style: italic;
        color: #d8c29d;
    }

    :global(.ProseMirror blockquote p) {
        margin: 0;
        line-height: 1.6;
    }

    /* Watermark Opening Quote */
    :global(.ProseMirror blockquote::before) {
        content: '“';
        position: absolute;
        font-family: serif;
        font-size: 3rem;
        line-height: 1;
        color: rgba(230, 192, 111, 0.2);
        pointer-events: none;
    }

    /*Details */
    :global(.ProseMirror div.fantasy-details) {
        display: flex;
        align-items: flex-start;
        gap: 0.5rem;
        margin: 1rem 0;
        padding: 0.75rem 1rem;
        /* 1. Make container transparent to stop rectangle bleeding */
        background: transparent;
        border: none;
        box-shadow: none;

        /* 2. Delegate fill color to the fantasy border inlay */
        --inlay-bg: rgba(30, 22, 16, 0.5);
    }

    /* 2. Built-in Expand/Collapse Button */
    :global(.ProseMirror div.fantasy-details > button) {
        background: transparent;
        border: none;
        color: #e6c06f;
        cursor: pointer;
        padding: 0;
        margin-top: 0.2rem;
        font-size: 0.75rem;
        line-height: 1;

        transition: transform 0.2s cubic-bezier(0.4, 0, 0.2, 1), color 0.2s ease;
    }

    /* Add a diamond glyph if the button icon is empty */
    :global(.ProseMirror div.fantasy-details > button::before) {
        content: '▶';
        display: inline-block;
    }

    :global(.ProseMirror div.fantasy-details > button:hover) {
        color: #fff0c2;
    }

    /* 3. Main Text Wrapper */
    :global(.ProseMirror div.fantasy-details > div) {
        flex: 1;
        display: flex;
        flex-direction: column;
    }

    /* 4. Summary / Header Line */
    :global(.ProseMirror div.fantasy-details summary) {
        font-family: serif;
        font-weight: 700;
        color: #e6c06f;
        outline: none;
        cursor: pointer;
        list-style: none;
    }

    /* 5. Collapsible Body (camelCase attribute match) */
    :global(.ProseMirror div[data-type="detailsContent"]) {
        margin-top: 0.5rem;
        padding-top: 0.5rem;
        border-top: 1px dashed rgba(230, 192, 111, 0.25);
        color: #d8c29d;
    }

    :global(.ProseMirror div.fantasy-details.is-open > button),
    :global(.ProseMirror div.fantasy-details:has(div[data-type="detailsContent"]:not([hidden])) > button) {
        transform: rotate(90deg);
        color: #ffd700;
    }
</style>
