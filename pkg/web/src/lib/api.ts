import { token } from "./stores.js"
import { browser } from "$app/environment";
import axios from "axios";

let url = browser && window.location.origin || "";
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

export async function register({ email, username, password }) {
    const endpoint = url + "/api/user/register";
    const response = await axios({
        url: endpoint,
        method: "POST",
        data: JSON.stringify({
            email,
            username,
            password,
        }),
    })
    token.set(response.data.token);
    return response.data;
}


export async function updateToken() {
    const endpoint = url + "/api/auth/token";
    const response = await axios({
        url: endpoint,
        method: "POST",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
    })
    token.set(response.data.token);
    return response.data;
}

export async function getTags() {
    const endpoint = url + "/api/tag";
    const response = await axios.get(endpoint);
    return response.data;
}

export async function getTag({ tag }) {
    const endpoint = url + "/api/tag/" + tag;
    const response = await axios.get(endpoint);
    return response.data;
}

export async function getRelatedTags({ tag }) {
    const endpoint = url + "/api/tag-related/" + tag;
    const response = await axios.get(endpoint);
    return response.data;
}
export async function getTagAutocomplete({ tag, positive }) {
    let endpoint = url + "/api/tag-autocomplete?tag=" + tag;
    if (positive) {
        endpoint = endpoint + "&positive=true";
    }
    const response = await axios.get(endpoint);
    return response.data;
}
export async function getPosts({ page, q, perPage }: { page: any, q?: any, perPage?: any }) {
    if (!perPage) {
        perPage = 20;
    }
    let endpoint = url + "/api/post?page=" + page + "&perPage=" + perPage;
    if (q) {
        endpoint = url + "/api/post?tags=" + q + "&page=" + page + "&perPage=" + perPage;
    }
    const response = await axios.get(endpoint);
    return response.data;
}


export async function getPost({ id }) {
    const endpoint = url + "/api/post/" + id;
    const response = await axios(endpoint);
    return response.data;
}

export async function getPostCount() {
    const endpoint = url + "/api/post-count";
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

export async function searchBlob({ file, onProgress }) {
    var formData = new FormData();
    formData.append("file", file);
    const endpoint = url + "/api/blob/search";
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
    const endpoint = url + "/api/post/" + id;
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

export async function postDelete({ id }) {
    const endpoint = url + "/api/post/" + id;
    const response = await axios({
        url: endpoint,
        method: "DELETE",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
    })
    return response.status == 200;
}

export async function getUserProfile() {
    const endpoint = url + "/api/user/profile";
    const response = await axios({
        url: endpoint,
        method: "GET",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
    });
    console.log(response.data);
    return response.data;
}

export async function updateTagNotes(id, { note }) {
    const endpoint = url + "/api/tag/" + id + "/note";
    const response = await axios({
        url: endpoint,
        method: "PUT",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
        data: {
            note
        }
    })
    return response.data;
}

export async function updateTag(id, { name, tagTypeId }) {
    const endpoint = url + "/api/tag/" + id;
    const response = await axios({
        url: endpoint,
        method: "PUT",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
        data: {
            name, tagTypeId
        }
    })
    return response.data;
}

export async function updateUserProfile({ email, username, }) {
    const endpoint = url + "/api/user/update";
    const response = await axios({
        url: endpoint,
        method: "PUT",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
        data: {
            email, username,
        }
    })
    return response.data;
}
export async function updateUserPassword({ old_password, new_password }) {
    const endpoint = url + "/api/user/update-password";
    const response = await axios({
        url: endpoint,
        method: "PUT",
        headers: {
            'Authorization': 'Bearer ' + current_token,
        },
        withCredentials: true,
        data: {
            old_password, new_password
        }
    })
    return response.data;
}

export async function getTagTypes() {
    const endpoint = url + "/api/tagtype";
    const response = await axios.get(endpoint);
    return response.data;
}