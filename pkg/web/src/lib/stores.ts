import { writable } from "svelte/store";
import { browser } from "$app/environment"

export const token = writable(browser && localStorage.getItem("apiToken") || "");

token.subscribe(value => {
    if (value != null) {
        browser && localStorage.setItem("apiToken", value);
    }
});