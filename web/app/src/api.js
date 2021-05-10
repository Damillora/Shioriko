import { token } from "./stores.js"
import axios from "axios";

let url = window.BASE_URL;
let current_token;
token.subscribe(value => {
    current_token = value;
})

export async function login({ username, password }) {
    const endpoint = url + "/api/auth/login";
    const response = await axios({
        url: endpoint,
        method: "POST",
        data: JSON.stringify({
            username,
            password,
        }),
    })
    token.set(response.data.token);
    return response.data;
}

export async function getPosts({ page }) {
    const endpoint = url + "/api/post?page=" + page;
    const response = await axios.get(endpoint);
    return response.data;
}

export async function getPostSearchTag({ page, q }) {
    if (q) {
        const endpoint = url + "/api/post?tags=" + q + "&page=" + page;
        const response = await axios(endpoint);
        return response.data;
    } else {
        const endpoint = url + "/api/post?page=" + page;
        const response = await axios(endpoint);
        return response.data;
    }
}

export async function getPost({ id }) {
    const endpoint = url + "/api/post/" + id;
    const response = await axios(endpoint);
    return response.data;
}

export async function uploadBlob({ file, onProgress }) {
    var formData = new FormData();
    formData.append("file", file);
    const endpoint = url + "/api/blob/upload";
    const response = await axios({
        url: endpoint,
        method: "POST",
        headers: {
            'Authorization': 'Bearer ' + current_token,
            'Content-Type': 'multipart/form-data',
        },
        withCredentials: true,
        data: formData,
        onUploadProgress: e => {
            if (onProgress) {
                onProgress(e)
            }
        }
    })
    return response.data;
}

export async function postCreate({ blob_id, source_url, tags }) {
    const endpoint = url + "/api/post/create";
    const response = await axios({
        url: endpoint,
        method: "POST",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
        data: {
            blob_id, source_url, tags
        }
    })
    return response.data;
}

export async function postUpdate(id, { source_url, tags }) {
    const endpoint = url + "/api/post/"+id;
    const response = await axios({
        url: endpoint,
        method: "POST",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
        data: {
            source_url, tags
        }
    })
    return response.data;
}