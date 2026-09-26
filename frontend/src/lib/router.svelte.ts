import Overview from "$lib/components/layout/Overview.svelte"
import Editor from "$lib/components/features/Editor.svelte"
import Reader from "$lib/components/features/Reader.svelte"
import type { Component } from "svelte"

class RouterState {

  #hash = $state(window.location.hash || '#/')

  constructor() {
    if (typeof window !== 'undefined') {
      window.addEventListener('hashchange', () => {
        this.#hash = window.location.hash || '#/'
      })
    }
  }

  // tl;dr full path without query string ("#/editor/character" => "/editor/character/id")
  get path() {
    return (this.#hash.split('?')[0] || '#/').replace(/^#/, '')
  }

  get current() {
    const p = this.path

    // Route: /reader/:section/:id (not optional)
    const readerMatch = p.match(/^\/reader\/([^/]+)\/([^/]+)$/)
    if (readerMatch) {
      return {
        component: Reader as Component,
        params: {
          section: readerMatch[1],
          id: readerMatch[2]
        }
      }
    }

    // Route: /editor/:section/:id? (id is optional, for edit only)
    const editorMatch = p.match(/^\/editor\/([^/]+)(?:\/([^/]+))?$/)
    if (editorMatch) {
      return {
        component: Editor as Component,
        params: {
          section: editorMatch[1],
          id: editorMatch[2]
        }
      }
    }

    return {
      component: Overview as Component,
      params: {}
    }
  }

}

export const router = new RouterState()
