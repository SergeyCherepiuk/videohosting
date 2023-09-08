import { parseEvent, type ServerSentEvent } from "../sse/sse"

export function uploadWithProgress(
    fields: Map<string, string | Blob>,
    endpoint: string,
    onProgress: (progress: number) => void,
) {
    let formData = new FormData()
    for (let entry of fields.entries()) {
        if (typeof entry[1] == "string") {
            formData.append(entry[0], entry[1])
            continue
        }
        formData.append(entry[0], entry[1])
    }

    fetch(endpoint, {
        method: "POST",
        headers: {
            "Accept": "text/event-stream",
            "Cache-Control": "no-cache",
            "Connection": "keep-alive",
        },
        body: formData,
    }).then(async (response) => {
        if (response.ok && response.body) {
            let reader = response.body.getReader()
            let decoder = new TextDecoder("utf-8");

            while (true) {
                let { done, value } = await reader.read()
                if (done) return
                
                let event: ServerSentEvent = parseEvent(decoder.decode(value))
                let data: { progress: number } = JSON.parse(event.data)
                onProgress(data.progress * 100)
            }   
        } 
    })
}