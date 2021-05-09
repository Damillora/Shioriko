import { token } from "./stores.js"

let url = window.BASE_URL;
let current_token;
const unsub_token = token.subscribe(value => {
    current_token = token;
})
export async function login({ username, password }) {
    const endpoint = url + "/api/auth/login";
    const response = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify({
            username,
            password,
        }),
    })
    const data = await response.json();
    token.set(data.token);
    return data;
}

export async function getPosts({ page }) {
    const endpoint = url + "/api/post?page=" + page;
    const response = await fetch(endpoint);
    const data = await response.json();
    return data;
}

export async function getPostsTag({ page, tag }) {
    const endpoint = url + "/api/post/tag/" + tag + "?page=" + page;
    const response = await fetch(endpoint);
    const data = await response.json();
    return data;
}

export async function getPost({ id }) {
    const endpoint = url + "/api/post/" + id;
    const response = await fetch(endpoint);
    const data = await response.json();
    return data;
}