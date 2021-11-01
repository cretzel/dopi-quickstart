class ItemService {

    async getItems() {
        const response = await fetch("/api/quickstart/items", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })

        if (response.ok) {
            return response.json();
        } else {
            console.log("Error fetching items");
        }
    }

    async postItem(item) {
        const response = await fetch("/api/quickstart/items", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(item)
        });
        if (response.ok) {
            return await response.json();
        }
        const json = await response.json();
        return Promise.reject(json.error)
    }

}

export default new ItemService()

