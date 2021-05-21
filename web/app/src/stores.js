import { writable } from "svelte/store";

const storedToken = localStorage.getItem("apiToken") ?? "";

export const token = writable(storedToken);
token.subscribe(value => {
    if (value != null) {
        localStorage.setItem("apiToken", value);
    }
});