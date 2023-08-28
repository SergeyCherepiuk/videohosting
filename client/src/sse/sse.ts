export type ServerSentEvent = {
    id: number,
    eventName: string,
    data: string
}

export function parseEvent(event: string): ServerSentEvent {
    let lines = event.trimEnd().split("\n")
    if (lines.length != 3) {
        throw new Error("input string should have 3 lines after end-trimming")
    }
    return {
        id: +lines.at(0)!.slice(4),
        eventName: lines.at(1)!.slice(7),
        data: lines.at(2)!.slice(6)
    }
}