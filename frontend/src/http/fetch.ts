
export const enum EMethods {
    Get = "GET",
    Post = "POST",
    Put = "PUT",
    Delete = "DELETE"
}

const BASE_URL = "http://localhost:6060/api/v1"

export function useFetch() {

    function $fetch(method: EMethods, url: string, body?: any) {
        const apiUrl = BASE_URL + url

        const options = {
            method,
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(body)
        }
        return fetch(apiUrl, options)
    }

    return { $fetch }
}