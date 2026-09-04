import { Events } from "@wailsio/runtime";

export const subscribe = (table: string, onUpdateCallback: Function) => {
  return Events.On("db:change", event => {
    const payload = event.data
    console.log("🔥 RAW EVENT RECEIVED:", event); // Let's see exactly what Wails gives us

    if (payload && payload.table === table) {
      console.log(`${table} table changed (${payload.action}), refreshing`)
      onUpdateCallback()
    }
  })
}
