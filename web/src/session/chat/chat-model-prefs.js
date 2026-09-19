// Browser-local model picker preferences (starred models + recent usage),
// persisted in localStorage and shared by every session page. Keys follow the
// `pi-web:v1:` convention used by other client-side settings.
const STARRED_KEY = 'pi-web:v1:model-starred';
const RECENTS_KEY = 'pi-web:v1:model-recents';
const MAX_RECENTS = 8;

function load(storage, key, fallback) {
  try {
    const raw = storage?.getItem?.(key);
    const parsed = raw ? JSON.parse(raw) : null;
    return parsed == null ? fallback : parsed;
  } catch {
    return fallback;
  }
}

function save(storage, key, value) {
  try {
    storage?.setItem?.(key, JSON.stringify(value));
  } catch {
    // storage unavailable (private mode, SSR) — prefs simply don't persist
  }
}

export function loadStarredModels(storage = globalThis.localStorage) {
  const value = load(storage, STARRED_KEY, []);
  return Array.isArray(value) ? value : [];
}

export function saveStarredModels(keys, storage = globalThis.localStorage) {
  save(storage, STARRED_KEY, keys);
}

export function loadRecentModels(storage = globalThis.localStorage) {
  const value = load(storage, RECENTS_KEY, []);
  return Array.isArray(value) ? value : [];
}

// Move key to the front; newest first, capped at MAX_RECENTS.
export function touchRecentModel(key, storage = globalThis.localStorage) {
  const next = [key, ...loadRecentModels(storage).filter((k) => k !== key)].slice(0, MAX_RECENTS);
  save(storage, RECENTS_KEY, next);
  return next;
}
