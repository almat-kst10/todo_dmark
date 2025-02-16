
import { IListItem } from "../types.dto"
import { EMethods, useFetch } from "./fetch"

const { $fetch } = useFetch()

type IListItemBody = Pick<IListItem, 'title'>
type IListItemUpdate = Pick<IListItem, 'title' | 'is_done'>

interface ITodoApi {
    list: () => any
    add: (body: IListItemBody) => any
    detail: (id: number) => any
    update: (id: number, body: IListItemUpdate) => any
    delete: (id: number) => any
}


export const todoApi: ITodoApi = {
    list: () => $fetch(EMethods.Get, "/list"),
    add: (body) => $fetch(EMethods.Post, "/list", body),
    detail: (id) => $fetch(EMethods.Get, `/list/${id}`),
    update: (id, body) => $fetch(EMethods.Put, `/list/${id}`, body),
    delete: (id) => $fetch(EMethods.Delete, `/list/${id}`),
}