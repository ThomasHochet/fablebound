import { LogError } from '../../bindings/world-builder/app'

// Usable in case of not needing a full try-catch-finally.
// e.g.: const result = await catchLog(Function(...), ctx)
export function catchLog<T>(promise: Promise<T>, contextMessage: string): Promise<T | null> {
  return promise.catch((err) => {
    const errorMsg = err instanceof Error ? err.message : String(err)
    console.error(`[${contextMessage}]`, err)

    LogError(contextMessage, errorMsg).catch(() => { })
    return null
  })
}

// One liner logging method inside the catch block
export function logError(context: string, err: unknown) {
  const message = err instanceof Error ? err.message : String(err)

  // Log to console for dev
  console.error(`[${context}]`, err)

  // Fire-Forget again !
  LogError(context, message).catch(() => {})
}
